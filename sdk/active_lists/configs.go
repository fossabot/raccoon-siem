package activeLists

type Config struct {
	Name   string `yaml:"name"`
	Fields []FieldConfig
}

type FieldConfig struct {
	Name   string `yaml:"name"`
	Kind   string `yaml:"kind"`
	Unique bool   `yaml:"unique"`
	PK     bool   `yaml:"pk"`
}
