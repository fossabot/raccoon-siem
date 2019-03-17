package filters

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type Filter struct {
	comparator
	name     string
	not      bool
	sections []SectionConfig
}

func (f *Filter) ID() string {
	return f.name
}

func (f *Filter) Pass(event normalization.Event) bool {
	for _, section := range f.sections {
		if !f.checkSection(event, section) {
			return f.not
		}
	}
	return !f.not
}

func (f *Filter) checkSection(event normalization.Event, section SectionConfig) bool {
	for _, cond := range section.Conditions {
		if !section.Or {
			if !f.conditionMatch(event, cond) {
				return section.Not
			}
		} else {
			if f.conditionMatch(event, cond) {
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

func (f *Filter) conditionMatch(event normalization.Event, cond ConditionConfig) bool {
	lv := event.GetAnyField(cond.Field)
	switch cond.RvSource {
	case RvSourceField:
		return f.compareValues(lv, event.GetAnyField(cond.Rv.(string)), cond.Op)
	case RvSourceDict:
		// TODO: ask dictionary for value
		return false
	case RvSourceAL:
		// TODO: ask active list for value
		return false
	default:
		return f.compareValues(lv, cond.Rv, cond.Op)
	}
}

func NewFilter(cfg Config) (*Filter, error) {
	return &Filter{
		name:     cfg.Name,
		not:      cfg.Not,
		sections: cfg.Sections,
	}, nil
}
