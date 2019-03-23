package destinations

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"sync"
)

func newConsoleDestination(cfg Config) (IDestination, error) {
	return &consoleDestination{name: cfg.Name}, nil
}

type consoleDestination struct {
	name string
	mu   sync.Mutex
}

func (r *consoleDestination) ID() string {
	return r.name
}

func (r *consoleDestination) Start() error {
	return nil
}

func (r *consoleDestination) Send(event *normalization.Event) {
	r.mu.Lock()
	fmt.Println(event)
	r.mu.Unlock()
}
