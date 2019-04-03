package normalizers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers/syslog/rfc3164"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers/syslog/rfc5424"
)

type syslogNormalizer struct {
	name    string
	mapping map[string]MappingConfig
	extra   []ExtraConfig
}

func (r *syslogNormalizer) ID() string {
	return r.name
}

func (r *syslogNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	event, created := createEventIfNil(event)
	callback := parserCallbackGenerator(r.mapping, event)

	if !rfc5424.Parse(data, callback) {
		if !rfc3164.Parse(data, callback) {
			return eventOrNil(event, created)
		}
	}

	return extraNormalize(event, r.extra)
}

func newSyslogNormalizer(cfg Config) (*syslogNormalizer, error) {
	return &syslogNormalizer{
		name:    cfg.Name,
		mapping: groupMappingBySourceField(cfg.Mapping),
		extra:   cfg.Extra,
	}, nil
}
