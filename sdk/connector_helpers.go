package sdk

import "github.com/tephrocactus/raccoon-siem/sdk/connectors"

func RegisterConnectors(
	configs []connectors.Config,
	channel connectors.OutputChannel,
) ([]connectors.IConnector, error) {
	uniqueConnectors := make(map[string]connectors.IConnector)
	result := make([]connectors.IConnector, 0)

	for _, config := range configs {
		if _, ok := uniqueConnectors[config.Name]; ok {
			continue
		}

		s, err := connectors.New(config, channel)
		if err != nil {
			return nil, err
		}

		uniqueConnectors[config.Name] = s
		result = append(result, s)
	}

	return result, nil
}

func RunConnectors(connectors []connectors.IConnector) error {
	for _, connector := range connectors {
		if err := connector.Run(); err != nil {
			return err
		}
	}
	return nil
}
