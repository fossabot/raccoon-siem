package aggregation

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"math"
	"sync"
	"time"
)

type Callback func(event *normalization.Event, hash string)

type bucket struct {
	uniqueHashes map[string]bool
	event        normalization.Event
	eventCount   int
	releaseAt    int64
}

type Rule struct {
	name            string
	filter          *filters.Filter
	outputChannel   chan *normalization.Event
	callback        Callback
	identicalFields []string
	uniqueFields    []string
	sumFields       []string
	sumDelimiter    byte
	threshold       int
	window          time.Duration
	mu              sync.Mutex
	buckets         map[string]*bucket
	ticker          *time.Ticker
}

func (r *Rule) ID() string {
	return r.name
}

func (r *Rule) Start() {
	if r.window > 0 {
		r.ticker = time.NewTicker(time.Second)
		go r.timeoutRoutine()
	}
}

func (r *Rule) Stop() {
	r.mu.Lock()

	if r.ticker != nil {
		r.ticker.Stop()
	}

	for key, bucket := range r.buckets {
		r.releaseBucket(key, bucket)
	}

	r.buckets = nil
	r.mu.Unlock()
}

func (r *Rule) Reset() {
	r.mu.Lock()
	r.buckets = make(map[string]*bucket)
	r.mu.Unlock()
}

func (r *Rule) Feed(event *normalization.Event) bool {
	if !r.filter.Pass(event) {
		return false
	}

	key := event.HashFields(r.identicalFields)
	if r.threshold == 1 {
		event.AggregatedEventCount = 1
		r.sendEvent(event, key)
		return true
	}

	r.mu.Lock()
	b := r.buckets[key]
	isFirstEvent := b == nil

	if isFirstEvent {
		b = &bucket{event: *event, releaseAt: time.Now().Add(r.window).Unix()}
		r.buckets[key] = b
		if len(r.uniqueFields) > 0 {
			b.uniqueHashes = make(map[string]bool)
		}
	}

	if len(r.uniqueFields) > 0 {
		uHash := event.HashFields(r.uniqueFields)
		if b.uniqueHashes[uHash] {
			r.mu.Unlock()
			return false
		}
		b.uniqueHashes[uHash] = true
	}

	if !isFirstEvent && len(r.sumFields) > 0 {
		helpers.SumEvents(&b.event, event, r.sumFields, r.sumDelimiter)
	}

	b.eventCount++
	if r.threshold > 0 && b.eventCount == r.threshold {
		r.releaseBucket(key, b)
	}

	r.mu.Unlock()
	return true
}

func (r *Rule) releaseBucket(key string, bucket *bucket) {
	delete(r.buckets, key)
	bucket.event.AggregatedEventCount = bucket.eventCount
	r.sendEvent(&bucket.event, key)
}

func (r *Rule) sendEvent(event *normalization.Event, key string) {
	event.AggregationRuleName = r.name
	if r.callback != nil {
		r.callback(event, key)
	} else if r.outputChannel != nil {
		r.outputChannel <- event
	}
}

func (r *Rule) timeoutRoutine() {
	var skip int64
	for range r.ticker.C {
		if skip > 0 {
			skip--
			continue
		}

		r.mu.Lock()
		if len(r.buckets) == 0 {
			r.mu.Unlock()
			continue
		}

		now := time.Now().Unix()
		nextReleaseMin := int64(math.MaxInt64)

		for key, bucket := range r.buckets {
			if bucket.releaseAt <= now {
				r.releaseBucket(key, bucket)
			} else if bucket.releaseAt < nextReleaseMin {
				nextReleaseMin = bucket.releaseAt
			}
		}

		r.mu.Unlock()
		skip = nextReleaseMin - now
	}
}

func NewRule(cfg Config, channel chan *normalization.Event, callback Callback) (*Rule, error) {
	r := &Rule{
		name:            cfg.Name,
		threshold:       cfg.Threshold,
		window:          cfg.Window,
		identicalFields: cfg.IdenticalFields,
		uniqueFields:    cfg.UniqueFields,
		sumFields:       cfg.SumFields,
		buckets:         make(map[string]*bucket),
		outputChannel:   channel,
		callback:        callback,
	}

	filter, err := filters.NewFilter(cfg.Filter)
	if err != nil {
		return nil, err
	}

	r.filter = filter

	if cfg.SumDelimiter != "" {
		delimiterBytes := []byte(cfg.SumDelimiter)
		if len(delimiterBytes) > 1 {
			return nil, fmt.Errorf("sum delimiter must be single byte ASCII character: %s", cfg.Name)
		}
		r.sumDelimiter = delimiterBytes[0]
	}

	return r, nil
}
