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
	default:
		panic(fmt.Errorf("unknown normalizer kind: %s", cfg.Kind))
	}
}
