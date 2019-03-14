package connectors

import (
	"fmt"
	"github.com/fln/nf9packet"
	"net"
	"strconv"
	"strings"
)

type netflowConnector struct {
	name          string
	url           string
	bufferSize    int
	channel       OutputChannel
	templateCache map[string]*nf9packet.TemplateRecord
}

func (r *netflowConnector) ID() string {
	return r.name
}

func (r *netflowConnector) Run() (err error) {
	addr, err := net.ResolveUDPAddr("udp", r.url)
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

func (r *netflowConnector) handleData(conn *net.UDPConn) {
	bufSize := 8192
	if r.bufferSize > 0 {
		bufSize = r.bufferSize
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

func (r *netflowConnector) process(input []byte, remote string) error {
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

			r.channel <- Output{
				Connector: r.name,
				Data:      []byte(sb.String()),
			}
		}
	}

	return nil
}

func newNetflowConnector(cfg Config, channel OutputChannel) (*netflowConnector, error) {
	return &netflowConnector{
		name:       cfg.Name,
		url:        cfg.URL,
		bufferSize: cfg.BufferSize,
		channel:    channel,
	}, nil
}
