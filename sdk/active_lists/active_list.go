package activeLists

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"sync"
	"time"
)

type activeList struct {
	mu        sync.RWMutex
	name      string
	records   map[string]record
	ttl       int64
	expTree   expirationTree
	expTicker *time.Ticker
	persistFn persistFn
}

func (r *activeList) get(key, field string) (v string) {
	r.mu.RLock()
	rec := r.records[key]
	if rec.Fields != nil {
		v = rec.Fields[field]
	}
	r.mu.RUnlock()
	return
}

func (r *activeList) set(key string, mapping []Mapping, event *normalization.Event) {
	r.mu.Lock()
	chLog := changeLog{Op: OpSet, Key: key, Version: time.Now().UnixNano()}

	rec := r.records[key]
	rec.Version = chLog.Version

	if r.ttl > 0 {
		rec.ExpiresAt = chLog.Version + r.ttl
	}

	if rec.Fields == nil {
		rec.Fields = make(map[string]string)
	}

	changes := 0
	for _, m := range mapping {
		currentValue := rec.Fields[m.ALField]
		var newValue string

		if m.Constant != nil {
			newValue = normalization.ToString(m.Constant)
		} else {
			newValue = normalization.ToString(event.GetAnyField(m.EventField))
		}

		if currentValue != newValue {
			rec.Fields[m.ALField] = newValue
			changes++
		}
	}

	r.expTree.touch(key, rec.ExpiresAt)

	if changes == 0 {
		r.mu.Unlock()
		return
	}

	r.records[key] = rec
	chLog.Record = rec
	r.persistFn(r.name, chLog)

	r.mu.Unlock()
}

func (r *activeList) del(key string) {
	r.mu.Lock()
	chLog := changeLog{Op: OpDel, Key: key, Version: time.Now().UnixNano()}
	delete(r.records, key)
	r.expTree.del(key)
	r.persistFn(r.name, chLog)
	r.mu.Unlock()
}

func (r *activeList) apply(chLog changeLog) {
	r.mu.Lock()

	rec := r.records[chLog.Key]
	if rec.Version > chLog.Version {
		r.mu.Unlock()
		return
	}

	if chLog.Op == OpDel {
		delete(r.records, chLog.Key)
		r.expTree.del(chLog.Key)
		r.mu.Unlock()
		return
	}

	r.records[chLog.Key] = chLog.Record
	r.expTree.touch(chLog.Key, chLog.Record.ExpiresAt)

	r.mu.Unlock()
}

func (r *activeList) expirationRoutine() {
	for range r.expTicker.C {
		if keysToExpire := r.expTree.getExpiredKeys(); len(keysToExpire) > 0 {
			r.mu.Lock()
			for _, key := range keysToExpire {
				delete(r.records, key)
				r.expTree.del(key)
			}
			r.mu.Unlock()
		}
	}
}

func newList(name string, ttl int64, persistFn persistFn) *activeList {
	al := &activeList{
		name:      name,
		records:   make(map[string]record),
		ttl:       ttl,
		expTree:   createExpirationTree(),
		expTicker: time.NewTicker(time.Second),
		persistFn: persistFn,
	}
	go al.expirationRoutine()
	return al
}
