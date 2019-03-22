package correlation

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"math"
	"sync"
	"time"
)

type timeoutCallback func(bucket *bucket, key string)
type eventCallback func(selector eventSelector, event *normalization.Event, bucket *bucket, key string)

type bucket struct {
	thresholdCount  int
	eventCount      int
	eventCountByTag map[string]int
	uniqueHashes    map[string]bool
	timeoutAt       int64
	events          []*normalization.Event
}

type eventSelector struct {
	tag       string
	filter    *filters.Filter
	threshold int
	recovery  bool
}

type baseRule struct {
	name            string
	selectors       []eventSelector
	filter          *filters.JoinFilter
	identicalFields []string
	uniqueFields    []string
	actions         map[string][]ActionConfig
	window          time.Duration
	mu              sync.Mutex
	buckets         map[string]*bucket
	ticker          *time.Ticker
	onEvent         eventCallback
	onTimeout       timeoutCallback
	outputFn        OutputFn
}

func (r *baseRule) ID() string {
	return r.name
}

func (r *baseRule) Start() {
	if r.window > 0 {
		r.ticker = time.NewTicker(time.Second)
		go r.timeoutRoutine()
	}
}

func (r *baseRule) Stop() {
	r.mu.Lock()
	if r.ticker != nil {
		r.ticker.Stop()
	}
	r.mu.Unlock()
}

func (r *baseRule) Feed(event *normalization.Event) bool {
	if event.CorrelationRuleName != r.name {
		for _, selector := range r.selectors {
			if selector.filter.Pass(event) {
				eventClone := event.Clone()
				eventClone.Tag = selector.tag
				return r.aggregate(selector, &eventClone)
			}
		}
	}
	return false
}

func (r baseRule) aggregate(selector eventSelector, event *normalization.Event) bool {
	key := event.HashFields(r.identicalFields)
	r.mu.Lock()

	b := r.buckets[key]
	if b == nil {
		b = &bucket{
			timeoutAt:       time.Now().Add(r.window).Unix(),
			eventCountByTag: make(map[string]int),
			uniqueHashes:    make(map[string]bool),
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

	b.events = append(b.events, event)
	b.eventCount++
	b.eventCountByTag[event.Tag]++

	r.onEvent(selector, event, b, key)

	r.mu.Unlock()
	return true
}

func (r *baseRule) isThresholdReached(b *bucket, callerKind string) bool {
	for _, selector := range r.selectors {
		if callerKind == RuleKindRecovery && selector.recovery {
			continue
		}

		if b.eventCountByTag[selector.tag] < selector.threshold {
			return false
		}
	}

	if r.filter != nil && !r.filter.Pass(b.events...) {
		return false
	}

	return true
}

func (r *baseRule) fireTrigger(kind string, b *bucket) {
	for _, action := range r.actions[kind] {
		switch action.Kind {
		case actions.KindRelease:
			correlationEvent := r.createCorrelationEvent(b)
			for _, cfg := range action.Release.EnrichmentConfigs {
				enrichment.Enrich(cfg, correlationEvent, b.events...)
			}
			if r.outputFn != nil {
				r.outputFn(correlationEvent)
			}
		}
	}
}

func (r *baseRule) createCorrelationEvent(b *bucket) *normalization.Event {
	newEvent := &normalization.Event{
		ID:                  uuid.NewV4().String(),
		Correlated:          true,
		Timestamp:           helpers.NowUnixMillis(),
		CorrelationRuleName: r.name,
		BaseEventCount:      len(b.events),
	}

	if len(b.events) > 0 {
		helpers.CopyFields(newEvent, b.events[0], r.identicalFields)
		for _, base := range b.events {
			newEvent.BaseEventIDs = append(newEvent.BaseEventIDs, base.ID)
		}
	}

	return newEvent
}

func (r *baseRule) deleteBucket(key string) {
	delete(r.buckets, key)
}

func (r *baseRule) resetBucketTimeout(b *bucket) {
	b.timeoutAt = time.Now().Add(r.window).Unix()
}

func (r *baseRule) timeoutRoutine() {
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
		closestTimeout := int64(math.MaxInt64)

		for key, bucket := range r.buckets {
			if bucket.timeoutAt <= now {
				r.onTimeout(bucket, key)
			} else if bucket.timeoutAt < closestTimeout {
				closestTimeout = bucket.timeoutAt
			}
		}

		r.mu.Unlock()
		skip = closestTimeout - now
	}
}

func newBaseRule(
	cfg Config,
	onEvent eventCallback,
	onTimeout timeoutCallback,
	outputFn OutputFn,
) (baseRule, error) {
	r := baseRule{
		name:            cfg.Name,
		window:          cfg.Window,
		identicalFields: cfg.IdenticalFields,
		uniqueFields:    cfg.UniqueFields,
		buckets:         make(map[string]*bucket),
		actions:         make(map[string][]ActionConfig),
		onEvent:         onEvent,
		onTimeout:       onTimeout,
		outputFn:        outputFn,
	}

	if cfg.Filter != nil {
		filter, err := filters.NewJoinFilter(*cfg.Filter)
		if err != nil {
			return r, err
		}
		r.filter = filter
	}

	for _, selectorCfg := range cfg.Selectors {
		selector := eventSelector{
			tag:       selectorCfg.Tag,
			threshold: selectorCfg.Threshold,
		}

		filter, err := filters.NewFilter(selectorCfg.Filter)
		if err != nil {
			return r, err
		}

		selector.filter = filter
		r.selectors = append(r.selectors, selector)
	}

	if len(r.selectors) == 0 {
		return r, fmt.Errorf("%s: at least one event selector required", r.name)
	}

	for kind, trigger := range cfg.Triggers {
		r.actions[kind] = trigger.Actions
	}

	return r, nil
}
