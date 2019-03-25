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
	mapping       []MappingConfig
}

func (r *kvNormalizer) ID() string {
	return r.name
}

func (r *kvNormalizer) Normalize(data []byte, event *normalization.Event) *normalization.Event {
	parsingResult, ok := kv.Parse(data, r.pairDelimiter, r.kvDelimiter)
	if !ok || len(parsingResult) == 0 {
		return event
	}
	return normalize(parsingResult, r.mapping, event)
}

func newKVNormalizer(cfg Config) (*kvNormalizer, error) {
	pairDelimiter, err := helpers.StringToSingleByte(cfg.PairDelimiter)
	if err != nil {
		return nil, err
	}

	kvDelimiter, err := helpers.StringToSingleByte(cfg.KVDelimiter)
	if err != nil {
		return nil, err
	}

	if cfg.PairDelimiter == cfg.KVDelimiter {
		return nil, errors.New("kv and pair separators must be different")
	}

	return &kvNormalizer{
		name:          cfg.Name,
		pairDelimiter: pairDelimiter,
		kvDelimiter:   kvDelimiter,
	}, nil
}
