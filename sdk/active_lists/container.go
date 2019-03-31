package activeLists

import (
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
)

type Container struct {
	lists   map[string]Config
	storage IStorage
}

func (r *Container) Get(listName, field string, keyFields []string, event *normalization.Event) string {
	data, err := r.storage.Get(listName, helpers.MakeKey(keyFields, event), field)
	if err != nil {
		return ""
	}
	return string(data)
}

func (r *Container) Set(listName string, keyFields []string, mapping []Mapping, event *normalization.Event) {
	key := helpers.MakeKey(keyFields, event)
	ttl := r.lists[listName]
	data := make(map[string]interface{})

	for _, m := range mapping {
		if m.Constant != nil {
			data[m.ALField] = normalization.ToString(m.Constant)
		} else {
			data[m.ALField] = normalization.ToString(event.GetAnyField(m.EventField))
		}
	}

	_ = r.storage.Put(listName, key, data, ttl.TTLDuration())
}

func (r *Container) Del(listName string, keyFields []string, event *normalization.Event) {
	_ = r.storage.Del(listName, helpers.MakeKey(keyFields, event))
}

func NewContainer(lists []Config, storageURL string) (*Container, error) {
	storage, err := newRedisStorage(storageURL)
	if err != nil {
		return nil, err
	}

	container := &Container{
		storage: storage,
		lists:   make(map[string]Config),
	}

	for _, listCfg := range lists {
		container.lists[listCfg.Name] = listCfg
	}

	return container, nil
}
