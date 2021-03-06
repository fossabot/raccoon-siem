package dictionaries

type Storage struct {
	data map[string]Dict
}

func NewStorage(dictionaries []Config) *Storage {
	s := &Storage{data: make(map[string]Dict)}
	for _, dict := range dictionaries {
		s.data[dict.Name] = dict.Data
	}
	return s
}

func (r *Storage) Get(dictionaryName string, key string) string {
	dictionary, found := r.data[dictionaryName]
	if !found {
		return ""
	}
	return dictionary[key]
}
