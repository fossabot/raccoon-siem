package activeLists

import (
	"context"
	"github.com/nats-io/go-nats"
	"github.com/olivere/elastic"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gopkg.in/vmihailenco/msgpack.v4"
	"runtime"
	"time"
)

type persistFn func(alName string, chLog changeLog)

type Container struct {
	cid    string
	lists  map[string]*activeList
	bus    *nats.Conn
	subs   map[string]*nats.Subscription
	esBulk *elastic.BulkProcessor
}

func (r *Container) Get(listName, field string, keyFields []string, event *normalization.Event) interface{} {
	key := makeKey(keyFields, event)
	list := r.lists[listName]
	if list != nil {
		return list.get(key, field)
	}
	return nil
}

func (r *Container) Set(listName string, keyFields []string, mapping []Mapping, event *normalization.Event) {
	key := makeKey(keyFields, event)
	list := r.lists[listName]
	if list != nil {
		list.set(key, mapping, event)
	}
}

func (r *Container) Del(listName string, keyFields []string, event *normalization.Event) {
	key := makeKey(keyFields, event)
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
	al := newList(cfg.Name, cfg.TTL.Nanoseconds(), r.persistAndBroadcast)
	r.lists[cfg.Name] = al

	sub, err := r.bus.Subscribe(alNamePrefix+cfg.Name, r.onChangeLog)
	if err != nil {
		return err
	}

	r.subs[cfg.Name] = sub
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
