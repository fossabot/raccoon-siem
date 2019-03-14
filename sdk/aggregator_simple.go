package sdk

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"time"
)

type simpleAggregator struct {
	baseAggregator
	aggregationCells *simpleAggregationCells
}

func (a *simpleAggregator) feed(event *normalization.Event, eventSpec *eventSpecification) {
	cellKey := event.HashFields(a.identicalFields)
	cell, exists := a.aggregationCells.get(cellKey)

	if !exists {
		cell = a.aggregationCells.create(cellKey)
	}

	cell.put(event, eventSpec)
}

func (a *simpleAggregator) feedRoot(event *normalization.Event, eventSpec *eventSpecification, cellKey string) {
	panic("not implemented")
}

func (a *simpleAggregator) reset() {
	a.aggregationCells.reset()
}

func newSimpleAggregator(
	settings *aggregation,
	eventSpecs []*eventSpecification,
	triggerFunc triggerCallback,
) *simpleAggregator {
	window := time.Duration(settings.window) * time.Second

	a := new(simpleAggregator)
	a.identicalFields = settings.identicalFields

	a.aggregationCells = newSimpleAggregationCells(
		eventSpecs,
		window,
		triggerFunc,
		a.sumFields,
	)

	return a
}
