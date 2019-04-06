package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/mutation"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

func Enrich(cfg Config, targetEvent *normalization.Event, sourceEvents ...*normalization.Event) {
	if cfg.filter != nil && !cfg.filter.Pass(targetEvent) {
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

	value := globals.Dictionaries.Get(cfg.ValueSourceName, key)
	if cfg.Mutation != nil {
		value = mutation.Mutate(cfg.Mutation, value)
	}

	targetEvent.SetAnyField(cfg.Field, value)
}

func fromEvent(cfg Config, targetEvent *normalization.Event, sourceEvents []*normalization.Event) {
	var sourceEvent *normalization.Event
	for _, src := range sourceEvents {
		if src.Tag == cfg.ValueSourceName {
			sourceEvent = src
			break
		}
	}

	if sourceEvent == nil {
		sourceEvent = targetEvent
	}

	sourceField := cfg.ValueSourceField
	if sourceField == "" {
		sourceField = cfg.Field
	}

	value := normalization.ToString(sourceEvent.GetAnyField(sourceField))
	if cfg.Mutation != nil {
		value = mutation.Mutate(cfg.Mutation, value)
	}
	targetEvent.SetAnyField(cfg.Field, value)
}

func fromAL(cfg Config, targetEvent *normalization.Event) {
	value := globals.ActiveLists.Get(cfg.ValueSourceName, cfg.ValueSourceField, cfg.KeyFields, targetEvent)
	if cfg.Mutation != nil {
		value = mutation.Mutate(cfg.Mutation, value)
	}
	targetEvent.SetAnyField(cfg.Field, value)
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
