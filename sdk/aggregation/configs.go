package aggregation

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"time"
)

type OutputFn func(event *normalization.Event)

type Config struct {
	Name            string         `json:"name,omitempty"`
	Filter          filters.Config `json:"filter,omitempty"`
	Threshold       int            `json:"threshold,omitempty"`
	Window          int64          `json:"window,omitempty"`
	IdenticalFields []string       `json:"identicalFields,omitempty"`
	UniqueFields    []string       `json:"uniqueFields,omitempty"`
	SumFields       []string       `json:"sumFields,omitempty"`
	Recovery        bool           `json:"recovery,omitempty"`
}

func (r *Config) ID() string {
	return r.Name
}

func (r *Config) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("aggregation rule: name required")
	}

	if err := r.Filter.Validate(); err != nil {
		return err
	}

	if r.Window == 0 && r.Threshold == 0 {
		return fmt.Errorf("aggregation rule: threshold and/or window required")
	}

	for _, f := range r.IdenticalFields {
		if !helpers.IsEventFieldAccessable(f) {
			return fmt.Errorf("aggregation rule: invalid event field %s", f)
		}
	}

	for _, f := range r.UniqueFields {
		if !helpers.IsEventFieldAccessable(f) {
			return fmt.Errorf("aggregation rule: invalid event field %s", f)
		}
	}

	for _, f := range r.SumFields {
		if !helpers.IsEventFieldAccessable(f) {
			return fmt.Errorf("aggregation rule: invalid event field %s", f)
		}
	}

	return nil
}

func (r *Config) WindowDuration() time.Duration {
	return time.Duration(r.Window) * time.Second
}
