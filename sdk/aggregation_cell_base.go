package sdk

import (
	"sync"
	"time"
)

type baseAggregationCell struct {
	mu               sync.RWMutex
	key              string
	aggregatedEvents aggregatedEvents
	timer            *time.Timer
}

func (ac *baseAggregationCell) stopTimer() bool {
	if ac.timer == nil {
		return true
	}
	return ac.timer.Stop()
}
