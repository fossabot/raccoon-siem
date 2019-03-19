package correlation

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"math"
	"sync"
	"time"
)

type timeoutCallback func(key string, bucket *bucket)

type bucket struct {
	thresholdCount  int
	eventCount      int
	eventCountByTag map[string]int
	timeoutAt       int64
	events          []*normalization.Event
}

type baseRule struct {
	name               string
	aggregationRules   []*aggregation.Rule
	filter             *filters.JoinFilter
	actions            map[string][]ActionConfig
	window             time.Duration
	mu                 sync.Mutex
	buckets            map[string]*bucket
	ticker             *time.Ticker
	outputChannel      chan *normalization.Event
	correlationChannel chan *normalization.Event
	onEvent            aggregation.Callback
	onTimeout          timeoutCallback
}

func (r *baseRule) ID() string {
	return r.name
}

func (r *baseRule) Start() {
	if r.window > 0 {
		r.ticker = time.NewTicker(time.Second)
		go r.timeoutRoutine()
	}

	for _, aggRule := range r.aggregationRules {
		aggRule.Start()
	}
}

func (r *baseRule) Stop() {
	r.mu.Lock()

	if r.ticker != nil {
		r.ticker.Stop()
	}

	for _, aggRule := range r.aggregationRules {
		aggRule.Stop()
	}

	r.mu.Unlock()
}

func (r *baseRule) Feed(event *normalization.Event) bool {
	if event.CorrelationRuleName != r.name {
		for _, aggRule := range r.aggregationRules {
			if aggRule.Feed(event) {
				return true
			}
		}
	}
	return false
}

func (r *baseRule) onEventWrapper(ar *aggregation.Rule, event *normalization.Event, key string) {
	r.mu.Lock()
	r.onEvent(ar, event, key)
	r.mu.Unlock()
}

func (r *baseRule) addEventToBucket(event *normalization.Event, key string) *bucket {
	b := r.buckets[key]
	if b == nil {
		b = &bucket{timeoutAt: time.Now().Add(r.window).Unix(), eventCountByTag: make(map[string]int)}
		r.buckets[key] = b
	}
	b.events = append(b.events, event)
	b.eventCount++
	b.eventCountByTag[event.AggregationRuleName]++
	return b
}

func (r *baseRule) isThresholdReached(b *bucket, callerKind string) bool {
	for _, aggRule := range r.aggregationRules {
		if callerKind == RuleKindRecovery && aggRule.IsRecovery() {
			continue
		}

		if b.eventCountByTag[aggRule.ID()] == 0 {
			return false
		}
	}

	if r.filter != nil && !r.filter.Pass(b.events...) {
		return false
	}

	return true
}

func (r *baseRule) fireTrigger(kind string, b *bucket) {
	fmt.Printf("rule: %s, trigger: %s\n", r.name, kind)

	for _, action := range r.actions[kind] {
		switch action.Kind {
		case actions.KindRelease:
			correlationEvent := r.createCorrelationEvent(b)
			actions.Release(action.Release, correlationEvent, b.events)
			r.toOutputChannel(correlationEvent)
			r.toCorrelationChannel(correlationEvent)
		}
	}
}

func (r *baseRule) createCorrelationEvent(b *bucket) *normalization.Event {
	newEvent := &normalization.Event{
		ID:                  uuid.NewV4().String(),
		Correlated:          true,
		Timestamp:           time.Now(),
		CorrelationRuleName: r.name,
		BaseEventCount:      len(b.events),
	}

	for _, base := range b.events {
		newEvent.BaseEventIDs = append(newEvent.BaseEventIDs, base.ID)
	}

	return newEvent
}

func (r *baseRule) toOutputChannel(events ...*normalization.Event) {
	if r.outputChannel != nil {
		for _, event := range events {
			r.outputChannel <- event
		}
	}
}

func (r *baseRule) toCorrelationChannel(events ...*normalization.Event) {
	if r.correlationChannel != nil {
		for _, event := range events {
			r.correlationChannel <- event
		}
	}
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
				r.onTimeout(key, bucket)
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
	onEvent aggregation.Callback,
	onTimeout timeoutCallback,
	outChannel chan *normalization.Event,
	correlationChannel chan *normalization.Event,
) (baseRule, error) {
	r := baseRule{
		name:               cfg.Name,
		window:             cfg.Window,
		buckets:            make(map[string]*bucket),
		actions:            make(map[string][]ActionConfig),
		outputChannel:      outChannel,
		correlationChannel: correlationChannel,
		onEvent:            onEvent,
		onTimeout:          onTimeout,
	}

	if cfg.Filter != nil {
		filter, err := filters.NewJoinFilter(*cfg.Filter)
		if err != nil {
			return r, err
		}
		r.filter = filter
	}

	for _, aggRuleCfg := range cfg.AggregationRules {
		aggRule, err := aggregation.NewRule(aggRuleCfg, nil, r.onEventWrapper)
		if err != nil {
			return r, err
		}
		r.aggregationRules = append(r.aggregationRules, aggRule)
	}

	for kind, trigger := range cfg.Triggers {
		r.actions[kind] = trigger.Actions
	}

	return r, nil
}
