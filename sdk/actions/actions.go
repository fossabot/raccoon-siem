package actions

import (
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/enrichment"
	"github.com/tephrocactus/raccoon-siem/sdk/globals"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

func Release(cfg ReleaseConfig, correlationEvent *normalization.Event, baseEvents ...*normalization.Event) {
	for _, cfg := range cfg.EnrichmentConfigs {
		enrichment.Enrich(cfg, correlationEvent, baseEvents...)
	}
}

func ActiveList(cfg ActiveListConfig, events ...*normalization.Event) {
	var targetEvent *normalization.Event

	if cfg.EventTag != "" {
		for _, event := range events {
			if event.Tag == cfg.EventTag {
				targetEvent = event
				break
			}
		}
	} else if len(events) > 0 {
		targetEvent = events[0]
	}

	if targetEvent == nil {
		return
	}

	switch cfg.Op {
	case activeLists.OpDel:
		globals.ActiveLists.Del(cfg.Name, cfg.KeyFields, targetEvent)
	case activeLists.OpSet:
		globals.ActiveLists.Set(cfg.Name, cfg.KeyFields, cfg.Mapping, targetEvent)
	}
}
