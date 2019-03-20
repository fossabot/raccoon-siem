package correlation

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"testing"
	"time"
)

var correlatedEvents []*normalization.Event
var rules []IRule

func TestCommonRule(t *testing.T) {
	dosRule, err := NewRule(buildTestCorrelationConfigDOS(), outputFnCommonTest)
	assert.Assert(t, err == nil)

	ddosRule, err := NewRule(buildTestCorrelationConfigDDOS(), outputFnCommonTest)
	assert.Assert(t, err == nil)

	rules = append(rules, dosRule, ddosRule)

	for _, event := range generateTestEvents() {
		inputFnCommonTest(event)
	}

	assert.Equal(t, len(correlatedEvents), 3)

	assert.Equal(t, correlatedEvents[0].CorrelationRuleName, "DoS")
	assert.Equal(t, correlatedEvents[0].Correlated, true)
	assert.Equal(t, correlatedEvents[0].Message, "DoS attack detected")
	assert.Equal(t, correlatedEvents[0].SourceIPAddress, "192.168.2.2")
	assert.Equal(t, correlatedEvents[0].DestinationIPAddress, "192.168.1.254")
	assert.Equal(t, correlatedEvents[0].BaseEventCount, 1)
	assert.Equal(t, correlatedEvents[0].AggregatedEventCount, 0)

	assert.Equal(t, correlatedEvents[1].CorrelationRuleName, "DoS")
	assert.Equal(t, correlatedEvents[1].Correlated, true)
	assert.Equal(t, correlatedEvents[1].Message, "DoS attack detected")
	assert.Equal(t, correlatedEvents[1].SourceIPAddress, "192.168.2.3")
	assert.Equal(t, correlatedEvents[1].DestinationIPAddress, "192.168.1.254")
	assert.Equal(t, correlatedEvents[1].BaseEventCount, 1)
	assert.Equal(t, correlatedEvents[1].AggregatedEventCount, 0)

	assert.Equal(t, correlatedEvents[2].CorrelationRuleName, "DDoS")
	assert.Equal(t, correlatedEvents[2].Correlated, true)
	assert.Equal(t, correlatedEvents[2].Message, "DDoS attack detected")
	assert.Equal(t, correlatedEvents[2].SourceIPAddress, "")
	assert.Equal(t, correlatedEvents[2].DestinationIPAddress, "192.168.1.254")
	assert.Equal(t, correlatedEvents[2].BaseEventCount, 2)
	assert.Equal(t, correlatedEvents[2].AggregatedEventCount, 0)

	for _, event := range correlatedEvents {
		fmt.Println(event)
	}
}

func inputFnCommonTest(event *normalization.Event) {
	for _, rule := range rules {
		rule.Feed(event)
	}
}

func outputFnCommonTest(event *normalization.Event) {
	correlatedEvents = append(correlatedEvents, event)
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
		Window: time.Second,
		IdenticalFields: []string{
			"SourceIPAddress",
			"DestinationIPAddress",
		},
		Triggers: map[string]TriggerConfig{
			TriggerFirstThreshold: {
				Actions: []ActionConfig{{
					Kind: actions.KindRelease,
					Release: actions.ReleaseConfig{
						MutateConfigs: []actions.MutateConfig{
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
						{Field: "Correlated", Op: filters.OpEQ, Value: false},
						{Field: "OriginServiceName", Op: filters.OpEQ, Value: "PeakFlow"},
						{Field: "Message", Op: filters.OpEQ, Value: "DoS"},
					},
				}},
			},
		}},
	}
}

func buildTestCorrelationConfigDDOS() Config {
	return Config{
		Name:            "DDoS",
		Window:          time.Second,
		IdenticalFields: []string{"DestinationIPAddress"},
		UniqueFields:    []string{"SourceIPAddress"},
		Triggers: map[string]TriggerConfig{
			TriggerFirstThreshold: {
				Actions: []ActionConfig{{
					Kind: actions.KindRelease,
					Release: actions.ReleaseConfig{
						MutateConfigs: []actions.MutateConfig{
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
						{Field: "CorrelationRuleName", Op: filters.OpEQ, Value: "DoS"},
					},
				}},
			},
		}},
	}
}
