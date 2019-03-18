package dictionary

type dictionaries map[string]map[string]string

type Config struct {
	Data dictionaries
}

type Storage struct {
	cfg Config
}

func NewDictionaryStorage(cfg Config) *Storage {
	return &Storage{cfg}
}

func (r *Storage) Get(dictionaryName string, field string) string {
	dictionary, found := r.cfg.Data[dictionaryName]
	if !found {
		return ""
	}
	return dictionary[field]
}