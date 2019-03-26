package correlation

import (
	"github.com/satori/go.uuid"
	"github.com/tephrocactus/raccoon-siem/sdk/actions"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/filters"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestRecoveryRule(t *testing.T) {
	cfg := buildTestCorrelationConfigRecovery()
	rule, err := NewRule(cfg, outputFnRecoveryTest)
	assert.Equal(t, err, nil)
	rule.Start()

	testEvents := []*normalization.Event{
		{
			ID:                uuid.NewV4().String(),
			OriginServiceName: "Nagios",
			Message:           "server rebooted",
			SourceDNSName:     "srv-example-01",
		},
		{
			ID:                uuid.NewV4().String(),
			OriginServiceName: "Nagios",
			Message:           "server rebooted",
			SourceDNSName:     "srv-example-02",
		},
		{
			ID:                uuid.NewV4().String(),
			OriginServiceName: "Nagios",
			Message:           "server booted",
			SourceDNSName:     "srv-example-01",
		},
	}

	testRules = append(testRules, rule)
	for _, event := range testEvents {
		inputFnRecoveryTest(event)
	}

	time.Sleep(2 * time.Second)
	assert.Equal(t, len(testCorrelatedEvents), 1)
	assert.Equal(t, testCorrelatedEvents[0].SourceDNSName, "srv-example-02")
}

func inputFnRecoveryTest(event *normalization.Event) {
	for _, rule := range testRules {
		rule.Feed(event)
	}
}

func outputFnRecoveryTest(event *normalization.Event) {
	testCorrelatedEvents = append(testCorrelatedEvents, event)
	inputFnCommonTest(event)
}

func buildTestCorrelationConfigRecovery() Config {
	return Config{
		Name:            "Test recovery rule",
		Kind:            RuleKindRecovery,
		Window:          1,
		IdenticalFields: []string{"SourceDNSName"},
		Triggers: map[string]TriggerConfig{
			TriggerTimeout: {
				Actions: []ActionConfig{{
					Kind: actions.KindRelease,
					Release: actions.ReleaseConfig{
						EnrichmentConfigs: []enrichment.Config{
							{Field: "Message", Constant: "Server rebooted and remains in down state"},
						},
					},
				}},
			},
		},
		Selectors: []EventSelector{
			{
				Tag:       "Reboot",
				Threshold: 1,
				Filter: filters.Config{
					Sections: []filters.SectionConfig{{
						Conditions: []filters.ConditionConfig{
							{Field: "OriginServiceName", Op: filters.OpEQ, Constant: "Nagios"},
							{Field: "Message", Op: filters.OpEQ, Constant: "server rebooted"},
						},
					}},
				},
			},
			{
				Tag:       "Boot",
				Threshold: 1,
				Recovery:  true,
				Filter: filters.Config{
					Sections: []filters.SectionConfig{{
						Conditions: []filters.ConditionConfig{
							{Field: "OriginServiceName", Op: filters.OpEQ, Constant: "Nagios"},
							{Field: "Message", Op: filters.OpEQ, Constant: "server booted"},
						},
					}},
				},
			},
		},
	}
}
