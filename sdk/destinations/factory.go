package destinations

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type IDestination interface {
	ID() string
	Start() error
	Send(event *normalization.Event)
}

func New(cfg Config) (IDestination, error) {
	switch cfg.Kind {
	case DestinationConsole:
		return newConsoleDestination(cfg)
	case DestinationNATS:
		return newNATS(cfg)
	case DestinationElastic:
		return newElastic(cfg)
	default:
		return nil, fmt.Errorf("unknown destination kind: %s", cfg.Kind)
	}
}
