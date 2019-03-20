package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/dictionary"
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"testing"
)

const (
	raccoon = "raccoon"
	weird   = "weird"
)

func TestEnrichment(t *testing.T) {
	fillDictionaryStorage()

	event := normalization.Event{
		Trace: "ERROR",
	}

	cfg := Config{
		ValueSourceKind: FromDict,
		ValueSourceName: raccoon,
		KeyFields:       []string{"Trace"},
		Field:           "Message",
	}
	Enrich(cfg, &event)
	assert.Equal(t, event.Message, "error")

	cfg = Config{
		ValueSourceKind: FromConst,
		Constant:        "1080",
		Field:           "RequestResults",
	}
	Enrich(cfg, &event)
	assert.Equal(t, event.RequestResults, int64(1080))

	cfg = Config{
		ValueSourceKind: FromConst,
		Constant:        "1081",
		Field:           "RequestResults",
		TriggerField:    "Message",
		TriggerValue:    "error",
	}
	Enrich(cfg, &event)
	assert.Equal(t, event.RequestResults, int64(1081))

	cfg = Config{
		ValueSourceKind: FromConst,
		Constant:        "1082",
		Field:           "RequestResults",
		TriggerField:    "Severity",
		TriggerValue:    "error",
	}
	Enrich(cfg, &event)
	assert.Equal(t, event.RequestResults, int64(1081))

}

func BenchmarkEnrich(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	fillDictionaryStorage()
	cfg := Config{
		ValueSourceKind: FromDict,
		ValueSourceName: raccoon,
		KeyFields:       []string{"Trace"},
		Field:           "Message",
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		event := normalization.Event{
			Trace: "ERROR",
		}

		Enrich(cfg, &event)
	}
}

func fillDictionaryStorage() {
	dictionariesData := make(map[string]map[interface{}]interface{})

	raccoonDict := make(map[interface{}]interface{})
	raccoonDict["ERROR"] = "error"
	raccoonDict["DEBUG"] = "debug"
	raccoonDict["INFO"] = "info"

	weirdDict := make(map[interface{}]interface{})
	weirdDict["1"] = "error"
	weirdDict["2"] = "debug"
	weirdDict["3"] = "info"

	dictionariesData[raccoon] = raccoonDict
	dictionariesData[weird] = weirdDict

	globals.DictionaryStorage = dictionary.NewDictionaryStorage(dictionary.Config{Data: dictionariesData})
}
