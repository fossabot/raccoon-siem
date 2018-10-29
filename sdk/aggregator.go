package sdk

import (
	"time"
)

type aggregator struct {
	baseAggregator
	aggregationCells     *aggregationCells
	rootAggregationCells *aggregationCells
	uniqueFields         []string
}

// Processes partial eventSpecs
func (a *aggregator) feed(event *Event, eventSpec *eventSpecification) {
	cellKey := event.HashFields(a.identicalFields)
	cell, exists := a.aggregationCells.get(cellKey)

	if !exists {
		cell = a.aggregationCells.create(cellKey)
	}

	uniqueKey := event.HashFields(a.uniqueFields)
	cell.Put(event, uniqueKey, eventSpec)
}

// Processes correlated eventSpecs
func (a *aggregator) feedRoot(event *Event, eventSpec *eventSpecification, cellKey string) {
	cell, exists := a.rootAggregationCells.get(cellKey)

	if !exists {
		cell = a.rootAggregationCells.create(cellKey)
	}

	cell.Put(event, defaultEventFieldsHash, eventSpec)
}

// Resets aggregator state
func (a *aggregator) reset() {
	a.aggregationCells.reset()
	a.rootAggregationCells.reset()
}

func newAggregator(
	settings *aggregation,
	eventSpecs []*eventSpecification,
	triggerFunc triggerCallback,
	rootTriggerFunc triggerCallback,
) *aggregator {
	window := time.Duration(settings.window) * time.Second

	a := new(aggregator)
	a.identicalFields = settings.identicalFields
	a.uniqueFields = settings.uniqueFields

	a.rootAggregationCells = newAggregationCells(
		true,
		settings.threshold,
		eventSpecs,
		window,
		rootTriggerFunc,
	)

	a.aggregationCells = newAggregationCells(
		false,
		-1,
		eventSpecs,
		window,
		triggerFunc,
	)

	return a
}
