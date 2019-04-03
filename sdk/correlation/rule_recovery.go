package correlation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type recoveryRule struct {
	baseRule
}

func (r *recoveryRule) onEvent(selector eventSelector, _ *normalization.Event, _ *bucket, key string) {
	if selector.recovery {
		r.deleteBucket(key)
		return
	}
}

func (r *recoveryRule) onTimeout(b *bucket, key string) {
	if r.isThresholdReached(b, RuleKindRecovery) {
		r.fireTrigger(TriggerTimeout, b)
	}
	r.deleteBucket(key)
}

func newRecoveryRule(cfg Config, outputFn OutputFn) (*recoveryRule, error) {
	r := &recoveryRule{}
	base, err := newBaseRule(cfg, r.onEvent, r.onTimeout, outputFn)
	if err != nil {
		return nil, err
	}
	r.baseRule = base
	return r, nil
}
