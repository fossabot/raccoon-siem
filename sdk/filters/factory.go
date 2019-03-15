package filters

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

type IFilter interface {
	ID() string
	Pass(events ...*normalization.Event) bool
}

func New(cfg Config) (IFilter, error) {
	if cfg.Join {
		return newJoinFilter(cfg)
	} else {
		return newFilter(cfg)
	}
}
