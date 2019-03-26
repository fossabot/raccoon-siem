package activeLists

import (
	"context"
	"encoding/json"
	"github.com/nats-io/go-nats"
	"github.com/olivere/elastic"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gopkg.in/vmihailenco/msgpack.v4"
	"io"
	"runtime"
	"time"
)

type persistFn func(alName string, chLog changeLog)

type Container struct {
	cid    string
	lists  map[string]*activeList
	bus    *nats.Conn
	subs   map[string]*nats.Subscription
	esCli  *elastic.Client
	esBulk *elastic.BulkProcessor
}

func (r *Container) Get(listName, field string, keyFields []string, event *normalization.Event) interface{} {
	key := helpers.MakeKey(keyFields, event)
	list := r.lists[listName]
	if list != nil {
		return list.get(key, field)
	}
	return nil
}

func (r *Container) Set(listName string, keyFields []string, mapping []Mapping, event *normalization.Event) {
	key := helpers.MakeKey(keyFields, event)
	list := r.lists[listName]
	if list != nil {
		list.set(key, mapping, event)
	}
}

func (r *Container) Del(listName string, keyFields []string, event *normalization.Event) {
	key := helpers.MakeKey(keyFields, event)
	list := r.lists[listName]
	if list != nil {
		list.del(key)
	}
}

func (r *Container) persistAndBroadcast(alName string, chLog changeLog) {
	chLog.CID = r.cid
	chLog.ALName = alName
	subject := alNamePrefix + alName

	if encodedChangeLog, err := msgpack.Marshal(&chLog); err == nil {
		_ = r.bus.Publish(subject, encodedChangeLog)
	}

	fieldsCopy := make(map[string]interface{})
	for k, v := range chLog.Record.Fields {
		fieldsCopy[k] = v
	}
	chLog.Record.Fields = fieldsCopy

	if chLog.Op == OpSet {
		r.esBulk.Add(elastic.NewBulkIndexRequest().
			Index(subject).
			Type("_doc").
			Id(chLog.Key).
			Version(chLog.Version).
			VersionType("external").
			Doc(chLog.Record).
			UseEasyJSON(true))
	} else if chLog.Op == OpDel {
		r.esBulk.Add(elastic.NewBulkDeleteRequest().
			Index(subject).
			Type("_doc").
			Id(chLog.Key).
			Version(chLog.Version).
			VersionType("external"))
	}
}

func (r *Container) onChangeLog(msg *nats.Msg) {
	chLog := changeLog{}
	if err := msgpack.Unmarshal(msg.Data, &chLog); err != nil {
		return
	}

	if chLog.CID == r.cid {
		return
	}

	if al := r.lists[chLog.ALName]; al != nil {
		al.apply(chLog)
	}
}

func (r *Container) addList(cfg Config) error {
	al := newList(cfg.Name, cfg.TTLDuration().Nanoseconds(), r.persistAndBroadcast)
	r.lists[cfg.Name] = al

	sub, err := r.bus.Subscribe(alNamePrefix+cfg.Name, r.onChangeLog)
	if err != nil {
		return err
	}

	r.subs[cfg.Name] = sub

	if err := r.fetchListRecords(cfg.Name, al); err != nil {
		return err
	}

	return nil
}

func (r *Container) fetchListRecords(name string, list *activeList) error {
	query := elastic.NewBoolQuery().
		MinimumNumberShouldMatch(1).
		Should(
			elastic.NewRangeQuery("ExpiresAt").Gt(time.Now().UnixNano()),
			elastic.NewTermQuery("ExpiresAt", 0),
		)

	scroll := r.esCli.
		Scroll(alNamePrefix + name).
		Type("_doc").
		Query(query).
		Version(true).
		Size(128)

	ctx := context.Background()
	defer scroll.Clear(ctx)

	for {
		res, err := scroll.Do(ctx)
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if res.Hits == nil || res.Hits.Hits == nil || len(res.Hits.Hits) == 0 {
			continue
		}

		for _, hit := range res.Hits.Hits {
			if hit.Version == nil {
				continue
			}

			chLog := changeLog{CID: r.cid, ALName: name, Op: OpSet, Key: hit.Id, Version: *hit.Version}
			if err := json.Unmarshal(*hit.Source, &chLog.Record); err != nil {
				continue
			}

			chLog.Record.Version = chLog.Version
			list.apply(chLog)
		}
	}

	return nil
}

func NewContainer(lists []Config, cid, busURL, storageURL string) (*Container, error) {
	bus, err := nats.Connect(busURL)
	if err != nil {
		return nil, err
	}

	esCli, err := elastic.NewClient(elastic.SetURL(storageURL))
	if err != nil {
		return nil, err
	}

	bulkSvc := elastic.NewBulkProcessorService(esCli).
		BulkActions(-1).
		BulkSize(-1).
		FlushInterval(time.Second).
		Workers(runtime.NumCPU())

	bulkProc, err := bulkSvc.Do(context.Background())
	if err != nil {
		return nil, err
	}

	container := &Container{
		cid:    cid,
		bus:    bus,
		esCli:  esCli,
		esBulk: bulkProc,
		subs:   make(map[string]*nats.Subscription),
		lists:  make(map[string]*activeList),
	}

	for _, listCfg := range lists {
		if err := container.addList(listCfg); err != nil {
			return nil, err
		}
	}

	return container, nil
}
