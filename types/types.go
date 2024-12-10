package types

import (
	"time"
)

// Main, top-level object for CDMs
type CausalDecisionModel struct {
	Schema       string        `json:"$schema"`
	Meta         Meta          `json:"meta"`
	Evaluatables []Evaluatable `json:"evaluatables,omitempty"`
	Diagrams     []Diagram     `json:"diagrams,omitempty"`
	Addons       any           `json:"addons,omitempty"`
}

// CDM with references to Evaluatables and Diagrams, instead of full objects
type RefsCausalDecisionModel struct {
	Schema       string                  `json:"$schema"`
	Meta         Meta                    `json:"meta"`
	Evaluatables []AssetRef[Evaluatable] `json:"evaluatables,omitempty"`
	Diagrams     []AssetRef[RefsDiagram] `json:"diagrams,omitempty"`
	Addons       any                     `json:"addons,omitempty"`
}

// A reference to another asset. Could be an Evaluatable, Diagram, Element, etc.
type AssetRef[T any] struct {
	UUID          string
	ContainingMap *map[string]T
}

// Base properties shared by all DI Assets
type Meta struct {
	UUID          string             `json:"uuid"`
	Name          string             `json:"name,omitempty"`
	Summary       string             `json:"summary,omitempty"`
	Documentation *MetaDocumentation `json:"documentation,omitempty"`
	Version       string             `json:"version,omitempty"`
	Draft         bool               `json:"draft,omitempty"`
	Creator       string             `json:"creator,omitempty"`
	CreatedDate   time.Time          `json:"createdDate,omitempty"`
	Updator       string             `json:"updator,omitempty"`
	UpdatedDate   time.Time          `json:"updatedDate,omitempty"`
}

// Full documentation for an asset/element/model
// Used by Meta
type MetaDocumentation struct {
	Content  string `json:"content,omitempty"`
	MIMEType string `json:"MIMEType,omitempty"`
}

// Properties of a Causal Decision Diagram
// Used by CausalDecisionModel
type Diagram struct {
	Meta         Meta               `json:"meta,omitempty"`
	Elements     []DiagramElement   `json:"elements,omitempty"`
	Dependencies []CausalDependency `json:"dependencies,omitempty"`
	Addons       any                `json:"addons,omitempty"`
}

type RefsDiagram struct {
	Meta         Meta                         `json:"meta,omitempty"`
	Elements     []AssetRef[DiagramElement]   `json:"elements,omitempty"`
	Dependencies []AssetRef[CausalDependency] `json:"dependencies,omitempty"`
	Addons       any                          `json:"addons,omitempty"`
}

// Properties for a Diagram Element in one of a Model's diagrams
// Used by Diagram
type DiagramElement struct {
	Meta               Meta            `json:"meta,omitempty"`
	CausalType         string          `json:"causalType,omitempty"`
	DiagramType        string          `json:"diaType,omitempty"`
	Content            *DiagramContent `json:"content,omitempty"`
	AssociatedElements []string        `json:"associatedElements,omitempty"`
}

// Custom version of DiagramElement, where Content is unrestricted
// Used by Diagram (TODO: Actually support this. Currently Elements is an array of DiagramElement.)
type CustomDiagramElement struct {
	Meta               Meta     `json:"meta,omitempty"`
	CausalType         string   `json:"causalType,omitempty"`
	DiagramType        string   `json:"diaType,omitempty"`
	Content            any      `json:"content,omitempty"`
	AssociatedElements []string `json:"associatedElements,omitempty"`
}

// Grab bag of properites
// Holds properties for a Point, Line, Box, Ellipse, or Text element
// Used by DiagramElement
type DiagramContent struct {
	Position        *Position  `json:"position,omitempty"`
	BoundingBoxSize *Size      `json:"boundingBoxSize,omitempty"`
	Vertices        []Position `json:"vertices,omitempty"`
	XRadius         float64    `json:"xRadius,omitempty"`
	YRadius         float64    `json:"yRadius,omitempty"`
	Text            string     `json:"text,omitempty"`
	Font            any        `json:"font,omitempty"`
}

// Used by DiagramContent
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Used by DiagramContent
type Size struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Used by Diagram
type CausalDependency struct {
	Meta   Meta   `json:"meta"`
	Source string `json:"source"`
	Target string `json:"target"`
}

// Properties for an Evaluatable Causal Model
// Used by CausalDecisionModel
type Evaluatable struct {
	Meta     Meta                 `json:"meta,omitempty"`
	Elements []EvaluatableElement `json:"evaluatables,omitempty"`
	Addons   any                  `json:"addons,omitempty"`
}

// Properties for an Evaluatable Element in one of a Model's evaluatables
// Used by Evaluatable
type EvaluatableElement struct {
	Meta            Meta               `json:"meta,omitempty"`
	CausalType      string             `json:"causalType,omitempty"`
	EvaluatableType string             `json:"evalType,omitempty"`
	Content         EvaluatableContent `json:"content,omitempty"`
	DefaultValue    any                `json:"defaultValue,omitempty"`
}

// Custom version of EvaluatableElement, where Content is unrestricted
// Used by Evaluatable (TODO: Actually support this. Currently Elements is an array of EvaluatableElement.)
type CustomEvaluatableElement struct {
	Meta            Meta   `json:"meta,omitempty"`
	CausalType      string `json:"causalType,omitempty"`
	EvaluatableType string `json:"evalType,omitempty"`
	Content         any    `json:"content,omitempty"`
	DefaultValue    any
}

// Grab bag of properties
// Holds properties for a Script, API Call, or Binary Base64 String element
// Used by EvaluatableElement
type EvaluatableContent struct {
	Script       string `json:"script,omitempty"`
	Language     string `json:"language,omitempty"`
	URIEndpoint  string `json:"uriEndpoint,omitempty"`
	Payload      any    `json:"payload,omitempty"`
	Base64String string `json:"base64String,omitempty"`
}

//---
// Based on this go.dev tutorial: https://go.dev/doc/tutorial/web-service-gin
//---
