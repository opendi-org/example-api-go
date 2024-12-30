package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"opendi.org/go-api/apiTypes"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// API wraps a GORM database object (gorm.DB).
// This allows me to define gin endpoint functions that always have access to the API database,
// by making them member functions. See ENDPOINT IMPLEMENTATIONS below
type API struct {
	database *gorm.DB
}

// Create a new API instance, with a test database file.
// Migrate all data types so that GORM knows how to deal with them.
func NewAPI() (*API, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&apiTypes.CausalDecisionModel{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&apiTypes.Meta{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&apiTypes.Diagram{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&apiTypes.DiaElement{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&apiTypes.CausalDependency{})
	if err != nil {
		return nil, err
	}

	return &API{database: db}, nil
}

// ENDPOINT IMPLEMENTATIONS

// GET /v0/models
// Returns a list of all apiTypes.Meta objects the user has access to.
// Currently limited to 10 results, no pagination, no sorting.
// TODO: Pagination, sorting
func (api *API) getModels(c *gin.Context) {
	var foundMetas []apiTypes.Meta
	api.database.Limit(10).Model(&apiTypes.Meta{}).Joins("JOIN causal_decision_models ON causal_decision_models.meta_id = meta.id").Distinct().Find(&foundMetas)

	c.IndentedJSON(http.StatusOK, foundMetas)
}

// GET /v0/models/:modelId/full
// Returns the FULL JSON object for the model with the given model ID (UUID)
// TODO: More consideration for possible duplicate database entries with the same UUID?
func (api *API) getModelById(c *gin.Context) {
	id := c.Param("modelId")

	foundCDM, err := api.retrieveFullModel(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error)
	}

	c.IndentedJSON(http.StatusOK, foundCDM)
}

// GET /v0/models/:modelId
// Returns the JSON meta object for the model with the given model ID (UUID)
// TODO: More consideration for possible duplicate database entries with the same UUID?
func (api *API) getModelMetaById(c *gin.Context) {
	id := c.Param("modelId")

	var foundMeta apiTypes.Meta
	//This JOIN statement ensures we only search for Meta objects that represent CDMs.
	api.database.Joins("JOIN causal_decision_models ON causal_decision_models.meta_id = meta.id").First(&foundMeta, "uuid = ?", id)

	if foundMeta.UUID != "" {
		c.IndentedJSON(http.StatusOK, foundMeta)
	} else {
		c.IndentedJSON(http.StatusNotFound, nil)
	}
}

// POST /v0/models
// Add the given model object to the database
// TODO: Better feedback for input validation errors. Consideration for attempting to POST a model with an already-stored UUID. Should use PUT instead?
func (api *API) postModel(c *gin.Context) {
	var newModel apiTypes.CausalDecisionModel

	if err := c.BindJSON(&newModel); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, newModel)
	}

	if exists := api.checkModelExists(newModel.Meta.UUID); !exists {
		api.database.Create(&newModel)
		c.IndentedJSON(http.StatusCreated, newModel)
	} else {
		c.IndentedJSON(http.StatusBadRequest, nil)
	}
}

func (api *API) putModel(c *gin.Context) {
	var updatedModel apiTypes.CausalDecisionModel

	if err := c.BindJSON(&updatedModel); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, updatedModel)
	}

	currentModel, err := api.retrieveFullModel(updatedModel.Meta.UUID)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	api.database.Session(&gorm.Session{FullSaveAssociations: true}).Model(&currentModel).Updates(&updatedModel)

	retrievedUpdatedModel, err := api.retrieveFullModel(updatedModel.Meta.UUID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Error retrieving model after update")
	}

	c.IndentedJSON(http.StatusOK, retrievedUpdatedModel)
}

// Query database for a model Meta object with UUID == modelId
// Returns true if a matching entry is found
func (api *API) checkModelExists(modelId string) bool {
	var foundMeta apiTypes.Meta
	//This JOIN statement ensures we only search for Meta objects that represent CDMs.
	api.database.Joins("JOIN causal_decision_models ON causal_decision_models.meta_id = meta.id").First(&foundMeta, "uuid = ?", modelId)

	return foundMeta.UUID == modelId
}

// Query database for a full CausalDecisionModel object with Meta.UUID == modelId
// Returns the complete model if a matching entry is found
func (api *API) retrieveFullModel(modelId string) (apiTypes.CausalDecisionModel, error) {
	if api.checkModelExists(modelId) {
		var modelMeta apiTypes.Meta
		api.database.First(&modelMeta, "uuid = ?", modelId)

		var model apiTypes.CausalDecisionModel
		api.database.Preload("Meta").Preload("Diagrams").Preload("Diagrams.Meta").Preload("Diagrams.Elements").Preload("Diagrams.Elements.Meta").Preload("Diagrams.Dependencies").Preload("Diagrams.Dependencies.Meta").First(&model, "meta_id = ?", modelMeta.ID)
		if model.Meta.UUID == modelId {
			return model, nil
		}
	}

	return apiTypes.CausalDecisionModel{}, errors.New("Model with ID " + modelId + " does not exist.")
}

// Main function.
// Set up API object.
// (Temporary) seed the database with a placeholder CDM object.
// Gin: Define endpoint paths for functions implemented above.
// Finally, set up a test API instance on localhost:8080.
func main() {

	api, err := NewAPI()
	if err != nil {
		panic(err)
	}

	newData := apiTypes.CausalDecisionModel{
		Schema: "Placeholder",
		Meta: apiTypes.Meta{
			UUID:    "18c731e4-6215-4908-b094-7be07ef17c98",
			Name:    "Test CDM Meta",
			Summary: "This is testing how GORM works with complicated data types",
		},
		Diagrams: []apiTypes.Diagram{
			{
				Meta: apiTypes.Meta{
					UUID:    "5fcacd4f-d14e-45bf-b1f4-65cf9498f642",
					Name:    "Test CDM Diagram",
					Summary: "This tests diagram data for the test CDM",
				},
				Elements: []apiTypes.DiaElement{
					{
						Meta: apiTypes.Meta{
							UUID: "ca843ab9-3058-4e9d-8633-18812c6a955b",
							Name: "Test Lever",
						},
						CausalType:  "Lever",
						DiagramType: "Box",
						Content:     []byte(`{"position": {"x": 0, "y": 0}, "boundingBoxSize": {"width": 400, "height": 200}}`),
					},
					{
						Meta: apiTypes.Meta{
							UUID: "3bf61246-1473-4e70-beca-9d60275aaeb7",
							Name: "Test Outcome",
						},
						CausalType:  "Outcome",
						DiagramType: "Box",
						Content:     []byte(`{"position": {"x": 500, "y": 0}, "boundingBoxSize": {"width": 400, "height": 200}}`),
					},
				},
				Dependencies: []apiTypes.CausalDependency{
					{
						Meta: apiTypes.Meta{
							UUID: "edfb963b-2031-426f-b8ab-393800dbd8ec",
							Name: "Test Lever --> Test Outcome",
						},
						Source: "ca843ab9-3058-4e9d-8633-18812c6a955b",
						Target: "3bf61246-1473-4e70-beca-9d60275aaeb7",
					},
				},
			},
		},
	}

	if !api.checkModelExists(newData.Meta.UUID) {
		api.database.Create(&newData)
	}

	router := gin.Default()
	router.GET("/v0/models", api.getModels)
	router.GET("/v0/models/:modelId/full", api.getModelById)
	router.GET("/v0/models/:modelId", api.getModelMetaById)

	router.POST("/v0/models", api.postModel)
	router.PUT("/v0/models", api.putModel)

	router.Run("localhost:8080")
}

//---
// Partially based on this go.dev tutorial: https://go.dev/doc/tutorial/web-service-gin
//---
