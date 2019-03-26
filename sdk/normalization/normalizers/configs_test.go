package normalizers

import (
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
						"eventField": "OriginHost",
						"sourceField": "host"
					}]
				}
			}] 
		}]
	}`)
)

func TestConfigUnmarshal(t *testing.T) {
	cfg := Config{}
	assert.Error(t, cfg.Unmarshal(withoutName), "name required")
	assert.Error(t, cfg.Unmarshal(withoutKind), "kind required")
	assert.Error(t, cfg.Unmarshal(withoutExpressions), "expressions required")
	assert.Error(t, cfg.Unmarshal(withoutDelimiters), "delimiters required")
	assert.Error(t, cfg.Unmarshal(withoutDelimiters), "delimiters required")
	assert.Error(t, cfg.Unmarshal(withoutMapping), "mapping required")
	assert.Error(t, cfg.Unmarshal(withoutSourceField), "source field required")
	assert.Error(t, cfg.Unmarshal(withoutEventField), "event field required")
	assert.Error(t, cfg.Unmarshal(withoutTriggerField), "trigger field required")
	assert.NilError(t, cfg.Unmarshal(full))

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