package correlation

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"testing"
)

var testCorrelatedEvents []*normalization.Event
var testRules []IRule

func TestCommonRule(t *testing.T) {
	testRules = nil
	testCorrelatedEvents = nil

	dosRule, err := NewRule(buildTestCorrelationConfigDOS(), outputFnCommonTest)
	assert.Assert(t, err == nil)

	ddosRule, err := NewRule(buildTestCorrelationConfigDDOS(), outputFnCommonTest)
	assert.Assert(t, err == nil)

	testRules = append(testRules, dosRule, ddosRule)

	for _, event := range generateTestEvents() {
		inputFnCommonTest(event)
	}

	assert.Equal(t, len(testCorrelatedEvents), 3)

	assert.Equal(t, testCorrelatedEvents[0].CorrelationRuleName, "DoS")
	assert.Equal(t, testCorrelatedEvents[0].Correlated, true)
	assert.Equal(t, testCorrelatedEvents[0].Message, "DoS attack detected")
	assert.Equal(t, testCorrelatedEvents[0].SourceIPAddress, "192.168.2.2")
	assert.Equal(t, testCorrelatedEvents[0].DestinationIPAddress, "192.168.1.254")
	assert.Equal(t, testCorrelatedEvents[0].BaseEventCount, 1)
	assert.Equal(t, testCorrelatedEvents[0].AggregatedEventCount, 0)

	assert.Equal(t, testCorrelatedEvents[1].CorrelationRuleName, "DoS")
	assert.Equal(t, testCorrelatedEvents[1].Correlated, true)
	assert.Equal(t, testCorrelatedEvents[1].Message, "DoS attack detected")
	assert.Equal(t, testCorrelatedEvents[1].SourceIPAddress, "192.168.2.3")
	assert.Equal(t, testCorrelatedEvents[1].DestinationIPAddress, "192.168.1.254")
	assert.Equal(t, testCorrelatedEvents[1].BaseEventCount, 1)
	assert.Equal(t, testCorrelatedEvents[1].AggregatedEventCount, 0)

	assert.Equal(t, testCorrelatedEvents[2].CorrelationRuleName, "DDoS")
	assert.Equal(t, testCorrelatedEvents[2].Correlated, true)
	assert.Equal(t, testCorrelatedEvents[2].Message, "DDoS attack detected")
	assert.Equal(t, testCorrelatedEvents[2].SourceIPAddress, "")
	assert.Equal(t, testCorrelatedEvents[2].DestinationIPAddress, "192.168.1.254")
	assert.Equal(t, testCorrelatedEvents[2].BaseEventCount, 2)
	assert.Equal(t, testCorrelatedEvents[2].AggregatedEventCount, 0)

	for _, event := range testCorrelatedEvents {
		fmt.Println(event)
	}
}

func inputFnCommonTest(event *normalization.Event) {
	for _, rule := range testRules {
		rule.Feed(event)
	}
}

func outputFnCommonTest(event *normalization.Event) {
	testCorrelatedEvents = append(testCorrelatedEvents, event)
	inputFnCommonTest(event)
}

func BenchmarkCommonRule(b *testing.B) {
	dosRule, _ := NewRule(buildTestCorrelationConfigDOS(), nil)
	events := generateTestEvents()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, event := range events {
			dosRule.Feed(event)
		}
	}
}

func generateTestEvents() (events []*normalization.Event) {
	for i := 2; i < 4; i++ {
		events = append(events, &normalization.Event{
			ID:                   uuid.NewV4().String(),
			OriginServiceName:    "PeakFlow",
			SourceIPAddress:      fmt.Sprintf("192.168.2.%d", i),
			DestinationIPAddress: "192.168.1.254",
			Message:              "DoS",
		})
	}
	return
}

func buildTestCorrelationConfigDOS() Config {
	return Config{
		Name:   "DoS",
		Window: 1,
		IdenticalFields: []string{
			"SourceIPAddress",
			"DestinationIPAddress",
		},
		Triggers: map[string]TriggerConfig{
			TriggerFirstThreshold: {
				Actions: []ActionConfig{{
					Kind: actions.KindRelease,
					Release: actions.ReleaseConfig{
						EnrichmentConfigs: []enrichment.Config{
							{Field: "Message", Constant: "DoS attack detected"},
						},
					},
				}},
			},
		},
		Selectors: []EventSelector{{
			Tag:       "DoS",
			Threshold: 1,
			Filter: filters.Config{
				Sections: []filters.SectionConfig{{
					Conditions: []filters.ConditionConfig{
						{Field: "Correlated", Op: filters.OpEQ, Constant: false},
						{Field: "OriginServiceName", Op: filters.OpEQ, Constant: "PeakFlow"},
						{Field: "Message", Op: filters.OpEQ, Constant: "DoS"},
					},
				}},
			},
		}},
	}
}

func buildTestCorrelationConfigDDOS() Config {
	return Config{
		Name:            "DDoS",
		Window:          1,
		IdenticalFields: []string{"DestinationIPAddress"},
		UniqueFields:    []string{"SourceIPAddress"},
		Triggers: map[string]TriggerConfig{
			TriggerFirstThreshold: {
				Actions: []ActionConfig{{
					Kind: actions.KindRelease,
					Release: actions.ReleaseConfig{
						EnrichmentConfigs: []enrichment.Config{
							{Field: "Incident", Constant: "true"},
							{Field: "Message", Constant: "DDoS attack detected"},
						},
					},
				}},
			},
		},
		Selectors: []EventSelector{{
			Tag:       "DDoS",
			Threshold: 2,
			Filter: filters.Config{
				Sections: []filters.SectionConfig{{
					Conditions: []filters.ConditionConfig{
						{Field: "CorrelationRuleName", Op: filters.OpEQ, Constant: "DoS"},
					},
				}},
			},
		}},
	}
}
