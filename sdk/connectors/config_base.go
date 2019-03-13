package connectors

type BaseConfig struct {
	Name          string
	URL           string
	OutputChannel chan []byte
}
