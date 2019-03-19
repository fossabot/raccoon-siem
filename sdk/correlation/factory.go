package correlation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

const (
	RuleKindCommon   = "common"
	RuleKindRecovery = "recover"
)

type IRule interface {
	ID() string
	Start()
	Stop()
	Feed(event *normalization.Event) bool
}

func NewRule(cfg Config, outChannel, correlationChannel chan *normalization.Event) (IRule, error) {
	switch cfg.Kind {
	case RuleKindRecovery:
		return newRecoveryRule(cfg, outChannel, correlationChannel)
	default:
		return newCommonRule(cfg, outChannel, correlationChannel)
	}
}
