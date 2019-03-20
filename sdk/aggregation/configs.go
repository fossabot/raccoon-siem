package aggregation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"time"
)

type OutputFn func(event *normalization.Event)

type Config struct {
	Name            string         `yaml:"name,omitempty"`
	Filter          filters.Config `yaml:"filter,omitempty"`
	Threshold       int            `yaml:"threshold,omitempty"`
	Window          time.Duration  `yaml:"window,omitempty"`
	IdenticalFields []string       `yaml:"identical_fields,omitempty"`
	UniqueFields    []string       `yaml:"unique_fields,omitempty"`
	SumFields       []string       `yaml:"sum_fields,omitempty"`
	Recovery        bool           `yaml:"recovery,omitempty"`
}

func (r *Config) ID() string {
	return r.Name
}
