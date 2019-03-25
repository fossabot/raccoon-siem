package collector

import (
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers"
)

type Config struct {
	Name         string                `yaml:"name,omitempty"`
	Connector    connectors.Config     `yaml:"connector,omitempty"`
	Normalizer   normalizers.Config    `yaml:"normalizer,omitempty"`
	Filters      []filters.Config      `yaml:"dropFilters,omitempty"`
	Enrichment   []enrichment.Config   `yaml:"enrichment,omitempty"`
	Rules        []aggregation.Config  `yaml:"rules,omitempty"`
	Destinations []destinations.Config `yaml:"destinations,omitempty"`
	ActiveLists  []activeLists.Config  `yaml:"activeLists,omitempty"`
	Dictionaries []dictionaries.Config `yaml:"dictionaries,omitempty"`
}

func (s *Config) ID() string {
	return s.Name
}

type cmdFlags struct {
	ID          string
	ConfigPath  string
	CoreURL     string
	BusURL      string
	StorageURL  string
	MetricsPort string
	Workers     int
}
