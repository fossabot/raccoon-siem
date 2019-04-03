package destinations

import (
	"fmt"
)

type IDestination interface {
	ID() string
	Start() error
	Send(data []byte)
}

func New(cfg Config) (IDestination, error) {
	switch cfg.Kind {
	case DestinationNATS:
		return newNATS(cfg)
	case DestinationElastic:
		return newElastic(cfg)
	default:
		return nil, fmt.Errorf("unknown destination kind: %s", cfg.Kind)
	}
}
