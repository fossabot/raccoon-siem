package filters

import (
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type Filter struct {
	comparator
	name     string
	not      bool
	sections []SectionConfig
}

func (r *Filter) ID() string {
	return r.name
}

func (r *Filter) Pass(event *normalization.Event) bool {
	for _, section := range r.sections {
		if !r.checkSection(event, section) {
			return r.not
		}
	}
	return !r.not
}

func (r *Filter) checkSection(event *normalization.Event, section SectionConfig) bool {
	for _, cond := range section.Conditions {
		if !section.Or {
			if !r.conditionMatch(event, cond) {
				return section.Not
			}
		} else {
			if r.conditionMatch(event, cond) {
				return !section.Not
			}
		}
	}

	if !section.Or {
		return !section.Not
	} else {
		return section.Not
	}
}

func (r *Filter) conditionMatch(event *normalization.Event, cond ConditionConfig) bool {
	lv := event.GetAnyField(cond.Field)
	switch cond.CMPSourceKind {
	case ValueSourceKindEvent:
		return r.compareValues(lv, event.GetAnyField(cond.CMPSourceField), cond.Op)
	case ValueSourceKindDict:
		key := event.GetAnyField(cond.KeyFields[0])
		rv := globals.Dictionaries.Get(cond.CMPSourceName, key)
		return r.compareValues(lv, rv, cond.Op)
	case ValueSourceKindAL:
		alValue := globals.ActiveLists.Get(cond.CMPSourceName, cond.CMPSourceField, cond.KeyFields, event)
		return r.compareValues(lv, alValue, cond.Op)
	default:
		return r.compareValues(lv, cond.Constant, cond.Op)
	}
}

func NewFilter(cfg Config) (*Filter, error) {
	return &Filter{
		name:     cfg.Name,
		not:      cfg.Not,
		sections: cfg.Sections,
	}, nil
}
