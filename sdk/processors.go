package sdk

type ProcessorTask struct {
	Connector string
	Data      []byte
}

type CorrelationChainTask *Event
type AggregationChainTask *Event
