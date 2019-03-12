package sdk

import (
	"fmt"
	"github.com/fln/nf9packet"
	"log"
	"net"
	"strconv"
	"strings"
)

type NetflowConnectorConfig struct {
	BaseConnectorConfig

	BufferSize int
}

type netflowConnector struct {
	config        NetflowConnectorConfig
	templateCache map[string]*nf9packet.TemplateRecord
}

func (s *netflowConnector) ID() string {
	return s.config.Name
}

func (s *netflowConnector) Run() (err error) {
	addr, err := net.ResolveUDPAddr("udp", s.config.URL)
	if err != nil {
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return
	}

	go s.handleData(conn)
	return
}

func (s *netflowConnector) handleData(conn *net.UDPConn) {
	bufSize := 8192
	if s.config.BufferSize > 0 {
		bufSize = s.config.BufferSize
	}

	buf := make([]byte, bufSize)
	for {
		length, remote, err := conn.ReadFrom(buf)

		if err != nil {
			if Debug {
				log.Println(err)
			}
			continue
		}

		if length == 0 {
			continue
		}

		_ = s.process(buf[:length], remote.String())
	}
}

func (s *netflowConnector) process(input []byte, remote string) error {
	pkt, err := nf9packet.Decode(input)

	if err != nil {
		return err
	}

	for _, t := range pkt.TemplateRecords() {
		templateKey := fmt.Sprintf("%s|%b|%v", remote, pkt.SourceId, t.TemplateId)
		s.templateCache[templateKey] = t
	}

	for _, set := range pkt.DataFlowSets() {
		templateKey := fmt.Sprintf("%s|%b|%v", remote, pkt.SourceId, set.Id)
		template, ok := s.templateCache[templateKey]

		if !ok {
			continue
		}

		records := template.DecodeFlowSet(&set)

		for _, r := range records {
			sb := strings.Builder{}

			for i := range r.Values {
				fName := template.Fields[i].Name()

				sb.WriteString(fName)
				sb.WriteString("=")

				if fName == "PROTOCOL" {
					protoUint := template.Fields[i].DataToUint64(r.Values[i])
					sb.WriteString(strconv.FormatUint(protoUint, 10))
				} else {
					sb.WriteString(template.Fields[i].DataToString(r.Values[i]))
				}

				sb.WriteString(" ")
			}

			s.config.OutputChannel <- &ProcessorTask{
				Connector: s.config.Name,
				Data:      []byte(sb.String()),
			}
		}
	}

	return nil
}

func newNetflowConnector(config NetflowConnectorConfig) (IConnector, error) {
	return &netflowConnector{config: config}, nil
}
