package correlation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type commonRule struct {
	baseRule
}

func (r *commonRule) onEvent(selector eventSelector, event *normalization.Event, b *bucket, _ string) {
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

func (r *commonRule) onTimeout(b *bucket, key string) {
	r.fireTrigger(TriggerTimeout, b)
	r.deleteBucket(key)
}

func newCommonRule(cfg Config, outputFn OutputFn) (*commonRule, error) {
	r := &commonRule{}
	base, err := newBaseRule(cfg, r.onEvent, r.onTimeout, outputFn)
	if err != nil {
		return nil, err
	}
	r.baseRule = base
	return r, nil
}
