package sdk

import (
	"fmt"
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
	Send(event *Event)
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
