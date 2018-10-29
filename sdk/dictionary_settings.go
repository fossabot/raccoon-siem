package sdk

type DictionaryData map[interface{}]interface{}

type DictionarySettings struct {
	Name string         `yaml:"name,omitempty"`
	Data DictionaryData `yaml:"data,omitempty"`
}

func (s *DictionarySettings) ID() string {
	return s.Name
}

func (s *DictionarySettings) compile() (*dictionary, error) {
	return new(dictionary).compile(s)
}
