package sdk

import (
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers"
)

type CollectorConfig struct {
	Name         string                `yaml:"rules,omitempty"`
	Connector    connectors.Config     `yaml:"connector,omitempty"`
	Normalizer   normalizers.Config    `yaml:"normalizer,omitempty"`
	Filters      []filters.Config      `yaml:"filters,omitempty"`
	Enrichment   []enrichment.Config   `yaml:"enrichment,omitempty"`
	Rules        []aggregation.Config  `yaml:"rules,omitempty"`
	Destinations []destinations.Config `yaml:"destinations,omitempty"`
}

func (s *CollectorConfig) ID() string {
	return s.Name
}
