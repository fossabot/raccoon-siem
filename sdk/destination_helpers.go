package sdk

func RegisterDestinations(destinationSettings []DestinationSettings) ([]IDestination, error) {
	uniqueDestinations := make(map[string]IDestination)
	result := make([]IDestination, 0, len(destinationSettings))

	for _, settings := range destinationSettings {
		if _, ok := uniqueDestinations[settings.Name]; ok {
			continue
		}

		dst, err := NewDestination(settings)

		if err != nil {
			return nil, err
		}

		uniqueDestinations[settings.Name] = dst
		result = append(result, dst)
	}

	return result, nil
}

func RunDestinations(destinations []IDestination) error {
	for _, dst := range destinations {
		if err := dst.Run(); err != nil {
			return err
		}
	}
	return nil
}

func GetDefaultDestinationSettings(elasticURL string, natsURL string) (settings []DestinationSettings) {
	if elasticURL != "" {
		storageSettings := DestinationSettings{
			Name:  RaccoonStorageName,
			Kind:  RaccoonStorageKind,
			Index: RaccoonStorageIndex,
			URL:   elasticURL,
		}
		settings = append(settings, storageSettings)
	}

	correlationBusSettings := DestinationSettings{
		Name:    RaccoonCorrelationBusName,
		Kind:    RaccoonCorrelationBusKind,
		Channel: RaccoonCorrelationBusChannel,
		URL:     natsURL,
	}

	settings = append(settings, correlationBusSettings)
	return
}
