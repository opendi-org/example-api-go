package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"opendi.org/go-api/types"
)

var placeholderData = []types.CausalDecisionModel{
	{
		Schema: "Placeholder",
		Meta: types.Meta{
			UUID:    "4dcb00d2-4467-4853-8ed4-973d208c2ef1",
			Name:    "Test CDM for API",
			Summary: "This is a meaningless CDM, used to populate a test dataset.",
			Draft:   true},
		Diagrams: []types.Diagram{
			{
				Meta: types.Meta{
					UUID:  "f2ce5b78-fb8d-4f50-9129-c33a53580395",
					Name:  "Test CDM - CDD",
					Draft: true,
				},
				Elements: []types.DiagramElement{
					{
						Meta: types.Meta{
							UUID: "ecc9a714-5810-4deb-8aff-f0b15b09c1c9",
							Name: "Test Lever",
						},
						CausalType:  "Lever",
						DiagramType: "Box",
						Content: &types.DiagramContent{
							Position:        &types.Position{X: 0, Y: 0},
							BoundingBoxSize: &types.Size{Width: 400, Height: 200},
						},
					},
					{
						Meta: types.Meta{
							UUID: "e13f7b7a-3417-437a-bdc6-d9fd3689f884",
							Name: "Test Outcome"},
						CausalType:  "Outcome",
						DiagramType: "Box",
						Content: &types.DiagramContent{
							Position:        &types.Position{X: 500, Y: 0},
							BoundingBoxSize: &types.Size{Width: 400, Height: 200},
						},
					},
				},
				Dependencies: []types.CausalDependency{
					{
						Meta: types.Meta{
							UUID: "271cabe6-ac60-4c7d-ac9b-81ecff0a78c3",
							Name: "Test Lever --> Test Outcome"},
						Source: "ecc9a714-5810-4deb-8aff-f0b15b09c1c9",
						Target: "e13f7b7a-3417-437a-bdc6-d9fd3689f884",
					},
				},
			},
		},
	},
}

var modelsMap = map[string]types.RefsCausalDecisionModel{
	"4dcb00d2-4467-4853-8ed4-973d208c2ef1": {
		Schema: "Placeholder",
		Meta: types.Meta{
			UUID:    "4dcb00d2-4467-4853-8ed4-973d208c2ef1",
			Name:    "Test CDM for API",
			Summary: "This is a meaningless CDM, used to populate a test dataset.",
			Draft:   true},
		Diagrams: []types.AssetRef[types.RefsDiagram]{
			{
				UUID:          "f2ce5b78-fb8d-4f50-9129-c33a53580395",
				ContainingMap: &diagramsMap,
			},
		},
	},
}
var diagramsMap = map[string]types.RefsDiagram{
	"f2ce5b78-fb8d-4f50-9129-c33a53580395": {
		Meta: types.Meta{
			UUID:  "f2ce5b78-fb8d-4f50-9129-c33a53580395",
			Name:  "Test CDM - CDD",
			Draft: true,
		},
		Elements: []types.AssetRef[types.DiagramElement]{
			{
				UUID:          "ecc9a714-5810-4deb-8aff-f0b15b09c1c9",
				ContainingMap: &diagramElementsMap,
			},
			{
				UUID:          "e13f7b7a-3417-437a-bdc6-d9fd3689f884",
				ContainingMap: &diagramElementsMap,
			},
		},
		Dependencies: []types.AssetRef[types.CausalDependency]{
			{
				UUID:          "271cabe6-ac60-4c7d-ac9b-81ecff0a78c3",
				ContainingMap: &dependenciesMap,
			},
		},
	},
}
var diagramElementsMap = map[string]types.DiagramElement{
	"ecc9a714-5810-4deb-8aff-f0b15b09c1c9": {
		Meta: types.Meta{
			UUID: "ecc9a714-5810-4deb-8aff-f0b15b09c1c9",
			Name: "Test Lever",
		},
		CausalType:  "Lever",
		DiagramType: "Box",
		Content: &types.DiagramContent{
			Position:        &types.Position{X: 0, Y: 0},
			BoundingBoxSize: &types.Size{Width: 400, Height: 200},
		},
	},
	"e13f7b7a-3417-437a-bdc6-d9fd3689f884": {
		Meta: types.Meta{
			UUID: "e13f7b7a-3417-437a-bdc6-d9fd3689f884",
			Name: "Test Outcome"},
		CausalType:  "Outcome",
		DiagramType: "Box",
		Content: &types.DiagramContent{
			Position:        &types.Position{X: 500, Y: 0},
			BoundingBoxSize: &types.Size{Width: 400, Height: 200},
		},
	},
}
var dependenciesMap = map[string]types.CausalDependency{
	"271cabe6-ac60-4c7d-ac9b-81ecff0a78c3": {
		Meta: types.Meta{
			UUID: "271cabe6-ac60-4c7d-ac9b-81ecff0a78c3",
			Name: "Test Lever --> Test Outcome"},
		Source: "ecc9a714-5810-4deb-8aff-f0b15b09c1c9",
		Target: "e13f7b7a-3417-437a-bdc6-d9fd3689f884",
	},
}
var evaluatablesMap map[string]types.Evaluatable

// ENDPOINT IMPLEMENTATIONS
func getModels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, placeholderData)
}

func postModel(c *gin.Context) {
	var newModel types.CausalDecisionModel

	if err := c.BindJSON(&newModel); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, newModel)
	}

	placeholderData = append(placeholderData, newModel)
	c.IndentedJSON(http.StatusCreated, newModel)
}

func main() {
	router := gin.Default()
	router.GET("/v0/models", getModels)

	router.POST("/v0/models", postModel)

	router.Run("localhost:8080")
}
