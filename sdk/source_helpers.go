package sdk

func RegisterSources(sourceSettings []SourceSettings, processorChannel chan *ProcessorTask) ([]ISource, error) {
	uniqueSources := make(map[string]ISource)
	result := make([]ISource, 0, len(sourceSettings))

	for i := range sourceSettings {
		settings := &sourceSettings[i]

		if _, ok := uniqueSources[settings.Name]; ok {
			continue
		}

		s, err := NewSource(settings, processorChannel)

		if err != nil {
			return nil, err
		}

		uniqueSources[settings.Name] = s
		result = append(result, s)
	}

	return result, nil
}

func RunSources(sources []ISource) error {
	for _, src := range sources {
		if err := src.Run(); err != nil {
			return err
		}
	}
	return nil
}
