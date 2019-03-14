package sdk

import (
	"github.com/nats-io/go-nats"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"runtime"
	"sync"
)

func newNATSDestination(settings DestinationSettings) IDestination {
	return &natsDestination{
		settings:  settings,
		inChannel: make(chan *normalization.Event, runtime.NumCPU()),
	}
}

type natsDestination struct {
	mu         sync.Mutex
	settings   DestinationSettings
	connection *nats.Conn
	inChannel  chan *normalization.Event
}

func (d *natsDestination) ID() string {
	return d.settings.Name
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

func (d *natsDestination) Send(event *normalization.Event) {
	d.inChannel <- event
}

func (d *natsDestination) spawnWorker() {
	go func() {
		for event := range d.inChannel {
			data, err := event.ToMsgPack()
			if err == nil {
				d.connection.Publish(d.settings.Channel, data)
			}
		}
	}()
}
