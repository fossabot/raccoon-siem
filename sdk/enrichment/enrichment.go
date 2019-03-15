package enrichment

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

const (
	FromConst = "const"
	FromDict  = "dict"
	FromAL    = "al"
)

type EnrichConfig struct {
	TargetField string `yaml:"targetField,omitempty"`
	KeyField    string `yaml:"keyField,omitempty"`
	From        string `yaml:"from,omitempty"`
	Const       string `const:"const,omitempty"`
}

func Enrich(cfg EnrichConfig, event *normalization.Event) *normalization.Event {
	switch cfg.From {
	case FromConst:
		event.Set(cfg.TargetField, cfg.Const, normalization.TimeUnitNone)
	case FromDict:

	}
	return event
}
