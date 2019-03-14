package sdk

func RegisterConnectors(configs []UniversalConnectorConfig, processorChannel chan *ProcessorTask) ([]IConnector, error) {
	uniqueConnectors := make(map[string]IConnector)
	result := make([]IConnector, 0)

	for _, config := range configs {
		if _, ok := uniqueConnectors[config.Name]; ok {
			continue
		}

		s, err := NewConnector(config, processorChannel)
		if err != nil {
			return nil, err
		}

		uniqueConnectors[config.Name] = s
		result = append(result, s)
	}

	return result, nil
}

func RunConnectors(connectors []IConnector) error {
	for _, connector := range connectors {
		if err := connector.Run(); err != nil {
			return err
		}
	}
	return nil
}
