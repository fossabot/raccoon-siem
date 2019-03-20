package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"strconv"
	"time"
)

func Enrich(cfg Config, event *normalization.Event) *normalization.Event {
	if cfg.TriggerField != "" {
		triggerValue, success := getStringValue(event.GetAnyField(cfg.TriggerField))
		if !success || cfg.TriggerValue != triggerValue {
			return event
		}
	}

	switch cfg.ValueSourceKind {
	case FromDict:
		srcValue := event.GetAnyField(cfg.KeyFields[0])
		key, success := getStringValue(srcValue)
		if !success {
			return event
		}

		value := globals.DictionaryStorage.Get(cfg.ValueSourceName, key)
		event.SetAnyField(cfg.Field, value, 0)
	case FromConst:
		event.SetAnyField(cfg.Field, cfg.Constant, normalization.TimeUnitNone)
	case FromAL:
	default:
		return event
	}
	return event
}

func getStringValue(src interface{}) (string, bool) {
	var key string
	switch src.(type) {
	case string:
		key = src.(string)
	case int64:
		key = strconv.FormatInt(src.(int64), 10)
	case time.Duration:
		key = strconv.FormatInt(src.(time.Duration).Nanoseconds(), 10)
	case time.Time:
		key = strconv.FormatInt(src.(time.Time).UnixNano(), 10)
	default:
		return "", false
	}
	return key, true
}
