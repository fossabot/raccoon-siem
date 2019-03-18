package enrichment

import (
	"github.com/tephrocactus/raccoon-siem/sdk/dictionary"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"testing"
)

const (
	raccoon = "raccoon"
	weird   = "weird"
)

func TestDictionary(t *testing.T) {
	dictionariesData := make(map[string]map[string]string)

	raccoonDict := make(map[string]string)
	raccoonDict["ERROR"] = "error"
	raccoonDict["DEBUG"] = "debug"
	raccoonDict["INFO"] = "info"

	weirdDict := make(map[string]string)
	weirdDict["1"] = "error"
	weirdDict["2"] = "debug"
	weirdDict["3"] = "info"

	dictionariesData[raccoon] = raccoonDict
	dictionariesData[weird] = weirdDict

	dictionary.MockDictionary = dictionary.NewDictionary(dictionary.Config{Data: dictionariesData})

	event := normalization.Event{}

	cfg := enrichment.EnrichConfig{
		From:        enrichment.FromDict,
		FromKey:     raccoon,
		KeyField:    "ERROR",
		TargetField: "Trace",
	}
	enrichment.Enrich(cfg, &event)
	assert.Equal(t, event.Trace, "error")

	cfg = enrichment.EnrichConfig{
		From:        enrichment.FromConst,
		Const:       "1080",
		TargetField: "RequestResults",
	}
	enrichment.Enrich(cfg, &event)
	assert.Equal(t, event.RequestResults, int64(1080))

	cfg = enrichment.EnrichConfig{
		From:        enrichment.FromField,
		KeyField:    "OriginIPAddress",
		TargetField: "DestinationIPAddress",
	}
	event.OriginIPAddress = "127.0.0.1"
	enrichment.Enrich(cfg, &event)
	assert.Equal(t, event.DestinationIPAddress, "127.0.0.1")
}
