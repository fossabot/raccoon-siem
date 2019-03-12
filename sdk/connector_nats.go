package sdk

import (
	"github.com/nats-io/go-nats"
)

type NatsConnectorConfig struct {
	BaseConnectorConfig
	Subject string
	Queue   string
}

type natsConnector struct {
	config NatsConnectorConfig
}

func (s *natsConnector) ID() string {
	return s.config.Name
}

func (s *natsConnector) Run() error {
	conn, err := nats.Connect(s.config.URL)
	if err != nil {
		return err
	}
	_, err = conn.QueueSubscribe(s.config.Subject, s.config.Queue, s.messageHandler)
	return err
}

func (s *natsConnector) messageHandler(msg *nats.Msg) {
	s.config.OutputChannel <- &ProcessorTask{
		Connector: s.config.Name,
		Data:      msg.Data,
	}
}

func newNatsConnector(config NatsConnectorConfig) (IConnector, error) {
	return &natsConnector{config: config}, nil
}
