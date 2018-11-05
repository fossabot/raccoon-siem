package sdk

import (
	"fmt"
	"log"
	"sync"
)

func newConsoleDestination(settings DestinationSettings) IDestination {
	return &consoleDestination{settings: settings}
}

type consoleDestination struct {
	mu       sync.Mutex
	settings DestinationSettings
}

func (d *consoleDestination) ID() string {
	return d.settings.Name
}

func (d *consoleDestination) Run() error {
	return nil
}

func (d *consoleDestination) Send(event *Event) {
	d.mu.Lock()
	fmt.Println(event)
	if event.Trace != "" {
		log.Println(event.Trace)
	}
	d.mu.Unlock()
}
