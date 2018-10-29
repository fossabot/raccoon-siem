package sdk

import (
	"github.com/mediocregopher/radix.v3"
	"strings"
)

func RegisterActiveLists(svcSettings ActiveListServiceSettings, settings []ActiveListSettings) error {
	if svcSettings.PoolSize == 0 {
		svcSettings.PoolSize = 10
	}

	pool, err := radix.NewPool("tcp", svcSettings.URL, svcSettings.PoolSize)

	if err != nil {
		return err
	}

	activeListsService = pool

	for _, setting := range settings {
		spec, err := setting.compile()

		if err != nil {
			return err
		}

		al := newActiveList(spec)

		if _, ok := activeListsByName[al.spec.name]; ok {
			continue
		}

		al.Run()

		activeListsByName[al.spec.name] = al
	}

	return nil
}

func makeActiveListKey(alName string, keyFields []string, event *Event) string {
	sb := strings.Builder{}
	sb.WriteString(alName + ":")

	for _, key := range keyFields {
		sb.WriteString(toString(event.GetFieldNoType(key)))
		sb.WriteString(".")
	}

	return sb.String()
}
