package sdk

type IFilter interface {
	ID() string
	Pass(events []*Event) bool
}
