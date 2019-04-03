package normalizers

import (
	"encoding/json"
	"gotest.tools/assert"
	"testing"
)

var (
	withoutName = []byte("{}")
	withoutKind = []byte(`{
		"name": "general",
		"kind": "",
		"expressions": [
			"\\w",
			"\\W"
		]
	}`)
	withoutExpressions = []byte(`{
		"name": "general",
		"kind": "regexp"
	}`)
	withoutDelimiters = []byte(`{
		"name": "general",
		"kind": "kv"
	}`)
	withoutMapping = []byte(`{
		"name": "general",
		"kind": "syslog"
	}`)
	withoutSourceField = []byte(`{
		"name": "general",
		"kind": "syslog",
		"mapping": [{
			"eventField": "OriginIPAddress"
		}]
	}`)
	withoutEventField = []byte(`{
		"name": "general",
		"kind": "syslog",
		"mapping": [{
			"sourceField": "ip"
		}]
	}`)
	withoutTriggerField = []byte(`{
		"name": "general",
		"kind": "syslog",
		"mapping": [{
			"eventField": "OriginIPAddress",
			"sourceField": "ip",
			"extra": [{
				"triggerValue": "debug"
			}] 
		}]
	}`)
	full = []byte(`{
		"name": "general",
		"kind": "syslog",
		"mapping": [{
			"eventField": "OriginIPAddress",
			"sourceField": "ip",
			"extra": [{
				"triggerField": "OriginSeverity",
				"triggerValue": "debug",
				"normalizer": {
					"name": "json-ip",
					"kind": "json",
					"mapping": [{
						"eventField": "OriginDomain",
						"sourceField": "host"
					}]
				}
			}] 
		}]
	}`)
)

func TestConfigUnmarshal(t *testing.T) {
	cfg := Config{}
	_ = json.Unmarshal(withoutName, &cfg)
	assert.Error(t, cfg.Validate(), "normalizer: name required")

	cfg = Config{}
	_ = json.Unmarshal(withoutKind, &cfg)
	assert.Error(t, cfg.Validate(), "normalizer: unknown kind ")
	cfg = Config{}
	_ = json.Unmarshal(withoutExpressions, &cfg)
	assert.Error(t, cfg.Validate(), "normalizer: expressions required")
	cfg = Config{}
	_ = json.Unmarshal(withoutDelimiters, &cfg)
	assert.Error(t, cfg.Validate(), "normalizer: delimiters required")
	cfg = Config{}
	_ = json.Unmarshal(withoutDelimiters, &cfg)
	assert.Error(t, cfg.Validate(), "normalizer: delimiters required")
	cfg = Config{}
	_ = json.Unmarshal(withoutMapping, &cfg)
	assert.Error(t, cfg.Validate(), "normalizer: mapping required")
	cfg = Config{}
	_ = json.Unmarshal(withoutSourceField, &cfg)
	assert.Error(t, cfg.Validate(), "mapping: source field required")
	cfg = Config{}
	_ = json.Unmarshal(withoutEventField, &cfg)
	assert.Error(t, cfg.Validate(), "mapping: invalid event field ")
	cfg = Config{}
	_ = json.Unmarshal(withoutTriggerField, &cfg)
	assert.Error(t, cfg.Validate(), "extra: trigger field required")
	cfg = Config{}
	_ = json.Unmarshal(full, &cfg)
	assert.NilError(t, cfg.Validate())

	assert.Equal(t, cfg.Name, "general")
	assert.Equal(t, cfg.Kind, "syslog")
	assert.Equal(t, len(cfg.Mapping), 1)
	assert.Equal(t, cfg.Mapping[0].SourceField, "ip")
	assert.Equal(t, cfg.Mapping[0].EventField, "OriginIPAddress")
	assert.Equal(t, len(cfg.Mapping[0].Extra), 1)
	assert.Equal(t, cfg.Mapping[0].Extra[0].TriggerField, "OriginSeverity")
	assert.Equal(t, cfg.Mapping[0].Extra[0].TriggerValue, "debug")
	assert.Equal(t, cfg.Mapping[0].Extra[0].Normalizer.Name, "json-ip")
}