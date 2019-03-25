package connectors

import (
	"github.com/nats-io/go-nats"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
)

type natsConnector struct {
	name    string
	url     string
	subject string
	queue   string
	channel OutputChannel
}

func (r *natsConnector) ID() string {
	return r.name
}

func (r *natsConnector) Start() error {
	conn, err := nats.Connect(r.url)
	if err != nil {
		return err
	}
	_, err = conn.QueueSubscribe(r.subject, r.queue, r.messageHandler)
	return err
}

func (r *natsConnector) messageHandler(msg *nats.Msg) {
	r.channel <- Output{
		Connector: r.name,
		Data:      helpers.CopyBytes(msg.Data),
	}
}

func newNATSConnector(cfg Config, channel OutputChannel) (*natsConnector, error) {
	return &natsConnector{
		name:    cfg.Name,
		url:     cfg.URL,
		subject: cfg.Subject,
		queue:   cfg.Queue,
		channel: channel,
	}, nil
}
