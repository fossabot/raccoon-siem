package activeLists

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"time"
)

const (
	OpGet        = "get"
	OpSet        = "set"
	OpDel        = "del"
	AlNamePrefix = "raccoon-al-"
)

type Config struct {
	Name string `json:"name"`
	TTL  int64  `json:"ttl"`
}

func (r *Config) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("active list: name required")
	}

	if r.TTL < 0 {
		return fmt.Errorf("active list: ttl must be positive or zero")
	}

	return nil
}

func (r *Config) TTLDuration() time.Duration {
	return time.Duration(r.TTL) * time.Second
}

type Mapping struct {
	EventField string      `json:"eventField"`
	ALField    string      `json:"activeListField"`
	Constant   interface{} `json:"constant"`
}

func (r *Mapping) Validate() error {
	if r.Constant == nil && !helpers.EventFieldHasGetter(r.EventField) {
		return fmt.Errorf("mapping: invalid event field %s", r.EventField)
	}
	return nil
}

type ScanCallback func(key string, value []byte)

type IStorage interface {
	Put(list, key string, data map[string]interface{}, ttl time.Duration) error
	Del(list, key string) error
	Get(list, key, field string) ([]byte, error)
}
