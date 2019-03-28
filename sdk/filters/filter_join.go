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

func (r *JoinFilter) ID() string {
	return r.name
}

func (r *JoinFilter) Pass(events ...*normalization.Event) bool {
	eventTags := make(map[string]*normalization.Event)
	for _, event := range events {
		eventTags[event.AggregationRuleName] = event
	}

	for _, section := range r.sections {
		if !r.checkSection(eventTags, section) {
			return r.not
		}
	}

	return !r.not
}

func (r *JoinFilter) checkSection(eventsByTag map[string]*normalization.Event, section JoinSectionConfig) bool {
	for _, cond := range section.Conditions {
		srcEvent := eventsByTag[cond.LeftTag]
		if srcEvent == nil {
			return section.Not
		}

		dstEvent := eventsByTag[cond.RightTag]
		if dstEvent == nil {
			return section.Not
		}

		if !section.Or {
			if !r.joinConditionMatch(srcEvent, cond.LeftField, dstEvent, cond.RightField, cond.Op) {
				return section.Not
			}
		} else {
			if r.joinConditionMatch(srcEvent, cond.LeftField, dstEvent, cond.RightField, cond.Op) {
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

func (r *JoinFilter) joinConditionMatch(
	srcEvent *normalization.Event,
	srcEventField string,
	dstEvent *normalization.Event,
	dstEventField string,
	op string,
) bool {
	return r.compareValues(
		srcEvent.GetAnyField(srcEventField),
		dstEvent.GetAnyField(dstEventField),
		op,
	)
}

func NewJoinFilter(cfg JoinConfig) (*JoinFilter, error) {
	return &JoinFilter{
		name:     cfg.Name,
		not:      cfg.Not,
		sections: cfg.Sections,
	}, nil
}
