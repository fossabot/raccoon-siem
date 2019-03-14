package nats

import (
	"github.com/nats-io/go-nats"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
)

type Config struct {
	connectors.BaseConfig
	Subject string
	Queue   string
}

type connector struct {
	cfg Config
}

func (r *connector) ID() string {
	return r.cfg.Name
}

func (r *connector) Run() error {
	conn, err := nats.Connect(r.cfg.URL)
	if err != nil {
		return err
	}
	_, err = conn.QueueSubscribe(r.cfg.Subject, r.cfg.Queue, r.messageHandler)
	return err
}

func (r *connector) messageHandler(msg *nats.Msg) {
	r.cfg.OutputChannel <- connectors.Output{
		Connector: r.cfg.Name,
		Data:      helpers.CopyBytes(msg.Data),
	}
}

func NewConnector(cfg Config) (*connector, error) {
	return &connector{cfg: cfg}, nil
}
