package sdk

func newEventParser(spec *parserSpecification) IParser {
	return &eventParser{spec: spec}
}

type eventParser struct {
	spec *parserSpecification
}

func (p *eventParser) AddSub(sub IParser) {}

func (p *eventParser) SubNames() []string {
	return p.spec.subs
}

func (p *eventParser) ID() string {
	return p.spec.name
}

func (p *eventParser) Parse(input []byte, target *Event) (*Event, error) {
	event := new(Event)

	if err := event.FromMsgPack(input); err != nil {
		return nil, err
	}

	return event, nil
}
