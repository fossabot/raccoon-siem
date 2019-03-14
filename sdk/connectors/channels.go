package connectors

type Output struct {
	Connector string
	Data      []byte
}

type OutputChannel chan Output
