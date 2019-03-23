package normalizers

import (
	"fmt"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

const (
	KindJSON   = "json"
	KindNative = "native"
	KindCEF    = "cef"
	KindRegexp = "regexp"
	KindSyslog = "syslog"
	KindKV     = "kv"
)

type INormalizer interface {
	ID() string
	Normalize(data []byte, event *normalization.Event) *normalization.Event
}

func New(cfg Config) (INormalizer, error) {
	if err := initExtraNormalizers(cfg.Mapping); err != nil {
		return nil, err
	}

	switch cfg.Kind {
	case KindSyslog:
		return newSyslogNormalizer(cfg)
	case KindJSON:
		return newJSONNormalizer(cfg)
	case KindCEF:
		return newCEFNormalizer(cfg)
	case KindRegexp:
		return newRegexpNormalizer(cfg)
	case KindKV:
		return newKVNormalizer(cfg)
	}

	panic(fmt.Errorf("unknown normalizer kind: %s", cfg.Kind))
}

func initExtraNormalizers(mapping []MappingConfig) (err error) {
	for m := range mapping {
		for e := range mapping[m].Extra {
			mapping[m].Extra[e].triggerValue = []byte(mapping[m].Extra[e].TriggerValue)
			mapping[m].Extra[e].normalizer, err = New(mapping[m].Extra[e].Normalizer)
			if err != nil {
				return err
			}
		}
	}
	return err
}
