package collector

import (
	"github.com/tephrocactus/raccoon-siem/collector"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers"
	"log"
	"runtime"
	"testing"
)

var input = []byte(`<34>Mar 14 22:14:15 mymachine rest: { "service": { "name": "rest", "version": "v1.0" }, "request": { "url": "/session", "status": 200 }, "message": "OK" }`)

func BenchmarkParsingNormalization(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

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
			{SourceField: "msg", Extra: &normalizers.ExtraConfig{Normalizer: extraNormalizer}},
		},
	})

	in := make(connectors.OutputChannel)
	processor := collector.InputProcessor{
		InputChannel: in,
		Normalizer:   mainNormalizer,
		Workers:      runtime.NumCPU(),
	}

	if err := processor.Start(); err != nil {
		log.Fatal(err)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		in <- connectors.Output{
			Connector: "test",
			Data:      helpers.CopyBytes(input),
		}
	}
}
