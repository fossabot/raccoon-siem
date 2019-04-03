package destinations

import (
	"context"
	"github.com/olivere/elastic"
	"runtime"
	"strings"
	"sync"
	"time"
)

func newElastic(cfg Config) (*elasticDestination, error) {
	return &elasticDestination{
		name:        cfg.Name,
		url:         cfg.URL,
		timeLayout:  "2006-01-02",
		index:       cfg.Index,
		staticIndex: cfg.StaticIndex,
	}, nil
}

type elasticDestination struct {
	mu          sync.Mutex
	name        string
	url         string
	index       string
	staticIndex bool
	esBulk      *elastic.BulkProcessor
	timeLayout  string
}

func (r *elasticDestination) ID() string {
	return r.name
}

func (r *elasticDestination) Start() error {
	esCli, err := elastic.NewClient(elastic.SetURL(r.url))
	if err != nil {
		return err
	}

	bulkSvc := elastic.NewBulkProcessorService(esCli).
		BulkActions(-1).
		BulkSize(-1).
		FlushInterval(time.Second).
		Workers(runtime.NumCPU())

	bp, err := bulkSvc.Do(context.Background())
	if err != nil {
		return err
	}

	r.esBulk = bp
	return nil
}

func (r *elasticDestination) Send(data []byte) {
	r.esBulk.Add(elastic.NewBulkIndexRequest().
		Index(r.makeFinalIndexName()).
		Type("_doc").
		Doc(string(data)))
}

func (r *elasticDestination) makeFinalIndexName() string {
	if !r.staticIndex {
		sb := strings.Builder{}
		sb.WriteString(r.index)
		sb.WriteString("-")
		sb.WriteString(time.Now().Format(r.timeLayout))
		return sb.String()
	}
	return r.index
}
