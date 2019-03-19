package correlation

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestCommonRule(t *testing.T) {
	correlationChannel := make(chan *normalization.Event, 64)
	outChannel := make(chan *normalization.Event)

	dosRule, err := NewRule(buildTestCorrelationConfigDOS(), outChannel, correlationChannel)
	assert.Assert(t, err == nil)

	ddosRule, err := NewRule(buildTestCorrelationConfigDDOS(), outChannel, correlationChannel)
	assert.Assert(t, err == nil)

	rules := []IRule{dosRule, ddosRule}
	var correlatedEvents []*normalization.Event

	go func() {
		for event := range correlationChannel {
			dosRule.Feed(event)
			ddosRule.Feed(event)
		}
	}()

	// Output routine
	go func() {
		for event := range outChannel {
			fmt.Println(event)
			correlatedEvents = append(correlatedEvents, event)
		}
	}()

	dosRule.Start()
	ddosRule.Start()

	for _, event := range generateTestEvents() {
		for _, r := range rules {
			r.Feed(event)
		}
	}

	<-time.After(5 * time.Second)
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
	assert.Equal(t, correlatedEvents[2].BaseEventCount, 1)
	assert.Equal(t, correlatedEvents[2].AggregatedEventCount, 0)
}

//func BenchmarkRule(b *testing.B) {
//	channel := make(chan normalization.Event)
//	rule, _ := NewRule(Config{
//		Filter:          getTestFilterConfig(),
//		Threshold:       100,
//		Window:          time.Second,
//		IdenticalFields: getTestIdenticalFields(),
//	}, channel, nil)
//
//	go func() {
//		for {
//			<-channel
//		}
//	}()
//
//	events := generateTestEvents()
//
//	b.ReportAllocs()
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		rule.Feed(events[0])
//	}
//}

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
		Triggers: map[string]TriggerConfig{
			TriggerFirstThreshold: {
				Actions: []ActionConfig{{
					Kind: actions.KindRelease,
					Release: actions.ReleaseConfig{
						MutateConfigs: []actions.MutateConfig{
							{Field: "Message", Constant: "DoS attack detected"},
							{Field: "SourceIPAddress", ValueSourceKind: actions.ValueSourceKindEvent, ValueSourceName: "DoS"},
							{Field: "DestinationIPAddress", ValueSourceKind: actions.ValueSourceKindEvent, ValueSourceName: "DoS"},
						},
					},
				}},
			},
		},
		AggregationRules: []aggregation.Config{{
			Name:      "DoS",
			Threshold: 1,
			IdenticalFields: []string{
				"SourceIPAddress",
				"DestinationIPAddress",
			},
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
		Name:   "DDoS",
		Window: time.Second,
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
		AggregationRules: []aggregation.Config{{
			Name:            "DDoS",
			Threshold:       2,
			IdenticalFields: []string{"DestinationIPAddress"},
			UniqueFields:    []string{"SourceIPAddress"},
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
