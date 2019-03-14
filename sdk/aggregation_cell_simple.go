package sdk

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"time"
)

type simpleAggregationCell struct {
	baseAggregationCell
	container *simpleAggregationCells
	threshold int
	sumFields []string
}

func (ac *simpleAggregationCell) put(event *normalization.Event, eventSpec *eventSpecification) {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	if ac.aggregatedEvents.sum(eventSpec.id, ac.sumFields, event, ac.threshold) {
		ac.callTrigger(triggerEveryThreshold, eventSpec, event)
		ac.reset(anySpecID)
	}
}

func (ac *simpleAggregationCell) callTrigger(trigger string, spec *eventSpecification, event *normalization.Event) {
	payload := &triggerPayload{
		eventSpec:  spec,
		eventSpecs: ac.container.eventSpecs,
		cellKey:    ac.key,
		events:     ac.aggregatedEvents.get(spec.id),
	}
	ac.container.triggerFunc(trigger, payload)
}

func (ac *simpleAggregationCell) reset(specID string) {
	ac.stopTimer()
	ac.aggregatedEvents.reset(anySpecID)
	ac.container.delete(ac.key)
}

func (ac *simpleAggregationCell) stopTimer() bool {
	if ac.timer == nil {
		return true
	}
	return ac.timer.Stop()
}

func newSimpleAggregationCell(container *simpleAggregationCells, key string) *simpleAggregationCell {
	ac := new(simpleAggregationCell)
	ac.key = key
	ac.container = container
	ac.aggregatedEvents = newAggregatedEvents()
	ac.threshold = container.eventSpecs[0].threshold
	ac.sumFields = container.sumFields

	if container.window > 0 {
		ac.timer = time.AfterFunc(container.window, func() {
			ac.callTrigger(triggerTimeout, ac.container.eventSpecs[0], nil)
			ac.reset(anySpecID)
		})
	}

	return ac
}
