package sdk

import (
	"fmt"
)

const (
	connectorListener = "listener"
	connectorNetflow  = "netflow"
	connectorNats     = "nats"
)

type BaseConnectorConfig struct {
	Name          string
	URL           string
	OutputChannel chan *ProcessorTask
}

type IConnector interface {
	ID() string
	Run() error
}

func NewConnector(config Config, processorChannel chan *ProcessorTask) (IConnector, error) {
	switch config.Kind {
	case connectorListener:
		return newListenerConnector(ListenerConnectorConfig{
			BaseConnectorConfig: BaseConnectorConfig{
				Name:          config.Name,
				URL:           config.URL,
				OutputChannel: processorChannel,
			},
			Protocol:   config.Protocol,
			Delimiter:  config.Delimiter,
			BufferSize: config.BufferSize,
		})
	case connectorNats:
		return newNatsConnector(NatsConnectorConfig{
			BaseConnectorConfig: BaseConnectorConfig{
				Name:          config.Name,
				URL:           config.URL,
				OutputChannel: processorChannel,
			},
			Subject: config.Subject,
			Queue:   config.Queue,
		})
	case connectorNetflow:
		return newNetflowConnector(NetflowConnectorConfig{
			BaseConnectorConfig: BaseConnectorConfig{
				Name:          config.Name,
				URL:           config.URL,
				OutputChannel: processorChannel,
			},
			BufferSize: config.BufferSize,
		})
	default:
		return nil, fmt.Errorf("unknown source type: %s", config.Kind)
	}
}
