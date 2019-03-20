package destinations

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"runtime"
	"strings"
	"sync"
	"time"
)

func newElasticDestination(cfg Config) (*elasticDestination, error) {
	return &elasticDestination{
		name:        cfg.Name,
		url:         cfg.URL,
		inChannel:   make(chan *normalization.Event),
		timeLayout:  "2006-01-02",
		index:       cfg.Index,
		staticIndex: cfg.StaticIndex,
	}, nil
}

type elasticDestination struct {
	mu            sync.Mutex
	name          string
	url           string
	index         string
	staticIndex   bool
	connection    *elastic.Client
	bulkProcessor *elastic.BulkProcessor
	inChannel     chan *normalization.Event
	timeLayout    string
}

func (r *elasticDestination) ID() string {
	return r.name
}

func (r *elasticDestination) Run() error {
	conn, bp, err := r.createConnection()

	if err != nil {
		return err
	}

	r.connection = conn
	r.bulkProcessor = bp
	r.spawnWorker()

	return nil
}

func (r *elasticDestination) Send(event *normalization.Event) {
	r.inChannel <- event
}

func (r *elasticDestination) createConnection() (*elastic.Client, *elastic.BulkProcessor, error) {
	conn, err := elastic.NewClient(elastic.SetURL(r.url))

	bps := elastic.NewBulkProcessorService(conn).
		BulkActions(-1).
		BulkSize(-1).
		FlushInterval(time.Second).
		Workers(runtime.NumCPU()).
		After(func(_ int64, _ []elastic.BulkableRequest, resp *elastic.BulkResponse, err error) {
			if err != nil {
				return
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

func (r *elasticDestination) spawnWorker() {
	go func() {
		for event := range r.inChannel {
			ts := time.Now()
			eventsToSend := []*normalization.Event{event}
			eventsToSend = append(eventsToSend, event.BaseEvents...)
			for _, eventToSend := range eventsToSend {
				request := elastic.NewBulkIndexRequest().
					Index(r.makeFinalIndexName(ts)).
					Type("_doc").
					Doc(eventToSend)
				r.bulkProcessor.Add(request)
			}
		}
	}()
}

func (r *elasticDestination) makeFinalIndexName(ts time.Time) string {
	if !r.staticIndex {
		sb := strings.Builder{}
		sb.WriteString(r.index)
		sb.WriteString("-")
		sb.WriteString(time.Now().Format(r.timeLayout))
		return sb.String()
	}
	return r.index
}
