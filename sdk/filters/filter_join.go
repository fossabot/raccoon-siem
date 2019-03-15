package filters

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type filterJoin struct {
	comparator
	not      bool
	name     string
	sections []*filterJoinSection
}

func (f *filterJoin) ID() string {
	return f.name
}

func (f *filterJoin) Pass(events ...*normalization.Event) bool {
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

func (f *filterJoin) compile(settings *FilterSettings) (*filterJoin, error) {
	// Name

	if settings.Name == "" {
		return nil, fmt.Errorf("filter must have a name")
	}

	f.name = settings.Name

	// Sections

	for _, sect := range settings.Sections {
		compiledJoinSect, err := sect.compileJoin()

		if err != nil {
			return nil, err
		}

		f.sections = append(f.sections, compiledJoinSect)
	}

	return f, nil
}

func (f *filterJoin) checkSection(
	correlationEvents []*normalization.Event,
	eventsBySpecID map[string]*normalization.Event,
	section *filterJoinSection,
) bool {
	for _, cond := range section.conditions {
		srcEvent := eventsBySpecID[cond.LeftEventID]
		dstEvent := eventsBySpecID[cond.RightEventID]

		if !section.or {
			if !f.joinConditionMatch(srcEvent, cond.LeftEventField, dstEvent, cond.RightEventField, cond.Operator) {
				return section.not
			}
		} else {
			if f.joinConditionMatch(srcEvent, cond.LeftEventField, dstEvent, cond.RightEventField, cond.Operator) {
				return !section.not
			}
		}
	}

	if !section.or {
		return !section.not
	} else {
		return section.not
	}
}

func (f *filterJoin) joinConditionMatch(
	srcEvent *normalization.Event,
	srcEventField string,
	dstEvent *normalization.Event,
	dstEventField string,
	op byte,
) bool {
	srcEventValue, srcEventType := srcEvent.GetField(srcEventField)
	dstEventValue := dstEvent.GetFieldNoType(dstEventField)
	return f.compareValues(srcEventValue, srcEventType, dstEventValue, op)
}
