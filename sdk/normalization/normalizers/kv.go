package normalizers

import (
	"errors"
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers/kv"
)

type kvNormalizer struct {
	name          string
	pairDelimiter byte
	kvDelimiter   byte
	mapping       map[string]MappingConfig
	extra         []ExtraConfig
}

func (r *kvNormalizer) ID() string {
	return r.name
}

func (r *kvNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	event, created := createEventIfNil(event)
	if !kv.Parse(data, r.pairDelimiter, r.kvDelimiter, parserCallbackGenerator(r.mapping, event)) {
		return eventOrNil(event, created)
	}
	return extraNormalize(event, r.extra)
}

func newKVNormalizer(cfg Config) (*kvNormalizer, error) {
	if cfg.PairDelimiter == cfg.KVDelimiter {
		return nil, errors.New("kv and pair separators must be different")
	}

	pairDelimiter, err := helpers.StringToSingleByte(cfg.PairDelimiter)
	if err != nil {
		return nil, err
	}

	kvDelimiter, err := helpers.StringToSingleByte(cfg.KVDelimiter)
	if err != nil {
		return nil, err
	}

	return &kvNormalizer{
		name:          cfg.Name,
		pairDelimiter: pairDelimiter,
		kvDelimiter:   kvDelimiter,
		mapping:       groupMappingBySourceField(cfg.Mapping),
		extra:         cfg.Extra,
	}, nil
}
