package enrichment

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

const (
	FromField = "field"
	FromConst = "const"
	FromDict  = "dict"
	FromAL    = "al"
)

type EnrichConfig struct {
	TargetField string `yaml:"targetField,omitempty"`
	Key         string `yaml:"key,omitempty"`
	From        string `yaml:"from,omitempty"`
	Const		string `const:"from,omitempty"`
}

func Enrich(cfg EnrichConfig, event *normalization.Event) *normalization.Event {
	switch cfg.From {
	case FromField:
		//value := event.Get(cfg.Key)
	case FromConst:
		event.Set(cfg.TargetField, []byte(cfg.Const), 0)
	case FromDict:
		return event
	default:
		return event
	}
}
