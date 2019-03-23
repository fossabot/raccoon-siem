package normalizers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers/syslog/rfc3164"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers/syslog/rfc5424"
)

type syslogNormalizer struct {
	name    string
	mapping []MappingConfig
}

func (r *syslogNormalizer) ID() string {
	return r.name
}

func (r *syslogNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	parsingResult, ok := rfc5424.Parse(data)
	if !ok || len(parsingResult) == 0 {
		parsingResult, ok = rfc3164.Parse(data)
		if !ok || len(parsingResult) == 0 {
			return event
		}
	}
	return normalize(parsingResult, r.mapping, event)
}

func newSyslogNormalizer(cfg Config) (*syslogNormalizer, error) {
	return &syslogNormalizer{
		name:    cfg.Name,
		mapping: cfg.Mapping,
	}, nil
}
