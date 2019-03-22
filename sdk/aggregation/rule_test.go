package aggregation

import (
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestRule(t *testing.T) {
	var aggregatedEvents []*normalization.Event
	outputFn := func(event *normalization.Event) {
		aggregatedEvents = append(aggregatedEvents, event)
	}

	threshold := 10
	rule, err := NewRule(Config{
		Name:            "netflow",
		Filter:          getTestFilterConfig(),
		Threshold:       threshold,
		IdenticalFields: getTestIdenticalFields(),
		SumFields:       getTestSumFields(),
	}, outputFn)
	assert.Assert(t, err == nil)

	for i := 0; i < threshold; i++ {
		for _, e := range getTestEvents() {
			rule.Feed(e)
		}
	}

	assert.Equal(t, len(aggregatedEvents), 2)

	uPorts := make(map[string]bool)
	for _, event := range aggregatedEvents {
		switch event.DestinationPort {
		case "80", "443":
			uPorts[event.DestinationPort] = true
			assert.Equal(t, event.RequestBytesIn, int64(10))
		default:
			t.Fatalf("enexpected event found: %v", event)
		}
	}

	assert.Equal(t, len(uPorts), 2)
}

func TestRuleWindow(t *testing.T) {
	var aggregatedEvents []*normalization.Event
	outputFn := func(event *normalization.Event) {
		aggregatedEvents = append(aggregatedEvents, event)
	}

	rule, err := NewRule(Config{
		Name:            "netflow",
		Filter:          getTestFilterConfig(),
		Window:          time.Second,
		IdenticalFields: getTestIdenticalFields(),
		SumFields:       getTestSumFields(),
	}, outputFn)
	assert.Assert(t, err == nil)

	rule.Start()

	for i := 0; i < 10; i++ {
		for _, e := range getTestEvents() {
			rule.Feed(e)
		}
	}

	time.Sleep(2 * time.Second)
	assert.Equal(t, len(aggregatedEvents), 2)
}

func BenchmarkRule(b *testing.B) {
	outputFn := func(event *normalization.Event) {}

	rule, _ := NewRule(Config{
		Name:            "netflow",
		Filter:          getTestFilterConfig(),
		Threshold:       100,
		Window:          time.Second,
		IdenticalFields: getTestIdenticalFields(),
	}, outputFn)

	events := getTestEvents()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rule.Feed(events[0])
	}
}

func getTestFilterConfig() filters.Config {
	return filters.Config{
		Sections: []filters.SectionConfig{{
			Conditions: []filters.ConditionConfig{{
				Field:    "OriginServiceName",
				Op:       filters.OpEQ,
				Constant: "netflow",
			}},
		}},
	}
}

func getTestEvents() (events []*normalization.Event) {
	validEventOne := &normalization.Event{
		OriginServiceName:        "netflow",
		SourceIPAddress:          "192.168.1.1",
		DestinationIPAddress:     "192.168.1.254",
		DestinationPort:          "443",
		RequestTransportProtocol: "tcp",
		RequestStatus:            "200",
		RequestBytesIn:           1,
	}

	validEventTwo := &normalization.Event{
		OriginServiceName:        "netflow",
		SourceIPAddress:          "192.168.2.1",
		DestinationIPAddress:     "192.168.2.254",
		DestinationPort:          "80",
		RequestTransportProtocol: "tcp",
		RequestStatus:            "201",
		RequestBytesIn:           1,
	}

	invalidEvent := &normalization.Event{
		Correlated: true,
	}

	events = append(events, validEventOne, validEventTwo, invalidEvent)
	return
}

func getTestIdenticalFields() []string {
	return []string{
		"SourceIPAddress",
		"DestinationIPAddress",
		"DestinationPort",
		"RequestTransportProtocol",
	}
}

func getTestSumFields() []string {
	return []string{"RequestBytesIn"}
}
