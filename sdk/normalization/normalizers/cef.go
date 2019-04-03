package normalizers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers/cef"
)

type cefNormalizer struct {
	name    string
	mapping map[string]MappingConfig
	extra   []ExtraConfig
}

func (r *cefNormalizer) ID() string {
	return r.name
}

func (r *cefNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	event, created := createEventIfNil(event)
	if !cef.Parse(data, parserCallbackGenerator(r.mapping, event)) {
		return eventOrNil(event, created)
	}
	return extraNormalize(event, r.extra)
}

func newCEFNormalizer(cfg Config) (*cefNormalizer, error) {
	return &cefNormalizer{
		name:    cfg.Name,
		mapping: groupMappingBySourceField(cfg.Mapping),
		extra:   cfg.Extra,
	}, nil
}
