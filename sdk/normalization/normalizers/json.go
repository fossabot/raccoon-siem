package normalizers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers/json"
)

type jsonNormalizer struct {
	name    string
	mapping map[string]MappingConfig
	extra   []ExtraConfig
}

func (r *jsonNormalizer) ID() string {
	return r.name
}

func (r *jsonNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	event, created := createEventIfNil(event)
	if !json.Parse(data, parserCallbackGenerator(r.mapping, event)) {
		return eventOrNil(event, created)
	}
	return extraNormalize(event, r.extra)
}

func newJSONNormalizer(cfg Config) (*jsonNormalizer, error) {
	return &jsonNormalizer{
		name:    cfg.Name,
		mapping: groupMappingBySourceField(cfg.Mapping),
		extra:   cfg.Extra,
	}, nil
}
