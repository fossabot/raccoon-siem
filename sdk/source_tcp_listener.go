package sdk

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func newTCPListenerSource(settings *SourceSettings, processorChannel chan *ProcessorTask) ISource {
	return &tcpListenerSource{
		processorChannel: processorChannel,
		settings:         settings,
	}
}

type tcpListenerSource struct {
	settings         *SourceSettings
	listener         net.Listener
	processorChannel chan *ProcessorTask
}

func (s *tcpListenerSource) ID() string {
	return s.settings.Name
}

func (s *tcpListenerSource) Run() (err error) {
	s.listener, err = net.Listen("tcp", s.settings.URL)

	if err != nil {
		return
	}

	go s.acceptConnections()

	return
}

func (s *tcpListenerSource) acceptConnections() {
	for {
		conn, err := s.listener.Accept()

		if err != nil {
			if Debug {
				fmt.Println(err)
			}
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *tcpListenerSource) handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s.processorChannel <- &ProcessorTask{
			Source: s.settings.Name,
			Data:   CopyBytes(scanner.Bytes()),
		}
	}

	if err := scanner.Err(); err != nil {
		if Debug {
			log.Println(err)
		}
		conn.Close()
	}
}
