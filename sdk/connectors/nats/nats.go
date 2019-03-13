package nats

import (
	"github.com/nats-io/go-nats"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
)

type Config struct {
	connectors.BaseConfig
	Subject string
	Queue   string
}

type connector struct {
	config Config
}

func (s *connector) Run() error {
	conn, err := nats.Connect(s.config.URL)
	if err != nil {
		return err
	}
	_, err = conn.QueueSubscribe(s.config.Subject, s.config.Queue, s.messageHandler)
	return err
}

func (s *connector) messageHandler(msg *nats.Msg) {
	s.config.OutputChannel <- msg.Data
}

func newNatsConnector(config Config) (*Conn, error) {
	return &connector{config: config}, nil
}
