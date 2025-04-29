/**
 * Define all Go runtime type representations for CDM structures.
 * These use JSON and GORM struct tags to define rules for JSON encoding/decoding
 * and GORM settings.
 * See https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go#using-struct-tags-to-control-encoding
 */

package apiTypes

import (
	"encoding/json"
	"time"
)

type CausalDecisionModel struct {
	ID             int                `gorm:"primaryKey" json:"-"`
	CreatedAt      time.Time          `json:"-"`
	UpdatedAt      time.Time          `json:"-"`
	Schema         string             `json:"$schema"`
	MetaID         int                `gorm:"index;constraint:OnDelete:CASCADE;" json:"-"`
	Meta           Meta               `json:"meta"`
	Diagrams       []Diagram          `gorm:"many2many:cdm_diagrams;constraint:OnDelete:CASCADE;" json:"diagrams,omitempty"`
	IOValues       []InputOutputValue `gorm:"many2many:cdm_inputOutputValues;constraint:OnDelete:CASCADE;" json:"inputOutputValues,omitempty"`
	Controls       []Control          `gorm:"many2many:cdm_controls;constraint:OnDelete:CASCADE;" json:"controls,omitempty"`
	RunnableModels []RunnableModel    `gorm:"many2many:cdm_runnableModels;constraint:OnDelete:CASCADE;" json:"runnableModels,omitempty"`
	EvalAssets     []EvalAsset        `gorm:"many2many:cdm_evaluatableAssets;constraint:OnDelete:CASCADE;" json:"evaluatableAssets,omitempty"`
	Addons         json.RawMessage    `json:"addons,omitempty"`
}

type Meta struct {
	ID            int             `gorm:"primaryKey" json:"-"`
	CreatedAt     time.Time       `json:"-"`
	UpdatedAt     time.Time       `json:"-"`
	UUID          string          `json:"uuid"`
	Name          string          `json:"name,omitempty"`
	Summary       string          `json:"summary,omitempty"`
	Documentation json.RawMessage `json:"documentation,omitempty"`
	Version       string          `json:"version,omitempty"`
	Draft         bool            `json:"draft,omitempty"`
	Creator       string          `json:"creator,omitempty"`
	CreatedDate   string          `json:"createdDate,omitempty"`
	Updator       string          `json:"updator,omitempty"`
	UpdatedDate   string          `json:"updatedDate,omitempty"`
}

type Diagram struct {
	ID           int                `gorm:"primaryKey" json:"-"`
	CreatedAt    time.Time          `json:"-"`
	UpdatedAt    time.Time          `json:"-"`
	MetaID       int                `json:"-"`
	Meta         Meta               `json:"meta"`
	Elements     []DiaElement       `gorm:"many2many:diagram_elements;constraint:OnDelete:CASCADE;" json:"elements,omitempty"`
	Dependencies []CausalDependency `gorm:"many2many:diagram_dependencies;constraint:OnDelete:CASCADE;" json:"dependencies,omitempty"`
	Addons       json.RawMessage    `json:"addons,omitempty"`
}

type DiaElement struct {
	ID         int             `gorm:"primaryKey" json:"-"`
	CreatedAt  time.Time       `json:"-"`
	UpdatedAt  time.Time       `json:"-"`
	MetaID     int             `json:"-"`
	Meta       Meta            `json:"meta"`
	CausalType string          `json:"causalType"`
	Position   json.RawMessage `json:"position"`
	Displays   []DiaDisplay    `gorm:"many2many:diagram_displays;constraint:OnDelete:CASCADE;" json:"displays,omitempty"`
	Addons     json.RawMessage `json:"addons,omitempty"`
}

type DiaDisplay struct {
	ID          int             `gorm:"primaryKey" json:"-"`
	CreatedAt   time.Time       `json:"-"`
	UpdatedAt   time.Time       `json:"-"`
	MetaID      int             `json:"-"`
	Meta        Meta            `json:"meta"`
	Content     json.RawMessage `json:"content"`
	DisplayType string          `json:"displayType"`
}

type CausalDependency struct {
	ID        int       `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	MetaID    int       `json:"-"`
	Meta      Meta      `json:"meta"`
	Source    string    `json:"source"`
	Target    string    `json:"target"`
}

type InputOutputValue struct {
	ID        int             `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	MetaID    int             `json:"-"`
	Meta      Meta            `json:"meta"`
	Data      json.RawMessage `json:"data"`
}

type Control struct {
	ID        int             `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	MetaID    int             `json:"-"`
	Meta      Meta            `json:"meta"`
	IOValues  json.RawMessage `json:"inputOutputValues"`
	Displays  json.RawMessage `json:"displays"`
}

type RunnableModel struct {
	ID        int             `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	MetaID    int             `json:"-"`
	Meta      Meta            `json:"meta"`
	Elements  []EvalElement   `gorm:"many2many:runnableModel_elements;constraint:OnDelete:CASCADE;" json:"elements"`
	Addons    json.RawMessage `json:"addons,omitempty"`
}

type EvalAsset struct {
	ID        int             `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	MetaID    int             `json:"-"`
	Meta      Meta            `json:"meta"`
	EvalType  string          `json:"evalType"`
	Content   json.RawMessage `json:"content"`
}

type EvalElement struct {
	ID           int             `gorm:"primaryKey" json:"-"`
	CreatedAt    time.Time       `json:"-"`
	UpdatedAt    time.Time       `json:"-"`
	MetaID       int             `json:"-"`
	Meta         Meta            `json:"meta"`
	Inputs       json.RawMessage `json:"inputs"`
	Outputs      json.RawMessage `json:"outputs"`
	FunctionName string          `json:"functionName"`
	EvasAssetID  int             `json:"-"`
	EvalAsset    string          `json:"evaluatableAsset"`
	Addons       json.RawMessage `json:"addons,omitempty"`
}
