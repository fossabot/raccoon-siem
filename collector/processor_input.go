package collector

import (
	"github.com/tephrocactus/raccoon-siem/sdk"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers"
)

type InputProcessor struct {
	InputChannel       connectors.OutputChannel
	AggregationChannel chan *normalization.Event
	OutputChannel      chan *normalization.Event
	Normalizer         normalizers.INormalizer
	Filters            []sdk.IFilter
	Workers            int
}

func (r *InputProcessor) Start() error {
	for i := 0; i < r.Workers; i++ {
		go r.inputRoutine()
	}
	return nil
}

func (r *InputProcessor) inputRoutine() {
	for input := range r.InputChannel {
		event := r.Normalizer.Normalize(input.Data, nil)
		if event == nil {
			continue
		}

		for _, filter := range r.Filters {
			filter.Pass()
		}

		event.ID = sdk.GetUUID()
		event.SourceID = input.Connector
	}
}
