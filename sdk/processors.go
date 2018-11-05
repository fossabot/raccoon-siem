package sdk

type ProcessorTask struct {
	Source string
	Data   []byte
}

type CorrelationChainTask *Event
type AggregationChainTask *Event
