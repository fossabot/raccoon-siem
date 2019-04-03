package connectors

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"log"
	"time"
)

type kafkaConnector struct {
	name    string
	reader  *kafka.Reader
	channel OutputChannel
	debug   bool
}

func (r *kafkaConnector) ID() string {
	return r.name
}

func (r *kafkaConnector) Start() error {
	go r.worker()
	return nil
}

func (r *kafkaConnector) worker() {
	for {
		m, err := r.reader.ReadMessage(context.Background())
		if err != nil {
			if r.debug {
				log.Println(err)
			}
			continue
		}

		if r.debug {
			log.Println(string(m.Value))
		}

		r.channel <- Output{
			Connector: r.name,
			Data:      helpers.CopyBytes(m.Value),
		}
	}
}

func newKafkaConnector(cfg Config, channel OutputChannel) (*kafkaConnector, error) {
	testConn, err := kafka.Dial("tcp", cfg.URL)
	if err != nil {
		return nil, err
	}
	_ = testConn.Close()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{cfg.URL},
		Topic:          cfg.Subject,
		GroupID:        cfg.Queue,
		Partition:      cfg.Partition,
		CommitInterval: time.Second,
		MinBytes:       4096,
		MaxBytes:       10e6,
	})

	if err := reader.SetOffset(kafka.LastOffset); err != nil {
		return nil, err
	}

	return &kafkaConnector{
		name:    cfg.Name,
		debug:   cfg.Debug,
		reader:  reader,
		channel: channel,
	}, nil
}
