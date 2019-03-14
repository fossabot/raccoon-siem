package sdk

import "github.com/tephrocactus/raccoon-siem/sdk/normalization"

type ProcessorTask struct {
	Connector string
	Data      []byte
}

type CorrelationChainTask *normalization.Event
type AggregationChainTask *normalization.Event
