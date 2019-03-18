package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/dictionary"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"strconv"
	"time"
)

const (
	FromConst = "const"
	FromField = "field"
	FromDict  = "dict"
	FromAL    = "al"
)

type EnrichConfig struct {
	// Поле event, куда записываем результат
	TargetField string `yaml:"targetField,omitempty"`
	KeyField    string `yaml:"keyField,omitempty"`
	From        string `yaml:"from,omitempty"`
	FromKey     string `yaml:"fromKey,omitempty"`
	Const       string `yaml:"const,omitempty"`
}

func Enrich(cfg EnrichConfig, event *normalization.Event) *normalization.Event {
	switch cfg.From {
	case FromField:
		srcValue := event.GetAnyField(cfg.KeyField)
		switch srcValue.(type) {
		case string:
			event.SetAnyField(cfg.TargetField, srcValue.(string), normalization.TimeUnitNone)
		case int64:
			event.SetAnyField(cfg.TargetField, strconv.FormatInt(srcValue.(int64), 10), normalization.TimeUnitNone)
		case time.Duration:
			duration := srcValue.(time.Duration)
			event.SetAnyField(cfg.TargetField, strconv.FormatInt(duration.Nanoseconds(), 10), normalization.TimeUnitNone)
		case time.Time:
			t := srcValue.(time.Time)
			event.SetAnyField(cfg.TargetField, strconv.FormatInt(t.UnixNano(), 10), normalization.TimeUnitNone)
		default:
			return event
		}
	case FromConst:
		event.SetAnyField(cfg.TargetField, cfg.Const, normalization.TimeUnitNone)
	case FromDict:
		value := dictionary.MockDictionary.Get(cfg.FromKey, cfg.KeyField)
		event.SetAnyField(cfg.TargetField, value, 0)
	default:
		return event
	}
	return event
}
