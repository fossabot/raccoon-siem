package filters

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type filter struct {
	comparator
	name     string
	not      bool
	sections []SectionConfig
}

func (f *filter) ID() string {
	return f.name
}

func (f *filter) Pass(events ...*normalization.Event) bool {
	for _, section := range f.sections {
		if !f.checkSection(events[0], section) {
			return f.not
		}
	}
	return !f.not
}

func (f *filter) checkSection(event *normalization.Event, section SectionConfig) bool {
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

func (f *filter) conditionMatch(event *normalization.Event, cond ConditionConfig) bool {
	lv := event.Get(cond.Field)
	switch cond.RvSource {
	case RvSourceField:
		return f.compareValues(lv, event.Get(cond.Rv.(string)), cond.Op)
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

func newFilter(cfg Config) (*filter, error) {
	return &filter{
		name:     cfg.Name,
		not:      cfg.Not,
		sections: cfg.Sections,
	}, nil
}
