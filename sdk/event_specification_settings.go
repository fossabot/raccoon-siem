package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
)

type EventSpecificationSettings struct {
	ID        string `yaml:"id,omitempty"`
	Threshold int    `yaml:"threshold,omitempty"`
	Filter    string `yaml:"filter,omitempty"`
}

func (s *EventSpecificationSettings) compile(filters []filters.IFilter) (*eventSpecification, error) {
	es := &eventSpecification{
		id:        s.ID,
		threshold: s.Threshold,
	}

	for _, f := range filters {
		if f.ID() == s.Filter {
			es.filter = f
			break
		}
	}

	if es.filter == nil {
		return nil, fmt.Errorf("event specification '%s' has no filter", es.id)
	}

	return es, nil
}
