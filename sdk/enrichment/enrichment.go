package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

func Enrich(cfg Config, targetEvent *normalization.Event, sourceEvents ...*normalization.Event) {
	if cfg.TriggerField != "" && cfg.TriggerValue != targetEvent.GetAnyField(cfg.TriggerField) {
		return
	}

	switch cfg.ValueSourceKind {
	case ValueSourceKindDict:
		fromDict(cfg, targetEvent)
	case ValueSourceKindEvent:
		fromEvent(cfg, targetEvent, sourceEvents)
	case ValueSourceKindAL:
		fromAL(cfg, targetEvent)
	default:
		setValue(cfg.Field, cfg.Constant, targetEvent)
	}
}

func fromDict(cfg Config, targetEvent *normalization.Event) {
	srcValue := targetEvent.GetAnyField(cfg.KeyFields[0])
	value := globals.Dictionaries.Get(cfg.ValueSourceName, srcValue)
	setValue(cfg.Field, value, targetEvent)
}

func fromEvent(cfg Config, targetEvent *normalization.Event, sourceEvents []*normalization.Event) {
	var sourceEvent *normalization.Event
	for _, src := range sourceEvents {
		if src.Tag == cfg.ValueSourceName {
			sourceEvent = src
			break
		}
	}

	sourceField := cfg.ValueSourceField
	if sourceField == "" {
		sourceField = cfg.Field
	}

	setValue(sourceField, sourceEvent.GetAnyField(sourceField), targetEvent)
}

func fromAL(cfg Config, targetEvent *normalization.Event) {
	alValue := globals.ActiveLists.Get(cfg.ValueSourceName, cfg.ValueSourceField, cfg.KeyFields, targetEvent)
	setValue(cfg.Field, alValue, targetEvent)
}

func setValue(field string, value interface{}, targetEvent *normalization.Event) {
	switch value.(type) {
	case string:
		targetEvent.SetAnyField(field, value.(string))
	case int64:
		targetEvent.SetIntField(field, value.(int64))
	case float64:
		targetEvent.SetFloatField(field, value.(float64))
	case bool:
		targetEvent.SetBoolField(field, value.(bool))
	}
}
