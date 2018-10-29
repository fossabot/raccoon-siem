package sdk

import (
	"fmt"
)

const (
	destinationConsole       = "console"
	destinationNATS          = "nats"
	destinationElasticsearch = "elasticsearch"
)

var knownDestinations = map[string]bool{
	destinationConsole:       true,
	destinationNATS:          true,
	destinationElasticsearch: true,
}

type IDestination interface {
	Run() error
	Send(event *Event)
}

func NewDestination(settings DestinationSettings) (IDestination, error) {
	switch settings.Kind {
	case destinationConsole:
		return newConsoleDestination(settings), nil
	case destinationNATS:
		return newNATSDestination(settings), nil
	case destinationElasticsearch:
		return newElasticsearchDestination(settings), nil
	default:
		return nil, fmt.Errorf("unknown destination type: %s", settings.Kind)
	}
}
