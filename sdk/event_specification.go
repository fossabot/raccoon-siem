package sdk

type eventSpecification struct {
	id        string
	threshold int
	filter    IFilter
}
