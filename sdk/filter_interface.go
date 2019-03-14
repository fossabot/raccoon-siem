package sdk

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

type IFilter interface {
	ID() string
	Pass(events []*normalization.Event) bool
}
