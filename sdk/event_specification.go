package sdk

import "github.com/tephrocactus/raccoon-siem/sdk/filters"

type eventSpecification struct {
	id        string
	threshold int
	filter    filters.IFilter
}
