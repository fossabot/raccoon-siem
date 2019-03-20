package aggregation

import (
	"github.com/satori/go.uuid"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"math"
	"sync"
	"time"
)

type Callback func(caller *Rule, event *normalization.Event, key string)

type bucket struct {
	uniqueHashes map[string]bool
	event        normalization.Event
	eventCount   int
	releaseAt    int64
}

type Rule struct {
	name            string
	filter          *filters.Filter
	identicalFields []string
	uniqueFields    []string
	sumFields       []string
	threshold       int
	window          time.Duration
	recovery        bool
	mu              sync.Mutex
	buckets         map[string]*bucket
	ticker          *time.Ticker
	outputFn        OutputFn
}

func (r *Rule) ID() string {
	return r.name
}

func (r *Rule) IsRecovery() bool {
	return r.recovery
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
	if event.AggregationRuleName == r.name {
		return false
	}

	if !r.filter.Pass(event) {
		return false
	}

	key := event.HashFields(r.identicalFields)

	r.mu.Lock()
	b := r.buckets[key]
	isFirstEvent := b == nil

	if isFirstEvent {
		b = &bucket{
			event:        *event,
			releaseAt:    time.Now().Add(r.window).Unix(),
			uniqueHashes: make(map[string]bool),
		}
		r.buckets[key] = b
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
		helpers.SumEvents(&b.event, event, r.sumFields)
	}

	b.eventCount++
	if r.threshold > 0 && b.eventCount == r.threshold {
		r.releaseBucket(key, b)
	}

	r.mu.Unlock()
	return true
}

func (r *Rule) releaseBucket(key string, bucket *bucket) {
	bucket.event.AggregatedEventCount = bucket.eventCount
	r.sendEvent(&bucket.event, key)
	delete(r.buckets, key)
}

func (r *Rule) sendEvent(event *normalization.Event, key string) {
	event.ID = uuid.NewV4().String()
	event.AggregationRuleName = r.name
	if r.outputFn != nil {
		r.outputFn(event)
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

func NewRule(cfg Config, outputFn OutputFn) (*Rule, error) {
	r := &Rule{
		name:            cfg.Name,
		threshold:       cfg.Threshold,
		window:          cfg.Window,
		recovery:        cfg.Recovery,
		identicalFields: cfg.IdenticalFields,
		uniqueFields:    cfg.UniqueFields,
		sumFields:       cfg.SumFields,
		buckets:         make(map[string]*bucket),
		outputFn:        outputFn,
	}

	filter, err := filters.NewFilter(cfg.Filter)
	if err != nil {
		return nil, err
	}

	r.filter = filter
	return r, nil
}
