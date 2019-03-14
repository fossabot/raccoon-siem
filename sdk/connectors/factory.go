package connectors

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors/listener"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors/nats"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors/netflow"
)

const (
	connectorListener = "listener"
	connectorNetflow  = "netflow"
	connectorNats     = "nats"
)

type IConnector interface {
	ID() string
	Run() error
}

func NewConnector(config UserConfig, channel OutputChannel) (IConnector, error) {
	switch config.Kind {
	case connectorListener:
		return listener.NewConnector(listener.Config{
			BaseConfig: BaseConfig{
				Name:          config.Name,
				URL:           config.URL,
				OutputChannel: channel,
			},
			Protocol:   config.Protocol,
			Delimiter:  config.Delimiter,
			BufferSize: config.BufferSize,
		})
	case connectorNats:
		return nats.NewConnector(nats.Config{
			BaseConfig: BaseConfig{
				Name:          config.Name,
				URL:           config.URL,
				OutputChannel: channel,
			},
			Subject: config.Subject,
			Queue:   config.Queue,
		})
	case connectorNetflow:
		return netflow.NewConnector(netflow.Config{
			BaseConfig: BaseConfig{
				Name:          config.Name,
				URL:           config.URL,
				OutputChannel: channel,
			},
			BufferSize: config.BufferSize,
		})
	default:
		return nil, fmt.Errorf("unknown connector kind: %s", config.Kind)
	}
}
