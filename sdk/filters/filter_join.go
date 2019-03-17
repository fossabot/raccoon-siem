package filters

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type JoinFilter struct {
	comparator
	not      bool
	name     string
	sections []JoinSectionConfig
}

func (f *JoinFilter) ID() string {
	return f.name
}

func (f *JoinFilter) Pass(events ...normalization.Event) bool {
	eventTags := make(map[string]normalization.Event)
	for _, event := range events {
		eventTags[event.AggregationRuleName] = event
	}

	for _, section := range f.sections {
		if !f.checkSection(eventTags, section) {
			return f.not
		}
	}

	return !f.not
}

func (f *JoinFilter) checkSection(eventsByTag map[string]normalization.Event, section JoinSectionConfig) bool {
	for _, cond := range section.Conditions {
		srcEvent := eventsByTag[cond.LeftTag]
		dstEvent := eventsByTag[cond.RightTag]

		if !section.Or {
			if !f.joinConditionMatch(srcEvent, cond.LeftField, dstEvent, cond.RightField, cond.Op) {
				return section.Not
			}
		} else {
			if f.joinConditionMatch(srcEvent, cond.LeftField, dstEvent, cond.RightField, cond.Op) {
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

func (f *JoinFilter) joinConditionMatch(
	srcEvent normalization.Event,
	srcEventField string,
	dstEvent normalization.Event,
	dstEventField string,
	op string,
) bool {
	return f.compareValues(
		srcEvent.GetAnyField(srcEventField),
		dstEvent.GetAnyField(dstEventField),
		op,
	)
}

func NewJoinFilter(cfg JoinConfig) (*JoinFilter, error) {
	return &JoinFilter{
		name:     cfg.Name,
		not:      cfg.Not,
		sections: cfg.JoinSections,
	}, nil
}
