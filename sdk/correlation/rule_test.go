package correlation

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestRulesDDOS(t *testing.T) {
	channel := make(chan normalization.Event)
	//chainChannel := make(chan normalization.Event)

	dosRule, err := NewRule(buildTestCorrelationConfigDOS(), channel)
	assert.Assert(t, err == nil)

	ddosRule, err := NewRule(buildTestCorrelationConfigDDOS(), channel)
	assert.Assert(t, err == nil)

	dosRule.Start()
	ddosRule.Start()

	//var aggregatedEvents []*normalization.Event
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//
	//go func() {
	//	timeout := time.After(time.Second)
	//	for {
	//		select {
	//		case <-timeout:
	//			wg.Done()
	//		case event := <-channel:
	//			aggregatedEvents = append(aggregatedEvents, &event)
	//		}
	//	}
	//}()

	for _, e := range getTestEvents() {
		dosRule.Feed(e)
	}

	//wg.Wait()
	//assert.Equal(t, len(aggregatedEvents), 2)
	//
	//uPorts := make(map[string]bool)
	//for _, event := range aggregatedEvents {
	//	switch event.DestinationPort {
	//	case "80", "443":
	//		uPorts[event.DestinationPort] = true
	//		assert.Equal(t, event.RequestBytesIn, int64(10))
	//	default:
	//		t.Fatalf("enexpected event found: %v", event)
	//	}
	//}
	//
	//assert.Equal(t, len(uPorts), 2)
	time.Sleep(5 * time.Second)
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
//	events := getTestEvents()
//
//	b.ReportAllocs()
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		rule.Feed(events[0])
//	}
//}

func getTestEvents() (events []*normalization.Event) {
	for i := 2; i < 4; i++ {
		for j := 1; j < 3; j++ {
			for c := 0; c < 5; c++ {
				events = append(events, &normalization.Event{
					OriginServiceName:    "PeakFlow",
					SourceIPAddress:      fmt.Sprintf("192.168.%d.%d", i, j),
					DestinationIPAddress: "192.168.1.254",
					Message:              "DoS",
				})
			}
		}
	}
	return
}

func buildTestCorrelationConfigDOS() Config {
	return Config{
		Name:   "DoS",
		Window: time.Second,
		AggregationRules: []aggregation.Config{{
			Name:      "DoS",
			Threshold: 2,
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
		AggregationRules: []aggregation.Config{{
			Name:            "DDoS",
			Threshold:       2,
			IdenticalFields: []string{"DestinationIPAddress"},
			UniqueFields:    []string{"SourceIPAddress"},
			SumFields:       []string{"SourceIPAddress"},
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
