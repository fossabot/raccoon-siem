package connectors

type BaseConfig struct {
	Name          string
	URL           string
	OutputChannel chan []byte
}

func (r *BaseConfig) ID() string {
	return r.Name
}
