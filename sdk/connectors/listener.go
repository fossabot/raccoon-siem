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

func (r *listenerConnector) Run() error {
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
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func newListenerConnector(cfg Config, channel OutputChannel) (*listenerConnector, error) {
	switch cfg.Protocol {
	case "tcp", "udp":
	default:
		return nil, fmt.Errorf("unknown protocol: %s", cfg.Protocol)
	}

	return &listenerConnector{
		name:       cfg.Name,
		url:        cfg.URL,
		protocol:   cfg.Protocol,
		delimiter:  cfg.Delimiter,
		bufferSize: cfg.BufferSize,
		channel:    channel,
	}, nil
}
