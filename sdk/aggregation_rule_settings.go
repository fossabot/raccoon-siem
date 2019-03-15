package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
)

type AggregationRuleSettings struct {
	Name        string                       `yaml:"name,omitempty"`
	Aggregation AggregationSettings          `yaml:"aggregation,omitempty"`
	Events      []EventSpecificationSettings `yaml:"events,omitempty"`
}

func (s *AggregationRuleSettings) ID() string {
	return s.Name
}

func (s *AggregationRuleSettings) Compile(filters []filters.IFilter) (*AggregationRule, error) {
	var err error

	ar := new(AggregationRule)

	// Name

	if s.Name == "" {
		return nil, fmt.Errorf("aggregation rule must have a name")
	}

	ar.name = s.Name

	// Aggregation spec

	ar.aggregation, err = s.Aggregation.compile()

	if err != nil {
		return nil, err
	}

	// Events

	for _, eventSetting := range s.Events {
		es, err := eventSetting.compile(filters)

		if err != nil {
			return nil, err
		}

		ar.eventSpecs = append(ar.eventSpecs, es)
	}

	return ar, nil
}
