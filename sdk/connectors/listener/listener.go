package connectors

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"net"
)

type Config struct {
	connectors.BaseConfig
	Protocol      string
	Delimiter     byte
	BufferSize    int
	OutputChannel chan []byte
}

type connector struct {
	cfg Config
}

func (r *connector) ID() string {
	return r.cfg.Name
}

func (r *connector) Run() error {
	listener, err := net.Listen(r.cfg.Protocol, r.cfg.URL)
	if err != nil {
		return err
	}
	go r.acceptConnections(listener)
	return nil
}

func (r *connector) acceptConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go r.handleConnection(conn)
	}
}

func (r *connector) handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(reader)
	scanner.Split(r.framer)

	for scanner.Scan() {
		r.cfg.OutputChannel <- helpers.CopyBytes(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		_ = conn.Close()
	}
}

func (r *connector) framer(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, r.cfg.Delimiter); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func NewConnector(config Config) (*connector, error) {
	switch config.Protocol {
	case "tcp", "udp":
	default:
		return nil, fmt.Errorf("unknown protocol: %s", config.Protocol)
	}
	return &connector{cfg: config}, nil
}
