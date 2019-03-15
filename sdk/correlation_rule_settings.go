package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
)

type CorrelationRuleSettings struct {
	Name        string                       `yaml:"name,omitempty"`
	Aggregation AggregationSettings          `yaml:"aggregation,omitempty"`
	Filter      string                       `yaml:"filter,omitempty"`
	Triggers    []TriggerSettings            `yaml:"triggers,omitempty"`
	Events      []EventSpecificationSettings `yaml:"events,omitempty"`
}

func (s *CorrelationRuleSettings) ID() string {
	return s.Name
}

func (s *CorrelationRuleSettings) Compile(filters []filters.IFilter) (*CorrelationRule, error) {
	var err error

	cr := new(CorrelationRule)

	// Name

	if s.Name == "" {
		return nil, fmt.Errorf("correlation rule must have a name")
	}

	cr.name = s.Name

	// filter

	for _, f := range filters {
		if f.ID() == s.Filter {
			cr.filter = f
			break
		}
	}

	// Aggregation spec

	cr.aggregation, err = s.Aggregation.compile()

	if err != nil {
		return nil, err
	}

	// Events

	for _, eventSetting := range s.Events {
		es, err := eventSetting.compile(filters)

		if err != nil {
			return nil, err
		}

		cr.eventSpecs = append(cr.eventSpecs, es)
	}

	// Actions and triggers

	cr.actions = make(map[string][]IAction)

	for _, trigger := range s.Triggers {
		kind, actSpecs, err := trigger.compile()

		if err != nil {
			return nil, err
		}

		actions := make([]IAction, 0)

		for _, actSpec := range actSpecs {
			actions = append(actions, NewAction(actSpec))
		}

		cr.actions[kind] = actions
	}

	return cr, nil
}
