package collector

import (
	"github.com/tephrocactus/raccoon-siem/sdk"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers"
	"runtime"
)

type InputProcessor struct {
	InputChannel       connectors.OutputChannel
	AggregationChannel chan *normalization.Event
	OutputChannel      chan *normalization.Event
	Normalizer         normalizers.INormalizer
	DropFilters        []filters.IFilter
}

func (r *InputProcessor) Start() error {
	for i := 0; i < runtime.NumCPU(); i++ {
		go r.inputRoutine()
	}
	return nil
}

func (r *InputProcessor) inputRoutine() {
inputLoop:
	for input := range r.InputChannel {
		event := r.Normalizer.Normalize(input.Data, nil)
		if event == nil {
			continue
		}

		for _, dropFilter := range r.DropFilters {
			if dropFilter.Pass(event) {
				continue inputLoop
			}
		}

		event.ID = sdk.GetUUID()
		event.SourceID = input.Connector
	}
}
