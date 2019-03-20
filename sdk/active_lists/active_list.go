package activeLists

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"strings"
	"sync"
)

type Value map[string]interface{}

type ActiveList struct {
	mu   sync.RWMutex
	data sync.Map
}

func (r *ActiveList) Set(keyFields []string, fields []string, sourceEvent *normalization.Event) {
	sb := strings.Builder{}
	for _, keyField := range keyFields {
		sourceEvent.
	}

	v := make(Value)
	for _, field := range fields {
		v[field] = sourceEvent.GetAnyField(field)
	}
	r.mu.Lock()
	r.data[key] = v
	r.mu.Unlock()
}

func New() *ActiveList {
	return &ActiveList{data: make(map[string]Value)}
}
