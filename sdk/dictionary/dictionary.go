package dictionary

type dictionaries map[string]map[string]string

type Config struct {
	Data dictionaries
}

type dictionary struct {
	cfg Config
}

func NewDictionary(cfg Config) *dictionary {
	return &dictionary{cfg}
}

func (r *dictionary) Get(dictionaryName string, field string) string {
	dictionary, found := r.cfg.Data[dictionaryName]
	if !found {
		return ""
	}
	return dictionary[field]
}