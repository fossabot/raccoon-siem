package sdk

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"runtime"
	"strings"
	"sync"
	"time"
)

func newElasticsearchDestination(settings DestinationSettings) IDestination {
	return &elasticsearchDestination{
		settings:   settings,
		inChannel:  make(chan *normalization.Event),
		timeLayout: "2006-01-02",
	}
}

type elasticsearchDestination struct {
	mu            sync.Mutex
	settings      DestinationSettings
	connection    *elastic.Client
	bulkProcessor *elastic.BulkProcessor
	inChannel     chan *normalization.Event
	timeLayout    string
}

func (d *elasticsearchDestination) ID() string {
	return d.settings.Name
}

func (d *elasticsearchDestination) Run() error {
	conn, bp, err := d.createConnection()

	if err != nil {
		return err
	}

	d.connection = conn
	d.bulkProcessor = bp
	d.spawnWorker()

	return nil
}

func (d *elasticsearchDestination) Send(event *normalization.Event) {
	d.inChannel <- event
}

func (d *elasticsearchDestination) createConnection() (*elastic.Client, *elastic.BulkProcessor, error) {
	conn, err := elastic.NewClient(elastic.SetURL(d.settings.URL))

	bps := elastic.NewBulkProcessorService(conn).
		BulkActions(-1).
		BulkSize(-1).
		FlushInterval(time.Second).
		Workers(runtime.NumCPU()).
		After(func(_ int64, _ []elastic.BulkableRequest, resp *elastic.BulkResponse, err error) {
			if err != nil {
				if Debug {
					fmt.Println(err)
				}
				return
			}

			if Debug {
				for _, failed := range resp.Failed() {
					fmt.Println(failed.Error.Reason)
				}
			}
		})

	if err != nil {
		return nil, nil, err
	}

	ctx := context.Background()
	bp, err := bps.Do(ctx)

	if err != nil {
		return nil, nil, err
	}

	err = bp.Start(ctx)

	if err != nil {
		return nil, nil, err
	}

	return conn, bp, nil
}

func (d *elasticsearchDestination) spawnWorker() {
	go func() {
		for event := range d.inChannel {
			ts := time.Now()
			eventsToSend := []*normalization.Event{event}
			eventsToSend = append(eventsToSend, event.BaseEvents...)
			for _, eventToSend := range eventsToSend {
				request := elastic.NewBulkIndexRequest().
					Index(d.makeFinalIndexName(ts)).
					Type("_doc").
					Doc(eventToSend)
				d.bulkProcessor.Add(request)
			}
		}
	}()
}

func (d *elasticsearchDestination) makeFinalIndexName(ts time.Time) string {
	if d.settings.RotateIndex {
		sb := strings.Builder{}
		sb.WriteString(d.settings.Index)
		sb.WriteString("-")
		sb.WriteString(time.Now().Format(d.timeLayout))
		return sb.String()
	}
	return d.settings.Index
}
