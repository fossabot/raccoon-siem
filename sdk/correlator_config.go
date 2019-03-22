package sdk

import (
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/correlation"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
)

type CorrelatorConfig struct {
	Name         string                `yaml:"name,omitempty"`
	Rules        []correlation.Config  `yaml:"rules,omitempty"`
	Destinations []destinations.Config `yaml:"destinations,omitempty"`
	ActiveLists  []activeLists.Config  `yaml:"activeLists,omitempty"`
}

func (s *CorrelatorConfig) ID() string {
	return s.Name
}
