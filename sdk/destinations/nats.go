package destinations

import (
	"github.com/nats-io/go-nats"
	"sync"
)

func newNATS(cfg Config) (*natsDestination, error) {
	return &natsDestination{
		name:    cfg.Name,
		url:     cfg.URL,
		subject: cfg.Subject,
	}, nil
}

type natsDestination struct {
	mu         sync.Mutex
	name       string
	url        string
	subject    string
	connection *nats.Conn
}

func (r *natsDestination) ID() string {
	return r.name
}

func (r *natsDestination) Start() error {
	conn, err := nats.Connect(r.url, nats.MaxReconnects(-1))
	if err != nil {
		return err
	}
	r.connection = conn
	return nil
}

func (r *natsDestination) Send(data []byte) {
	_ = r.connection.Publish(r.subject, data)
}
