package collector

import (
	"github.com/tephrocactus/raccoon-siem/collector"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers"
	"log"
	"testing"
	"time"
)

var inputs = [][]byte{
	[]byte(`<34>Mar 14 22:14:15 mymachine rest: { "service": { "name": "rest", "version": "v1.0" }, "request": { "url": "/session", "status": 200 }, "message": "OK" }`),
	[]byte(`<34>Mar 14 22:15:15 myserver cron: job success`),
}

func TestCollector(t *testing.T) {
	in, out := startCollector()

	go func() {
		for range out {

		}
	}()

	for _, input := range inputs {
		in <- connectors.Output{
			Connector: "test",
			Data:      input,
		}
	}

	time.Sleep(time.Second)
}

func BenchmarkCollector(b *testing.B) {
	in, out := startCollector()

	go func() {
		for range out {

		}
	}()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			in <- connectors.Output{
				Connector: "test",
				Data:      input,
			}
		}
	}
}

func startCollector() (connectors.OutputChannel, chan normalization.Event) {
	extraNormalizer, _ := normalizers.New(normalizers.Config{
		Name: "json",
		Kind: normalizers.KindJSON,
		Mapping: []normalizers.MappingConfig{
			{SourceField: "service.name", EventField: "OriginServiceName"},
			{SourceField: "service.version", EventField: "OriginServiceVersion"},
			{SourceField: "request.url", EventField: "RequestURL"},
			{SourceField: "message", EventField: "Message"},
		},
	})

	mainNormalizer, _ := normalizers.New(normalizers.Config{
		Name: "syslog",
		Kind: normalizers.KindSyslog,
		Mapping: []normalizers.MappingConfig{
			{SourceField: "host", EventField: "OriginDNSName"},
			{SourceField: "timestamp", EventField: "OriginTimestamp"},
			{SourceField: "severity", EventField: "OriginSeverity"},
			{SourceField: "app", EventField: "OriginProcessName"},
			{SourceField: "msg", Extra: []normalizers.ExtraConfig{{
				Normalizer:   extraNormalizer,
				TriggerField: "app",
				TriggerValue: []byte("rest"),
			}}},
		},
	})

	dropFilter, _ := filters.NewFilter(filters.Config{
		Name: "drop cron",
		Sections: []filters.SectionConfig{{
			Conditions: []filters.ConditionConfig{{
				Field: "OriginProcessName",
				Op:    filters.OpEQ,
				Value: "cron",
			}},
		}},
	})

	inChannel := make(connectors.OutputChannel)
	outChannel := make(chan normalization.Event)
	processor := collector.Processor{
		InputChannel:  inChannel,
		OutputChannel: outChannel,
		Normalizer:    mainNormalizer,
		DropFilters:   []*filters.Filter{dropFilter},
	}

	if err := processor.Start(); err != nil {
		log.Fatal(err)
	}

	return inChannel, outChannel
}
