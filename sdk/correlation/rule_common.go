package correlation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type commonRule struct {
	baseRule
}

func (r *commonRule) onEvent(ar *aggregation.Rule, event *normalization.Event, key string) {
	b := r.addEventToBucket(event, key)

	r.fireTrigger(TriggerEveryEvent, b)
	if b.eventCount == 1 {
		r.fireTrigger(TriggerFirstEvent, b)
	} else {
		r.fireTrigger(TriggerSubsequentEvents, b)
	}

	if r.isThresholdReached(b, RuleKindCommon) {
		b.thresholdCount++
		r.fireTrigger(TriggerEveryThreshold, b)
		if b.thresholdCount == 1 {
			r.fireTrigger(TriggerFirstThreshold, b)
		} else {
			r.fireTrigger(TriggerSubsequentThresholds, b)
		}
	}
}

func (r *commonRule) onTimeout(key string, b *bucket) {
	r.fireTrigger(TriggerTimeout, b)
	r.deleteBucket(key)
}

func newCommonRule(cfg Config, outChannel, correlationChannel chan *normalization.Event) (*commonRule, error) {
	r := &commonRule{}
	base, err := newBaseRule(cfg, r.onEvent, r.onTimeout, outChannel, correlationChannel)
	if err != nil {
		return nil, err
	}
	r.baseRule = base
	return r, nil
}
