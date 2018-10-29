package sdk

func RegisterAggregationRules(
	settings []AggregationRuleSettings,
	filterSettings []FilterSettings,
	aggregationChainChannel chan AggregationChainTask,
) ([]IAggregationRule, error) {
	filters, err := RegisterFilters(filterSettings)

	if err != nil {
		return nil, err
	}

	rulesToReturn := make([]IAggregationRule, 0)
	uniqueRules := make(map[string]*AggregationRule)

	for _, setting := range settings {
		if _, ok := uniqueRules[setting.Name]; ok {
			continue
		}

		rule, err := setting.Compile(filters)

		if err != nil {
			return nil, err
		}

		rule.aggregationChain = aggregationChainChannel

		uniqueRules[rule.name] = rule
		rulesToReturn = append(rulesToReturn, rule)
	}

	return rulesToReturn, nil
}

func RunAggregationRules(rules []IAggregationRule) {
	for i := range rules {
		rules[i].Run()
	}
}
