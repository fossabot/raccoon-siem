package sdk

import "github.com/tephrocactus/raccoon-siem/sdk/filters"

// Loads, compiles and runs correlation rules
func RegisterCorrelationRules(
	settings []CorrelationRuleSettings,
	filters []filters.IFilter,
	correlationChainChannel chan CorrelationChainTask,
) ([]ICorrelationRule, error) {
	rulesToReturn := make([]ICorrelationRule, 0)
	uniqueCorrelationRules := make(map[string]*CorrelationRule)

	for _, setting := range settings {
		if _, ok := uniqueCorrelationRules[setting.Name]; ok {
			continue
		}

		rule, err := setting.Compile(filters)

		if err != nil {
			return nil, err
		}

		rule.correlationChain = correlationChainChannel

		uniqueCorrelationRules[rule.name] = rule
		rulesToReturn = append(rulesToReturn, rule)
	}

	return rulesToReturn, nil
}

func RunCorrelationRules(rules []ICorrelationRule) {
	for _, rule := range rules {
		rule.Run()
	}
}
