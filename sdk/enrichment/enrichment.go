package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
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
		setValueFromInterface(cfg.Field, cfg.Constant, targetEvent)
	}
}

func fromDict(cfg Config, targetEvent *normalization.Event) {
	key := helpers.MakeKey(cfg.KeyFields, targetEvent)
	targetEvent.SetAnyField(cfg.Field, globals.Dictionaries.Get(cfg.ValueSourceName, key))
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

	setValueFromInterface(sourceField, sourceEvent.GetAnyField(sourceField), targetEvent)
}

func fromAL(cfg Config, targetEvent *normalization.Event) {
	alValue := globals.ActiveLists.Get(cfg.ValueSourceName, cfg.ValueSourceField, cfg.KeyFields, targetEvent)
	setValueFromInterface(cfg.Field, alValue, targetEvent)
}

func setValueFromInterface(field string, value interface{}, targetEvent *normalization.Event) {
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
