package connectors

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"net"
)

type listenerConnector struct {
	name       string
	url        string
	protocol   string
	delimiter  byte
	bufferSize int
	channel    OutputChannel
}

func (r *listenerConnector) ID() string {
	return r.name
}

func (r *listenerConnector) Start() error {
	listener, err := net.Listen(r.protocol, r.url)
	if err != nil {
		return err
	}
	go r.acceptConnections(listener)
	return nil
}

func (r *listenerConnector) acceptConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go r.handleConnection(conn)
	}
}

func (r *listenerConnector) handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(reader)
	scanner.Split(r.framer)

	for scanner.Scan() {
		r.channel <- Output{
			Connector: r.name,
			Data:      helpers.CopyBytes(scanner.Bytes()),
		}
	}

	if err := scanner.Err(); err != nil {
		_ = conn.Close()
	}
}

func (r *listenerConnector) framer(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, r.delimiter); i >= 0 {
		// We have a full newline-terminated line.
		if r.delimiter == '\n' {
			return i + 1, dropCR(data[0:i]), nil
		}
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		if r.delimiter == '\n' {
			return len(data), dropCR(data), nil
		}
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func newListenerConnector(cfg Config, channel OutputChannel) (*listenerConnector, error) {
	switch cfg.Protocol {
	case "tcp", "udp":
	default:
		return nil, fmt.Errorf("unknown protocol: %s", cfg.Protocol)
	}

	delimiter := byte('\n')
	if cfg.Delimiter != "" {
		d, err := helpers.StringToSingleByte(cfg.Delimiter)
		if err != nil {
			return nil, err
		}
		delimiter = d
	}

	return &listenerConnector{
		name:       cfg.Name,
		url:        cfg.URL,
		protocol:   cfg.Protocol,
		delimiter:  delimiter,
		bufferSize: cfg.BufferSize,
		channel:    channel,
	}, nil
}
