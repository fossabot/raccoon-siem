package correlator

import (
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/correlation"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
)

type Config struct {
	Name         string                `yaml:"name,omitempty"`
	Connector    connectors.Config     `yaml:"connector,omitempty"`
	Rules        []correlation.Config  `yaml:"rules,omitempty"`
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
