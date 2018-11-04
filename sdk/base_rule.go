package sdk

type BaseRule struct {
	name        string
	aggregation *aggregation
	eventSpecs  []*eventSpecification
}

func (r *BaseRule) ID() string {
	return r.name
}
