package correlation

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"math"
	"sync"
	"time"
)

type bucket struct {
	thresholdCount int
	subBuckets     map[string]*subBucket
	resetAt        int64
}

type subBucket struct {
	event      normalization.Event
	eventCount int
}

type Rule struct {
	name             string
	aggregationRules []*aggregation.Rule
	filter           *filters.JoinFilter
	actions          map[string][]ActionConfig
	window           time.Duration
	resetWindowOn    string
	mu               sync.Mutex
	buckets          map[string]*bucket
	ticker           *time.Ticker
	outputChannel    chan normalization.Event
}

func (r *Rule) Start() {
	if r.window > 0 {
		r.ticker = time.NewTicker(time.Second)
		go r.timeoutRoutine()
	}

	for _, aggRule := range r.aggregationRules {
		aggRule.Start()
	}
}

func (r *Rule) Stop() {
	r.mu.Lock()

	if r.ticker != nil {
		r.ticker.Stop()
	}

	for _, aggRule := range r.aggregationRules {
		aggRule.Stop()
	}

	r.mu.Unlock()
}

func (r *Rule) Feed(event *normalization.Event) {
	if event.CorrelationRuleName != r.name {
		for _, aggRule := range r.aggregationRules {
			if aggRule.Feed(event) {
				break
			}
		}
	}
}

func (r *Rule) eventReady(caller *aggregation.Rule, event *normalization.Event, hash string) {
	r.mu.Lock()

	if caller.IsUnexpected() {
		r.resetBucket(hash)
		r.mu.Unlock()
		return
	}

	b := r.buckets[hash]
	if b == nil {
		b = &bucket{
			resetAt:    time.Now().Add(r.window).Unix(),
			subBuckets: make(map[string]*subBucket),
		}
		r.buckets[hash] = b
	}

	sb := b.subBuckets[event.AggregationRuleName]
	if sb == nil {
		sb = &subBucket{}
		b.subBuckets[event.AggregationRuleName] = sb
	}

	sb.event = *event
	sb.eventCount++

	if sb.eventCount == 1 {
		r.fireTrigger(TriggerFirstEvent, b, sb)
	} else {
		r.fireTrigger(TriggerSubsequentEvents, b, sb)
	}

	r.fireTrigger(TriggerEveryEvent, b, sb)

	thresholdReached := true
	eventsForFilter := make([]*normalization.Event, 0, len(r.aggregationRules))
	for _, aggRule := range r.aggregationRules {
		if !aggRule.IsUnexpected() {
			sb, ok := b.subBuckets[aggRule.ID()]
			if !ok {
				thresholdReached = false
				break
			}
			eventsForFilter = append(eventsForFilter, &sb.event)
		}
	}

	if thresholdReached {
		if r.filter != nil && !r.filter.Pass(eventsForFilter...) {
			r.resetBucket(hash)
			r.mu.Unlock()
			return
		}

		b.thresholdCount++
		if b.thresholdCount == 1 {
			r.fireTrigger(TriggerFirstThreshold, b, nil)
		} else {
			r.fireTrigger(TriggerSubsequentThresholds, b, nil)
		}

		r.fireTrigger(TriggerEveryThreshold, b, nil)
	}

	r.mu.Unlock()
}

func (r *Rule) releaseEvent(event *normalization.Event) {
	if r.outputChannel != nil {
		r.outputChannel <- *event
	}
}

func (r *Rule) resetBucket(key string) {
	delete(r.buckets, key)
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
		nextResetMin := int64(math.MaxInt64)

		for key, bucket := range r.buckets {
			if bucket.resetAt <= now {
				r.fireTrigger(TriggerTimeout, bucket, nil)
				r.resetBucket(key)
			} else if bucket.resetAt < nextResetMin {
				nextResetMin = bucket.resetAt
			}
		}

		r.mu.Unlock()
		skip = nextResetMin - now
	}
}

func (r *Rule) fireTrigger(kind string, bucket *bucket, subBucket *subBucket) {
	if r.resetWindowOn == kind {
		bucket.resetAt = time.Now().Add(r.window).Unix()
	}
	if subBucket != nil {
		fmt.Printf("rule: %s, trigger: %s, srcIP: %s, count: %d\n",
			r.name, kind, subBucket.event.SourceIPAddress, subBucket.event.AggregatedEventCount)
	} else {
		fmt.Printf("rule: %s, trigger: %s\n", r.name, kind)
	}
}

func NewRule(cfg Config, channel chan normalization.Event) (*Rule, error) {
	r := &Rule{
		name:          cfg.Name,
		window:        cfg.Window,
		buckets:       make(map[string]*bucket),
		actions:       make(map[string][]ActionConfig),
		outputChannel: channel,
	}

	if cfg.Filter != nil {
		filter, err := filters.NewJoinFilter(*cfg.Filter)
		if err != nil {
			return nil, err
		}
		r.filter = filter
	}

	for _, aggRuleCfg := range cfg.AggregationRules {
		aggRule, err := aggregation.NewRule(aggRuleCfg, nil, r.eventReady)
		if err != nil {
			return nil, err
		}
		r.aggregationRules = append(r.aggregationRules, aggRule)
	}

	for _, trigger := range cfg.Triggers {
		r.actions[trigger.Kind] = trigger.Actions
	}

	return r, nil
}
