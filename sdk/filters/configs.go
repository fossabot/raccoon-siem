package filters

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

const (
	CMPSourceKindEvent   = "event"
	CMPSourceKindContant = "constant"
	CMPSourceKindAL      = "activeList"
)

type Config struct {
	Name     string          `json:"name,omitempty"`
	Not      bool            `json:"not,omitempty"`
	Sections []SectionConfig `json:"sections,omitempty"`
}

func (r *Config) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("filter: name required")
	}

	if len(r.Sections) == 0 {
		return fmt.Errorf("filter: sections required")
	}

	for i := range r.Sections {
		if err := r.Sections[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (r *Config) ID() string {
	return r.Name
}

type JoinConfig struct {
	Name     string              `json:"name,omitempty"`
	Not      bool                `json:"not,omitempty"`
	Sections []JoinSectionConfig `json:"sections,omitempty"`
}

func (r *JoinConfig) ID() string {
	return r.Name
}

func (r *JoinConfig) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("filter: name required")
	}

	if len(r.Sections) == 0 {
		return fmt.Errorf("filter: sections required")
	}

	for i := range r.Sections {
		if err := r.Sections[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

type SectionConfig struct {
	Or         bool              `json:"or,omitempty"`
	Not        bool              `json:"not,omitempty"`
	Conditions []ConditionConfig `json:"conditions,omitempty"`
}

func (r *SectionConfig) Validate() error {
	if len(r.Conditions) == 0 {
		return fmt.Errorf("section: conditions required")
	}

	for i := range r.Conditions {
		if err := r.Conditions[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

type ConditionConfig struct {
	Field          string      `json:"field,omitempty"`
	KeyFields      []string    `json:"keyFields,omitempty"`
	Op             string      `json:"op,omitempty"`
	Constant       interface{} `json:"constant,omitempty"`
	CMPSourceKind  string      `json:"cmpSourceKind,omitempty"`
	CMPSourceName  string      `json:"cmpSourceName,omitempty"`
	CMPSourceField string      `json:"cmpSourceField,omitempty"`
}

func (r *ConditionConfig) Validate() error {
	if !helpers.EventFieldHasGetter(r.Field) {
		return fmt.Errorf("condition: invalid event field %s", r.Field)
	}

	switch r.Op {
	case OpEQ, OpNEQ, OpGT, OpGTorEQ, OpLT, OpLTorEQ, OpInSubnet, OpNotInSubnet, OpContains, OpNotContains:
	default:
		return fmt.Errorf("condition: unknown operator %s", r.Op)
	}

	switch r.CMPSourceKind {
	case CMPSourceKindAL:
		return r.validateAL()
	case CMPSourceKindEvent:
		return r.validateEvent()
	default:
		return r.validateConstant()
	}
}

func (r *ConditionConfig) validateConstant() error {
	r.Constant = normalization.ToFieldType(r.Field, r.Constant)
	if r.Constant == nil {
		return fmt.Errorf("condition: constant type must be convertable to field type")
	}
	return nil
}

func (r *ConditionConfig) validateAL() error {
	if len(r.KeyFields) == 0 {
		return fmt.Errorf("condition: key fields required")
	}

	for _, f := range r.KeyFields {
		if !helpers.EventFieldHasGetter(f) {
			return fmt.Errorf("condition: invalid event field %s", f)
		}
	}

	if r.CMPSourceName == "" {
		return fmt.Errorf("condition: active list name required")
	}

	if r.CMPSourceField == "" {
		return fmt.Errorf("condition: active list field required")
	}

	return nil
}

func (r *ConditionConfig) validateEvent() error {
	if !helpers.EventFieldHasGetter(r.CMPSourceField) {
		return fmt.Errorf("condition: invalid event field %s", r.CMPSourceField)
	}

	if !helpers.AreEventFieldTypesEqual(r.Field, r.CMPSourceField) {
		return fmt.Errorf("condition: right and left event fields must have the same type")
	}

	return nil
}

type JoinSectionConfig struct {
	Or         bool                  `json:"or,omitempty"`
	Not        bool                  `json:"not,omitempty"`
	Conditions []JoinConditionConfig `json:"conditions,omitempty"`
}

func (r *JoinSectionConfig) Validate() error {
	if len(r.Conditions) == 0 {
		return fmt.Errorf("section: conditions required")
	}

	for i := range r.Conditions {
		if err := r.Conditions[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

type JoinConditionConfig struct {
	LeftTag    string `json:"leftTag,omitempty"`
	LeftField  string `json:"leftField,omitempty"`
	Op         string `json:"op,omitempty"`
	RightTag   string `json:"rightTag,omitempty"`
	RightField string `json:"rightField,omitempty"`
}

func (r *JoinConditionConfig) Validate() error {
	if r.LeftTag == "" {
		return fmt.Errorf("condition: left event tag required")
	}

	if r.RightTag == "" {
		return fmt.Errorf("condition: right event tag required")
	}

	if !helpers.EventFieldHasGetter(r.LeftField) {
		return fmt.Errorf("condition: invalid event field")
	}

	if !helpers.EventFieldHasGetter(r.RightField) {
		return fmt.Errorf("condition: invalid event field")
	}

	switch r.Op {
	case OpEQ, OpNEQ, OpGT, OpGTorEQ, OpLT, OpLTorEQ, OpInSubnet, OpNotInSubnet, OpContains, OpNotContains:
	default:
		return fmt.Errorf("condition: unknown operator %s", r.Op)
	}

	if !helpers.AreEventFieldTypesEqual(r.LeftField, r.RightField) {
		return fmt.Errorf("condition: right and left event fields must have the same type")
	}

	return nil
}
