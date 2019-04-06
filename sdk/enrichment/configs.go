package enrichment

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/mutation"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

const (
	ValueSourceKindConst = "constant"
	ValueSourceKindDict  = "dictionary"
	ValueSourceKindEvent = "event"
	ValueSourceKindAL    = "activeList"
)

type Config struct {
	Field            string            `json:"field,omitempty"`
	Constant         interface{}       `json:"constant,omitempty"`
	KeyFields        []string          `json:"keyFields,omitempty"`
	ValueSourceKind  string            `json:"valueSourceKind,omitempty"`
	ValueSourceName  string            `json:"valueSourceName,omitempty"`
	ValueSourceField string            `json:"valueSourceField,omitempty"`
	Mutation         []mutation.Config `json:"mutation,omitempty"`
	Filter           *filters.Config   `json:"filters,omitempty"`
	filter           *filters.Filter   `json:"-"`
}

func (r *Config) Validate() error {
	if !helpers.EventFieldHasSetter(r.Field) {
		return fmt.Errorf("enrichment: invalid event field %s", r.Field)
	}

	for i := range r.Mutation {
		if err := r.Mutation[i].Validate(); err != nil {
			return err
		}
	}

	if r.Filter != nil {
		if err := r.Filter.Validate(); err != nil {
			return err
		}

		filter, err := filters.NewFilter(*r.Filter)
		if err != nil {
			return err
		}

		r.filter = filter
	}

	switch r.ValueSourceKind {
	case ValueSourceKindEvent:
		return r.validateEvent()
	case ValueSourceKindAL:
		return r.validateAL()
	case ValueSourceKindDict:
		return r.validateDict()
	default:
		return r.validateConstant()
	}
}

func (r *Config) validateConstant() error {
	r.Constant = normalization.ToFieldType(r.Field, r.Constant)
	if r.Constant == nil {
		return fmt.Errorf("enrichment: constant type must be convertable to field type")
	}
	return nil
}

func (r *Config) validateAL() error {
	if len(r.KeyFields) == 0 {
		return fmt.Errorf("enrichment: key fields required")
	}

	for _, f := range r.KeyFields {
		if !helpers.EventFieldHasGetter(f) {
			return fmt.Errorf("enrichment: invalid event field %s", f)
		}
	}

	if r.ValueSourceName == "" {
		return fmt.Errorf("enrichment: active list name required")
	}

	if r.ValueSourceField == "" {
		return fmt.Errorf("enrichment: active list field required")
	}

	return nil
}

func (r *Config) validateDict() error {
	if len(r.KeyFields) == 0 {
		return fmt.Errorf("enrichment: key fields required")
	}

	for _, f := range r.KeyFields {
		if !helpers.EventFieldHasGetter(f) {
			return fmt.Errorf("enrichment: invalid event field %s", f)
		}
	}

	if r.ValueSourceName == "" {
		return fmt.Errorf("enrichment: dictionary name required")
	}

	return nil
}

func (r *Config) validateEvent() error {
	if r.ValueSourceField != "" && !helpers.EventFieldHasGetter(r.ValueSourceField) {
		return fmt.Errorf("enrichment: invalid event field %s", r.ValueSourceField)
	}

	return nil
}
