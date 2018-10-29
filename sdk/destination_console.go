package sdk

import (
	"fmt"
	"log"
	"sync"
)

func newConsoleDestination(settings DestinationSettings) IDestination {
	return new(consoleDestination)
}

type consoleDestination struct {
	mu sync.Mutex
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
