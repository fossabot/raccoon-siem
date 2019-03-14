package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type aggregation struct {
	window          int
	threshold       int
	identicalFields []string
	uniqueFields    []string
	sumFields       []string
}

type AggregationSettings struct {
	Window          int      `yaml:"window,omitempty"`
	Threshold       int      `yaml:"threshold,omitempty"`
	IdenticalFields []string `yaml:"identicalFields,omitempty"`
	UniqueFields    []string `yaml:"uniqueFields,omitempty"`
	SumFields       []string `yaml:"sumFields,omitempty"`
}

func (s *AggregationSettings) compile() (*aggregation, error) {
	agg := &aggregation{
		window:    s.Window,
		threshold: s.Threshold,
	}

	if agg.threshold == 0 {
		agg.threshold = 1
	}

	// Identical fields

	for _, field := range s.IdenticalFields {
		_, err := ValidateEventFieldAndGetType(field)

		if err != nil {
			return nil, err
		}

		agg.identicalFields = append(agg.identicalFields, field)
	}

	// Unique fields

	for _, field := range s.UniqueFields {
		_, err := ValidateEventFieldAndGetType(field)

		if err != nil {
			return nil, err
		}

		agg.uniqueFields = append(agg.uniqueFields, field)
	}

	// Sum fields

	for _, field := range s.SumFields {
		ft, err := ValidateEventFieldAndGetType(field)

		if err != nil {
			return nil, err
		}

		switch ft {
		case normalization.FieldTypeInt, normalization.FieldTypeFloat, normalization.FieldTypeDuration:
		default:
			return nil, fmt.Errorf("field '%s' can not be summarized", field)
		}

		agg.sumFields = append(agg.sumFields, field)
	}

	return agg, nil
}
