package sdk

import (
	"fmt"
	"sync"
	"time"
)

func NewMeter(label string, avgPeriod int) *Meter {
	return &Meter{
		label:     label,
		avgPeriod: avgPeriod,
	}
}

type Meter struct {
	mu        sync.Mutex
	label     string
	count     int
	recentEPS []int
	avgPeriod int
}

func (m *Meter) Run() *Meter {
	if m.avgPeriod <= 0 {
		m.avgPeriod = 60
	}

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				m.report()
			}
		}
	}()

	return m
}

func (m *Meter) Feed() {
	m.mu.Lock()
	m.count++
	m.mu.Unlock()
}

func (m *Meter) report() {
	m.mu.Lock()

	m.recentEPS = append(m.recentEPS, m.count)
	m.count = 0

	if len(m.recentEPS) == m.avgPeriod {
		sum := 0

		for _, cnt := range m.recentEPS {
			sum += cnt
		}

		fmt.Printf("[%s] Average EPS: %d\n", m.label, sum/len(m.recentEPS))
		m.recentEPS = nil
	}

	m.mu.Unlock()
}
