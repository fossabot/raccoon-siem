package sdk

import (
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers"
)

type CollectorConfig struct {
	DefaultComponentSettings `yaml:",inline"`
	Connectors               []string `yaml:"sources,omitempty"`
	Parsers                  []string `yaml:"parsers,omitempty"`
	Destinations             []string `yaml:"destinations,omitempty"`
	AggregationRules         []string `yaml:"rules,omitempty"`
	Filters                  []string `yaml:"filters,omitempty"`
}

func (s *CollectorConfig) ID() string {
	return s.Name
}

type CollectorPackage struct {
	DefaultComponentSettings `yaml:",inline"`
	Normalizers              []normalizers.Config      `yaml:"normalizers,omitempty"`
	Connectors               []connectors.Config       `yaml:"connectors,omitempty"`
	Destinations             []DestinationSettings     `yaml:"destinations,omitempty"`
	Dictionaries             []DictionarySettings      `yaml:"dictionaries,omitempty"`
	AggregationRules         []AggregationRuleSettings `yaml:"rules,omitempty"`
	AggregationFilters       []FilterSettings          `yaml:"aggregationFilters,omitempty"`
	Filters                  []FilterSettings          `yaml:"filters,omitempty"`
}
