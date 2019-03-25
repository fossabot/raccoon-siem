package normalizers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers/cef"
)

type cefNormalizer struct {
	name    string
	mapping []MappingConfig
}

func (r *cefNormalizer) ID() string {
	return r.name
}

func (r *cefNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	parsingResult, ok := cef.Parse(data)
	if !ok || len(parsingResult) == 0 {
		return event
	}
	return normalize(parsingResult, r.mapping, event)
}

func newCEFNormalizer(cfg Config) (*cefNormalizer, error) {
	return &cefNormalizer{
		name:    cfg.Name,
		mapping: cfg.Mapping,
	}, nil
}
