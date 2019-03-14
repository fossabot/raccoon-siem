package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

const (
	DestinationConsole       = "console"
	DestinationNATS          = "nats"
	DestinationElasticsearch = "elasticsearch"
)

var knownDestinations = map[string]bool{
	DestinationConsole:       true,
	DestinationNATS:          true,
	DestinationElasticsearch: true,
}

type IDestination interface {
	ID() string
	Run() error
	Send(event *normalization.Event)
}

func NewDestination(settings DestinationSettings) (IDestination, error) {
	switch settings.Kind {
	case DestinationConsole:
		return newConsoleDestination(settings), nil
	case DestinationNATS:
		return newNATSDestination(settings), nil
	case DestinationElasticsearch:
		return newElasticsearchDestination(settings), nil
	default:
		return nil, fmt.Errorf("unknown destination type: %s", settings.Kind)
	}
}
