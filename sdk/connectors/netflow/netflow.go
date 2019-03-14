package netflow

import (
	"fmt"
	"github.com/fln/nf9packet"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"net"
	"strconv"
	"strings"
)

type Config struct {
	connectors.BaseConfig
	BufferSize int
}

type connector struct {
	cfg           Config
	templateCache map[string]*nf9packet.TemplateRecord
}

func (r *connector) ID() string {
	return r.cfg.Name
}

func (r *connector) Run() (err error) {
	addr, err := net.ResolveUDPAddr("udp", r.cfg.URL)
	if err != nil {
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return
	}

	go r.handleData(conn)
	return
}

func (r *connector) handleData(conn *net.UDPConn) {
	bufSize := 8192
	if r.cfg.BufferSize > 0 {
		bufSize = r.cfg.BufferSize
	}

	buf := make([]byte, bufSize)
	for {
		length, remote, err := conn.ReadFrom(buf)

		if err != nil {
			continue
		}

		if length == 0 {
			continue
		}

		_ = r.process(buf[:length], remote.String())
	}
}

func (r *connector) process(input []byte, remote string) error {
	pkt, err := nf9packet.Decode(input)

	if err != nil {
		return err
	}

	for _, t := range pkt.TemplateRecords() {
		templateKey := fmt.Sprintf("%s|%b|%v", remote, pkt.SourceId, t.TemplateId)
		r.templateCache[templateKey] = t
	}

	for _, set := range pkt.DataFlowSets() {
		templateKey := fmt.Sprintf("%s|%b|%v", remote, pkt.SourceId, set.Id)
		template, ok := r.templateCache[templateKey]

		if !ok {
			continue
		}

		records := template.DecodeFlowSet(&set)

		for _, rec := range records {
			sb := strings.Builder{}

			for i := range rec.Values {
				fName := template.Fields[i].Name()

				sb.WriteString(fName)
				sb.WriteString("=")

				if fName == "PROTOCOL" {
					protoUint := template.Fields[i].DataToUint64(rec.Values[i])
					sb.WriteString(strconv.FormatUint(protoUint, 10))
				} else {
					sb.WriteString(template.Fields[i].DataToString(rec.Values[i]))
				}

				sb.WriteString(" ")
			}

			r.cfg.OutputChannel <- connectors.Output{
				Connector: r.cfg.Name,
				Data:      []byte(sb.String()),
			}
		}
	}

	return nil
}

func NewConnector(config Config) (*connector, error) {
	return &connector{cfg: config}, nil
}
