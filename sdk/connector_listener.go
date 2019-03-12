package sdk

import (
	"bufio"
	"fmt"
	"net"
)

type ListenerConnectorConfig struct {
	BaseConnectorConfig
	Protocol      string
	Delimiter     string
	BufferSize    int
	OutputChannel chan *ProcessorTask
}

type ListenerConnector struct {
	config ListenerConnectorConfig
}

func (r *ListenerConnector) ID() string {
	return r.config.Name
}

func (r *ListenerConnector) Run() error {
	listener, err := net.Listen(r.config.Protocol, r.config.URL)
	if err != nil {
		return err
	}
	go r.acceptConnections(listener)
	return nil
}

func (r *ListenerConnector) acceptConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go r.handleConnection(conn)
	}
}

func (r *ListenerConnector) handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		r.config.OutputChannel <- &ProcessorTask{
			Connector: r.config.Name,
			Data:      CopyBytes(scanner.Bytes()),
		}
	}

	if err := scanner.Err(); err != nil {
		_ = conn.Close()
	}
}

func newListenerConnector(config ListenerConnectorConfig) (IConnector, error) {
	switch config.Protocol {
	case "tcp", "udp":
	default:
		return nil, fmt.Errorf("unknown protocol: %s", config.Protocol)
	}
	return &ListenerConnector{config: config}, nil
}
