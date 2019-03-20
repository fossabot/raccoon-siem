package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"time"
)

func Enrich(cfg Config, event *normalization.Event) *normalization.Event {
	if cfg.TriggerField != "" && cfg.TriggerValue != event.GetAnyField(cfg.TriggerField) {
		return event
	}

	switch cfg.ValueSourceKind {
	case FromDict:
		srcValue := event.GetAnyField(cfg.KeyFields[0])
		value := globals.DictionaryStorage.Get(cfg.ValueSourceName, srcValue)
		setValue(cfg.Field, value, event)
	case FromConst:
		setValue(cfg.Field, cfg.Constant, event)
	case FromAL:
	default:
		return event
	}
	return event
}

func setValue(field string, value interface{}, event *normalization.Event) {
	switch value.(type) {
	case string:
		event.SetAnyField(field, value.(string), 0)
	case int64:
		event.SetIntField(field, value.(int64))
	case time.Duration:
		event.SetDurationField(field, value.(time.Duration))
	}
}
