package sdk

import (
	"sync"
	"time"
)

type baseAggregationCells struct {
	mu          sync.RWMutex
	eventSpecs  []*eventSpecification
	window      time.Duration
	triggerFunc triggerCallback
}
