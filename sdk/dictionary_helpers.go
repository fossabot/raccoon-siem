package sdk

func RegisterDictionaries(settings []DictionarySettings) error {
	for _, setting := range settings {
		dict, err := setting.compile()

		if err != nil {
			return err
		}

		if _, ok := dictionariesByName[dict.name]; ok {
			continue
		}

		dictionariesByName[dict.name] = dict
	}

	return nil
}
