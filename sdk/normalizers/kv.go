package normalizers

import (
	"errors"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"github.com/tephrocactus/raccoon-siem/sdk/normalizers/parsers/kv"
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
	if !ok {
		return event
	}
	return normalize(parsingResult, r.mapping, event)
}

func newKVNormalizer(cfg Config) (*kvNormalizer, error) {
	if cfg.PairDelimiter == 0 {
		return nil, errors.New("pair separators cannot be empty")
	}

	if cfg.KVDelimiter == 0 {
		return nil, errors.New("kv separators cannot be empty")
	}

	if cfg.PairDelimiter == cfg.KVDelimiter {
		return nil, errors.New("kv and pair separators must be different")
	}

	return &kvNormalizer{
		name:          cfg.Name,
		pairDelimiter: cfg.PairDelimiter,
		kvDelimiter:   cfg.KVDelimiter,
	}, nil
}
