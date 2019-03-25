package actions

import (
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
)

const (
	KindRelease    = "release"
	KindActiveList = "al"
)

type ReleaseConfig struct {
	EnrichmentConfigs []enrichment.Config `yaml:"enrichment,omitempty"`
}

type ActiveListConfig struct {
	Name      string                `yaml:"name,omitempty"`
	Op        string                `yaml:"op,omitempty"`
	KeyFields []string              `yaml:"keyFields,omitempty"`
	Mapping   []activeLists.Mapping `yaml:"mapping,omitempty"`
	EventTag  string                `yaml:"eventTag,omitempty"`
}
