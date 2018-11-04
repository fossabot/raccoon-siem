package sdk

import (
	"fmt"
)

const (
	sourceNATS    = "nats"
	sourceTCP     = "tcp"
	sourceUDP     = "udp"
	sourceNetflow = "netflow"
)

type ISource interface {
	ID() string
	Run() error
}

func NewSource(
	settings *SourceSettings,
	processorChannel chan *ProcessorTask,
) (ISource, error) {
	switch settings.Kind {
	case sourceNATS:
		return newNATSSource(settings, processorChannel), nil
	case sourceTCP:
		return newTCPListenerSource(settings, processorChannel), nil
	case sourceUDP:
		return newUDPListenerSource(settings, processorChannel), nil
	case sourceNetflow:
		return newNetflowSource(settings, processorChannel), nil
	default:
		return nil, fmt.Errorf("unknown source type: %s", settings.Kind)
	}
}
