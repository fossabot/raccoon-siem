package correlation

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type recoveryRule struct {
	baseRule
}

func (r *recoveryRule) onEvent(ar *aggregation.Rule, event *normalization.Event, key string) {
	if ar.IsRecovery() {
		r.deleteBucket(key)
		return
	}
	r.addEventToBucket(event, key)
}

func (r *recoveryRule) onTimeout(key string, b *bucket) {
	if r.isThresholdReached(b, RuleKindRecovery) {
		r.fireTrigger(TriggerTimeout, b)
	}
	r.deleteBucket(key)
}

func newRecoveryRule(cfg Config, outChannel, correlationChannel chan *normalization.Event) (*recoveryRule, error) {
	r := &recoveryRule{}
	base, err := newBaseRule(cfg, r.onEvent, r.onTimeout, outChannel, correlationChannel)
	if err != nil {
		return nil, err
	}

	hasRecoveryAggregationRule := false
	for _, ar := range base.aggregationRules {
		if ar.IsRecovery() {
			hasRecoveryAggregationRule = true
			break
		}
	}

	if !hasRecoveryAggregationRule {
		return nil, fmt.Errorf("%s: at least one recovery aggregation rule required", base.name)
	}

	r.baseRule = base
	return r, nil
}
