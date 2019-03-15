package filters

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type joinFilter struct {
	comparator
	not      bool
	name     string
	sections []JoinSectionConfig
}

func (f *joinFilter) ID() string {
	return f.name
}

func (f *joinFilter) Pass(events ...*normalization.Event) bool {
	eventsBySpecID := make(map[string]*normalization.Event)

	for _, e := range events {
		eventsBySpecID[e.CorrelatorEventSpecID] = e
	}

	for _, section := range f.sections {
		if !f.checkSection(events, eventsBySpecID, section) {
			return f.not
		}
	}

	return !f.not
}

func (f *joinFilter) checkSection(
	correlationEvents []*normalization.Event,
	eventsBySpecID map[string]*normalization.Event,
	section JoinSectionConfig,
) bool {
	for _, cond := range section.Conditions {
		srcEvent := eventsBySpecID[cond.LeftEventID]
		dstEvent := eventsBySpecID[cond.RightEventID]

		if !section.Or {
			if !f.joinConditionMatch(srcEvent, cond.LeftEventField, dstEvent, cond.RightEventField, cond.Operator) {
				return section.Not
			}
		} else {
			if f.joinConditionMatch(srcEvent, cond.LeftEventField, dstEvent, cond.RightEventField, cond.Operator) {
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

func (f *joinFilter) joinConditionMatch(
	srcEvent *normalization.Event,
	srcEventField string,
	dstEvent *normalization.Event,
	dstEventField string,
	op string,
) bool {
	return f.compareValues(
		srcEvent.Get(srcEventField),
		dstEvent.Get(dstEventField),
		op,
	)
}

func newJoinFilter(cfg Config) (*joinFilter, error) {
	return &joinFilter{
		name:     cfg.Name,
		not:      cfg.Not,
		sections: cfg.JoinSections,
	}, nil
}
