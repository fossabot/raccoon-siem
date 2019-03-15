package collector

import (
	"github.com/tephrocactus/raccoon-siem/collector"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
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
	channel := startCollector()
	for _, input := range inputs {
		channel <- connectors.Output{
			Connector: "test",
			Data:      input,
		}
	}
	time.Sleep(time.Second)
}

func BenchmarkCollector(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	channel := startCollector()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			channel <- connectors.Output{
				Connector: "test",
				Data:      input,
			}
		}
	}
}

func startCollector() connectors.OutputChannel {
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
			{SourceField: "msg", Extra: &normalizers.ExtraConfig{
				Normalizer:   extraNormalizer,
				TriggerField: "app",
				TriggerValue: []byte("rest"),
			}},
		},
	})

	dropFilter, _ := filters.New(filters.Config{
		Name: "drop cron",
		Sections: []filters.SectionConfig{{
			Conditions: []filters.ConditionConfig{{
				Field: "OriginProcessName",
				Op:    filters.OpEQ,
				Rv:    "cron",
			}},
		}},
	})

	channel := make(connectors.OutputChannel)
	processor := collector.InputProcessor{
		InputChannel: channel,
		Normalizer:   mainNormalizer,
		DropFilters:  []filters.IFilter{dropFilter},
	}

	if err := processor.Start(); err != nil {
		log.Fatal(err)
	}

	return channel
}
