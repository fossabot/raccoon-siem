package connectors

import (
	"fmt"
)

const (
	connectorListener = "listener"
	connectorNetflow  = "netflow"
	connectorNats     = "nats"
)

type IConnector interface {
	ID() string
	Start() error
}

type Output struct {
	Connector string
	Data      []byte
}

type OutputChannel chan Output

func New(cfg Config, channel OutputChannel) (IConnector, error) {
	switch cfg.Kind {
	case connectorListener:
		return newListenerConnector(cfg, channel)
	case connectorNats:
		return newNATSConnector(cfg, channel)
	case connectorNetflow:
		return newNetflowConnector(cfg, channel)
	default:
		return nil, fmt.Errorf("unknown connector kind: %s", cfg.Kind)
	}
}
