package normalizers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers/parsers/json"
)

type jsonNormalizer struct {
	name    string
	mapping []MappingConfig
}

func (r *jsonNormalizer) ID() string {
	return r.name
}

func (r *jsonNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	parsingResult, ok := json.Parse(data)
	if !ok {
		return event
	}
	return normalize(parsingResult, r.mapping, event)
}

func newJSONNormalizer(cfg Config) (*jsonNormalizer, error) {
	return &jsonNormalizer{
		name:    cfg.Name,
		mapping: cfg.Mapping,
	}, nil
}
