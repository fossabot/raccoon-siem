package dictionaries

type Dict map[interface{}]interface{}

type Config struct {
	Name string `yaml:"name"`
	Data Dict   `yaml:"data"`
}
