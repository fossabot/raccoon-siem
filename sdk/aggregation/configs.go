package aggregation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"time"
)

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
