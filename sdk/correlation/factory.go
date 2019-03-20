package correlation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type IRule interface {
	ID() string
	Start()
	Stop()
	Feed(event *normalization.Event) bool
}

func NewRule(cfg Config, outputFn OutputFn) (IRule, error) {
	switch cfg.Kind {
	case RuleKindRecovery:
		return newRecoveryRule(cfg, outputFn)
	default:
		return newCommonRule(cfg, outputFn)
	}
}
