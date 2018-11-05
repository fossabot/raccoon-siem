package sdk

import (
	"bufio"
	"log"
	"net"
)

func newUDPListenerSource(settings *SourceSettings, processorChannel chan *ProcessorTask) ISource {
	return &udpListenerSource{
		processorChannel: processorChannel,
		settings:         settings,
	}
}

type udpListenerSource struct {
	settings         *SourceSettings
	connection       *net.UDPConn
	processorChannel chan *ProcessorTask
}

func (s *udpListenerSource) ID() string {
	return s.settings.Name
}

func (s *udpListenerSource) Run() (err error) {
	addr, err := net.ResolveUDPAddr("udp", s.settings.URL)

	if err != nil {
		return
	}

	s.connection, err = net.ListenUDP("udp", addr)

	if err != nil {
		return
	}

	go s.handleData()

	return
}

func (s *udpListenerSource) handleData() {
	reader := bufio.NewReader(s.connection)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for {
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
		}
	}
}
