package dictionaries

type Dict map[string]interface{}

type Config struct {
	Name string `yaml:"name"`
	Data Dict   `yaml:"data"`
}
