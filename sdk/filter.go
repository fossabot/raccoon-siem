package sdk

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type filter struct {
	comparator
	not       bool
	name      string
	sections  []*filterSection
	variables map[string]*variable
}

func (f *filter) ID() string {
	return f.name
}

func (f *filter) Pass(events []*normalization.Event) bool {
	for _, section := range f.sections {
		if !f.checkSection(events[0], section) {
			return f.not
		}
	}
	return !f.not
}

func (f *filter) compile(settings *FilterSettings) (*filter, error) {
	f.not = settings.Not
	f.variables = make(map[string]*variable)

	// Name

	if settings.Name == "" {
		return nil, fmt.Errorf("filter must have a name")
	}

	f.name = settings.Name

	// Sections

	for _, sect := range settings.Sections {
		compiledSect, err := sect.compile()

		if err != nil {
			return nil, err
		}

		f.sections = append(f.sections, compiledSect)
	}

	// Variables

	for _, expr := range settings.Variables {
		vb, err := new(variable).compile(expr)

		if err != nil {
			return nil, err
		}

		if _, ok := f.variables[vb.name]; ok {
			continue
		}

		f.variables[vb.name] = vb
	}

	return f, nil
}

func (f *filter) checkSection(event *normalization.Event, section *filterSection) bool {
	for _, cond := range section.conditions {
		if !section.or {
			if !f.conditionMatch(event, cond) {
				return section.not
			}
		} else {
			if f.conditionMatch(event, cond) {
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

func (f *filter) conditionMatch(event *normalization.Event, cond *filterCondition) bool {
	if cond.incFilter != nil {
		return cond.incFilter.Pass([]*normalization.Event{event})
	}

	lv, err := cond.leftValue.resolve(f.variables, event)

	if err != nil {
		DebugError(err)
		return false
	}

	rv, err := cond.rightValue.resolve(f.variables, event)

	if err != nil {
		DebugError(err)
		return false
	}

	return f.compareValues(lv, cond.leftValue.kind, rv, cond.operator)
}
