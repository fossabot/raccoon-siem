package parsers

type BaseConfig struct {
	Name string
}

func (r *BaseConfig) ID() string {
	return r.Name
}
