package sdk

import (
	"fmt"
	"github.com/fln/nf9packet"
	"log"
	"net"
	"strconv"
	"strings"
)

func newNetflowSource(settings *SourceSettings, processorChannel chan ProcessorTask) ISource {
	return &netflowSource{
		processorChannel: processorChannel,
		settings:         settings,
		templateCache:    make(map[string]*nf9packet.TemplateRecord),
	}
}

type netflowSource struct {
	settings         *SourceSettings
	connection       *net.UDPConn
	processorChannel chan ProcessorTask
	templateCache    map[string]*nf9packet.TemplateRecord
}

func (s *netflowSource) Run() (err error) {
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

func (s *netflowSource) handleData() {
	bufSize := 8192

	if s.settings.Buffer > 0 {
		bufSize = s.settings.Buffer
	}

	buf := make([]byte, bufSize)

	for {
		length, remote, err := s.connection.ReadFrom(buf)

		if err != nil {
			if Debug {
				log.Println(err)
			}
			continue
		}

		if length == 0 {
			continue
		}

		s.process(buf[:length], remote.String())
	}
}

func (s *netflowSource) process(input []byte, remote string) error {
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

			s.processorChannel <- []byte(sb.String())
		}
	}

	return nil
}
