package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/mutation"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"log"
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
		ValueSourceKind: ValueSourceKindDict,
		ValueSourceName: raccoon,
		KeyFields:       []string{"Trace"},
		Field:           "Message",
	}
	Enrich(cfg, &event)
	assert.Equal(t, event.Message, "error")

	cfg = Config{
		ValueSourceKind: ValueSourceKindConst,
		Constant:        "1080",
		Field:           "RequestResults",
	}
	Enrich(cfg, &event)
	assert.Equal(t, event.RequestResults, int64(1080))

	cfg = Config{
		ValueSourceKind: ValueSourceKindConst,
		Constant:        "1081",
		Field:           "RequestResults",
		TriggerField:    "Message",
		TriggerValue:    "error",
	}
	Enrich(cfg, &event)
	assert.Equal(t, event.RequestResults, int64(1081))

	cfg = Config{
		ValueSourceKind: ValueSourceKindConst,
		Constant:        "1082",
		Field:           "RequestResults",
		TriggerField:    "Severity",
		TriggerValue:    "error",
	}

	Enrich(cfg, &event)
	assert.Equal(t, event.RequestResults, int64(1081))
}

func TestMutation(t *testing.T) {
	event := &normalization.Event{
		RequestUser: "Tephro@gmail.com",
	}

	configs := []Config{
		{
			Field:            "RequestReferrer",
			ValueSourceKind:  ValueSourceKindEvent,
			ValueSourceField: "RequestUser",
			Mutation:         []mutation.Config{{Kind: mutation.KindRegexp, Expression: "([^@]+)@.+"}},
		},
		{
			Field:           "RequestReferrer",
			ValueSourceKind: ValueSourceKindEvent,
			Mutation:        []mutation.Config{{Kind: mutation.KindLower}},
		},
		{
			Field:           "RequestReferrer",
			ValueSourceKind: ValueSourceKindEvent,
			Mutation:        []mutation.Config{{Kind: mutation.KindSubstring, Start: 0, End: 3}},
		},
	}

	for i := range configs {
		if err := configs[i].Validate(); err != nil {
			log.Fatal(err)
		}
		Enrich(configs[i], event)
	}

	assert.Equal(t, event.RequestReferrer, "tep")
}

func BenchmarkEnrich(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	fillDictionaryStorage()
	cfg := Config{
		ValueSourceKind: ValueSourceKindDict,
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
	raccoonDict := dictionaries.Config{
		Name: raccoon,
		Data: map[string]string{
			"ERROR": "error",
			"DEBUG": "debug",
			"INFO":  "info",
		},
	}

	weirdDict := dictionaries.Config{
		Name: weird,
		Data: map[string]string{
			"1": "error",
			"2": "debug",
			"3": "info",
		},
	}

	globals.Dictionaries = dictionaries.NewStorage([]dictionaries.Config{raccoonDict, weirdDict})
}
