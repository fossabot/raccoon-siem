package sdk

type BaseRule struct {
	name        string
	aggregation *aggregation
	eventSpecs  []*eventSpecification
}
