package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/dictionary"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"strconv"
	"time"
)

const (
	FromField = "field"
	FromConst = "const"
	FromDict  = "dict"
	FromAL    = "al"
)

type EnrichConfig struct {
	// Поле event, куда записываем результат
	TargetField string `yaml:"targetField,omitempty"`
	Key         string `yaml:"key,omitempty"`
	From        string `yaml:"from,omitempty"`
	FromKey     string `yaml:"fromKey,omitempty"`
	Const       string `yaml:"const,omitempty"`
}

func Enrich(cfg EnrichConfig, event *normalization.Event) *normalization.Event {
	switch cfg.From {
	case FromField:
		srcValue := event.Get(cfg.Key)
		switch srcValue.(type) {
		case string:
			event.Set(cfg.TargetField, []byte(srcValue.(string)), 0)
		case int64:
			event.Set(cfg.TargetField, []byte(strconv.FormatInt(srcValue.(int64), 10)), 0)
		case time.Duration:
			duration := srcValue.(time.Duration)
			event.Set(cfg.TargetField, []byte(strconv.FormatInt(duration.Nanoseconds(), 10)), 0)
		case time.Time:
			t := srcValue.(time.Time)
			event.Set(cfg.TargetField, []byte(strconv.FormatInt(t.UnixNano(), 10)), 0)
		default:
			return event
		}
	case FromConst:
		event.Set(cfg.TargetField, []byte(cfg.Const), 0)
	case FromDict:
		value := dictionary.MockDictionary.Get(cfg.FromKey, cfg.Key)
		event.Set(cfg.TargetField, []byte(value), 0)
	default:
		return event
	}
	return event
}
