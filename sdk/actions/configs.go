package actions

import "github.com/tephrocactus/raccoon-siem/sdk/enrichment"

const (
	KindRelease    = "release"
	KindActiveList = "al"
)

type ReleaseConfig struct {
	EnrichmentConfigs []enrichment.Config `yaml:"enrichment,omitempty"`
}
