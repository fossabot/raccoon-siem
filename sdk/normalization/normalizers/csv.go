package normalizers

import (
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers/csv"
)

type csvNormalizer struct {
	name      string
	delimiter byte
	mapping   map[string]MappingConfig
	extra     []ExtraConfig
}

func (r *csvNormalizer) ID() string {
	return r.name
}

func (r *csvNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	event, created := createEventIfNil(event)
	if !csv.Parse(data, r.delimiter, parserCallbackGenerator(r.mapping, event)) {
		return eventOrNil(event, created)
	}
	return extraNormalize(event, r.extra)
}

func newCSVNormalizer(cfg Config) (*csvNormalizer, error) {
	valueDelimiter, err := helpers.StringToSingleByte(cfg.Delimiter)
	if err != nil {
		return nil, err
	}

	return &csvNormalizer{
		name:      cfg.Name,
		delimiter: valueDelimiter,
		mapping:   groupMappingBySourceField(cfg.Mapping),
		extra:     cfg.Extra,
	}, nil
}
