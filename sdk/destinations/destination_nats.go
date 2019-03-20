package destinations

import (
	"github.com/nats-io/go-nats"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"runtime"
	"sync"
)

func newNATSDestination(cfg Config) (*natsDestination, error) {
	return &natsDestination{
		name:      cfg.Name,
		url:       cfg.URL,
		subject:   cfg.Subject,
		inChannel: make(chan *normalization.Event, runtime.NumCPU()),
	}, nil
}

type natsDestination struct {
	mu         sync.Mutex
	name       string
	url        string
	subject    string
	connection *nats.Conn
	inChannel  chan *normalization.Event
}

func (r *natsDestination) ID() string {
	return r.name
}

func (r *natsDestination) Run() error {
	conn, err := nats.Connect(r.url, nats.MaxReconnects(-1))

	if err != nil {
		return err
	}

	r.connection = conn

	for i := 0; i < runtime.NumCPU(); i++ {
		r.spawnWorker()
	}

	return nil
}

func (r *natsDestination) Send(event *normalization.Event) {
	r.inChannel <- event
}

func (r *natsDestination) spawnWorker() {
	go func() {
		for event := range r.inChannel {
			data, err := event.ToMsgPack()
			if err == nil {
				_ = r.connection.Publish(r.subject, data)
			}
		}
	}()
}
