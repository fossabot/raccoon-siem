package connectors

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"time"
)

type kafkaConnector struct {
	name    string
	reader  *kafka.Reader
	channel OutputChannel
}

func (r *kafkaConnector) ID() string {
	return r.name
}

func (r *kafkaConnector) Start() error {
	if err := r.reader.SetOffset(kafka.LastOffset); err != nil {
		return err
	}
	go r.worker()
	return nil
}

func (r *kafkaConnector) worker() {
	for {
		if m, err := r.reader.ReadMessage(context.Background()); err == nil {
			r.channel <- Output{
				Connector: r.name,
				Data:      helpers.CopyBytes(m.Value),
			}
		}
	}
}

func newKafkaConnector(cfg Config, channel OutputChannel) (*kafkaConnector, error) {
	return &kafkaConnector{
		name:    cfg.Name,
		channel: channel,
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{cfg.URL},
			Topic:          cfg.Subject,
			GroupID:        cfg.Queue,
			Partition:      cfg.Partition,
			CommitInterval: time.Second,
		}),
	}, nil
}
