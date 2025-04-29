/**
 * Define test data for ensuring the database works with complex CDM examples.
 */

package apiTypes

import (
	"github.com/google/uuid"
)

// Generates a simple adder CDM. Should contain two Levers with interactive sliders for sum inputs,
// and one Outcome with a non-interactive slider showing the sum output
func GetTestModel() CausalDecisionModel {
	// Pre-generate UUIDs that need to match
	modelUUID := uuid.NewString()
	diagramUUID := uuid.NewString()

	sumInput1UUID := uuid.NewString()
	sumInput2UUID := uuid.NewString()
	sumOutputUUID := uuid.NewString()

	sumInput1ControlUUID := uuid.NewString()
	sumInput2ControlUUID := uuid.NewString()
	sumOutputControlUUID := uuid.NewString()

	sumInput1DisplayUUID := uuid.NewString()
	sumInput2DisplayUUID := uuid.NewString()
	sumOutputDisplayUUID := uuid.NewString()

	input1ToOutputDependencyUUID := uuid.NewString()
	input2ToOutputDependencyUUID := uuid.NewString()

	runnableModelUUID := uuid.NewString()
	addEvalElementUUID := uuid.NewString()
	addEvalAssetUUID := uuid.NewString()

	return CausalDecisionModel{
		Schema: "Placeholder",
		Meta: Meta{
			UUID:        modelUUID,
			Name:        "Test Model",
			Summary:     "This model confirms that the API and database work correctly.\n\nIt uses many of the Go runtime type definitions for database storage to confirm you're storing all of these components correctly:\n- Diagrams\n- Diagram Elements\n- Displays\n- Runnable models\n- Evaluatable elements\n- Evaluatable assets\n- Controls\n- Input/Output values",
			Version:     "0.1",
			Draft:       true,
			CreatedDate: "2025-04-29T12:31:37-04:00",
		},
		RunnableModels: []RunnableModel{
			{
				Meta: Meta{
					UUID: runnableModelUUID,
					Name: "Test runnable model",
				},
				Elements: []EvalElement{
					{
						Meta: Meta{
							UUID: addEvalElementUUID,
							Name: "Sum inputs",
						},
						Inputs:       []byte(`["` + sumInput1UUID + `","` + sumInput2UUID + `"]`),
						Outputs:      []byte(`["` + sumOutputUUID + `"]`),
						FunctionName: "add",
						EvalAsset:    addEvalAssetUUID,
					},
				},
			},
		},
		EvalAssets: []EvalAsset{
			{
				Meta: Meta{
					UUID: addEvalAssetUUID,
					Name: "Add script",
				},
				EvalType: "Script",
				Content:  []byte(`{"script":"KGZ1bmN0aW9uICgpIHsKICBjb25zdCBhZGQgPSBmdW5jdGlvbiAodmFscykgewogICAgbGV0IHN1bSA9IDA7CiAgICB2YWxzLmZvckVhY2goKHZhbCkgPT4gewogICAgICBzdW0gKz0gdmFsOwogICAgfSk7CiAgICByZXR1cm4gW3N1bV07CiAgfTsKCiAgcmV0dXJuIHsgZnVuY01hcDogeyAiYWRkIjogYWRkIH0gfTsKfSkoKTs=","language":"javascript"}`),
			},
		},
		IOValues: []InputOutputValue{
			{
				Meta: Meta{
					UUID: sumInput1UUID,
					Name: "Add Input 1",
				},
				Data: []byte(`30`),
			},
			{
				Meta: Meta{
					UUID: sumInput2UUID,
					Name: "Add Input 2",
				},
				Data: []byte(`27`),
			},
			{
				Meta: Meta{
					UUID: sumOutputUUID,
					Name: "Add Output",
				},
				Data: []byte(`null`),
			},
		},
		Controls: []Control{
			{
				Meta: Meta{
					UUID: sumInput1ControlUUID,
					Name: "Control: Sum Input 1",
				},
				IOValues: []byte(`["` + sumInput1UUID + `"]`),
				Displays: []byte(`["` + sumInput1DisplayUUID + `"]`),
			},
			{
				Meta: Meta{
					UUID: sumInput2ControlUUID,
					Name: "Control: Sum Input 2",
				},
				IOValues: []byte(`["` + sumInput2UUID + `"]`),
				Displays: []byte(`["` + sumInput2DisplayUUID + `"]`),
			},
			{
				Meta: Meta{
					UUID: sumOutputControlUUID,
					Name: "Control: Sum Output",
				},
				IOValues: []byte(`["` + sumOutputUUID + `"]`),
				Displays: []byte(`["` + sumOutputDisplayUUID + `"]`),
			},
		},
		Diagrams: []Diagram{
			{
				Meta: Meta{
					UUID:        diagramUUID,
					Name:        "Test Diagram",
					CreatedDate: "2025-04-29T12:31:37-04:00",
				},
				Elements: []DiaElement{
					{
						Meta: Meta{
							UUID: sumInput1UUID,
							Name: "Sum Input 1",
						},
						CausalType: "Lever",
						Position:   []byte(`{"x":87.5,"y":256.5}`),
						Displays: []DiaDisplay{
							{
								Meta: Meta{
									UUID: sumInput1DisplayUUID,
									Name: "",
								},
								DisplayType: "controlRange",
								Content:     []byte(`{"controlParameters":{"min":0,"max":50,"step":1,"value":25,"isInteractive":true}}`),
							},
						},
					},
					{
						Meta: Meta{
							UUID: sumInput2UUID,
							Name: "Sum Input 2",
						},
						CausalType: "Lever",
						Position:   []byte(`{"x":89.5,"y":439.5}`),
						Displays: []DiaDisplay{
							{
								Meta: Meta{
									UUID: sumInput2DisplayUUID,
									Name: "",
								},
								DisplayType: "controlRange",
								Content:     []byte(`{"controlParameters":{"min":0,"max":50,"step":1,"value":25,"isInteractive":true}}`),
							},
						},
					},
					{
						Meta: Meta{
							UUID: sumOutputUUID,
							Name: "Sum Output",
						},
						CausalType: "Outcome",
						Position:   []byte(`{"x":500,"y":348}`),
						Displays: []DiaDisplay{
							{
								Meta: Meta{
									UUID: sumOutputDisplayUUID,
									Name: "Sum:",
								},
								DisplayType: "controlRange",
								Content:     []byte(`{"controlParameters":{"min":0,"max":100,"step":1,"value":50,"isInteractive":false}}`),
							},
						},
					},
				},
				Dependencies: []CausalDependency{
					{
						Meta: Meta{
							UUID: input1ToOutputDependencyUUID,
							Name: "Sum input 1 --> Sum output",
						},
						Source: sumInput1UUID,
						Target: sumOutputUUID,
					},
					{
						Meta: Meta{
							UUID: input2ToOutputDependencyUUID,
							Name: "Sum input 2 --> Sum output",
						},
						Source: sumInput2UUID,
						Target: sumOutputUUID,
					},
				},
			},
		},
	}
}
