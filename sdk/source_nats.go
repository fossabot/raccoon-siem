package sdk

import (
	"github.com/nats-io/go-nats"
)

func newNATSSource(
	settings *SourceSettings,
	processorChannel chan *ProcessorTask,
) ISource {
	return &natsSource{
		processorChannel: processorChannel,
		settings:         settings,
	}
}

type natsSource struct {
	settings         *SourceSettings
	connection       *nats.Conn
	subscription     *nats.Subscription
	processorChannel chan *ProcessorTask
}

func (s *natsSource) ID() string {
	return s.settings.Name
}

func (s *natsSource) Run() error {
	conn, err := nats.Connect(s.settings.URL)

	if err != nil {
		return err
	}

	s.connection = conn

	if !s.settings.LoadBalance {
		if err := s.subscribe(); err != nil {
			return err
		}
	} else {
		if err := s.subscribeQueue(s.settings.Queue); err != nil {
			return err
		}
	}

	return nil
}

func (s *natsSource) subscribeQueue(queueName string) error {
	sub, err := s.connection.QueueSubscribe(s.settings.Channel, queueName, s.messageHandler)

	if err != nil {
		return err
	}

	s.subscription = sub
	return nil
}

func (s *natsSource) subscribe() error {
	sub, err := s.connection.Subscribe(s.settings.Channel, s.messageHandler)

	if err != nil {
		return err
	}

	s.subscription = sub
	return nil
}

func (s *natsSource) messageHandler(msg *nats.Msg) {
	s.processorChannel <- &ProcessorTask{
		Source: s.settings.Name,
		Data:   msg.Data,
	}
}
