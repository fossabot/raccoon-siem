package sdk

import (
	"github.com/nats-io/go-nats"
	"runtime"
	"sync"
)

func newNATSDestination(settings DestinationSettings) IDestination {
	return &natsDestination{
		settings:  settings,
		inChannel: make(chan *Event, runtime.NumCPU()),
	}
}

type natsDestination struct {
	mu         sync.Mutex
	settings   DestinationSettings
	connection *nats.Conn
	inChannel  chan *Event
}

func (d *natsDestination) Run() error {
	conn, err := nats.Connect(
		d.settings.URL,
		nats.MaxReconnects(-1),
	)

	if err != nil {
		return err
	}

	d.connection = conn

	for i := 0; i < runtime.NumCPU(); i++ {
		d.spawnWorker()
	}

	return nil
}

func (d *natsDestination) Send(event *Event) {
	d.inChannel <- event
}

func (d *natsDestination) spawnWorker() {
	go func() {
		for event := range d.inChannel {
			data, err := event.ToBSON()
			if err == nil {
				d.connection.Publish(d.settings.Channel, data)
			}
		}
	}()
}
