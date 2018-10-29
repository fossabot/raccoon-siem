package sdk

import (
	"github.com/mediocregopher/radix.v3"
	"time"
)

var activeListsService *radix.Pool
var activeListsByName = make(map[string]*activeList)

type alMultiValueType map[string]interface{}

var removeExpiredKeysFromSetScript = radix.NewEvalScript(1, `
	local keysInSet = redis.pcall('ZRANGE', KEYS[1], 0, -1)
	for k, v in pairs(keysInSet) do
		local exists = redis.pcall('EXISTS', v)
		if exists == 0 then
			redis.pcall('ZREM', KEYS[1], v)
		end
	end
`)

func newActiveList(spec *activeListSpecification) *activeList {
	return &activeList{
		spec:    spec,
		service: activeListsService,
	}
}

type activeListSpecification struct {
	name string
	ttl  int64
}

type activeList struct {
	spec    *activeListSpecification
	service *radix.Pool
}

func (al *activeList) Run() {
	go func() {
		for {
			err := al.service.Do(removeExpiredKeysFromSetScript.Cmd(nil, al.spec.name))

			if err != nil {
				DebugError(err)
			}

			time.Sleep(15 * time.Second)
		}
	}()
}

func (al *activeList) Add(key string, value alMultiValueType) error {
	return al.service.Do(radix.Pipeline(
		radix.FlatCmd(nil, "HMSET", key, value),
		radix.FlatCmd(nil, "EXPIRE", key, al.spec.ttl),
		radix.FlatCmd(nil, "ZADD", al.spec.name, "INCR", 1, key),
	))
}

func (al *activeList) Get(key string) (value alMultiValueType, err error) {
	value = make(alMultiValueType)
	err = al.service.Do(radix.FlatCmd(&value, "HGETALL", key))
	return
}

func (al *activeList) Delete(key string) error {
	return al.service.Do(radix.Pipeline(
		radix.FlatCmd(nil, "DEL", key),
		radix.FlatCmd(nil, "ZREM", al.spec.name, key),
	))
}

func (al *activeList) Size() (size int64, err error) {
	err = al.service.Do(radix.FlatCmd(&size, "ZCARD", al.spec.name))
	return
}

func (al *activeList) Count(key string) (count int64, err error) {
	err = al.service.Do(radix.FlatCmd(&count, "ZSCORE", al.spec.name, key))
	return
}

func (al *activeList) Clear() error {
	records := make([]string, 0)
	err := al.service.Do(radix.FlatCmd(&records, "ZRANGE", al.spec.name, 0, -1))

	if err != nil {
		return err
	}

	return al.service.Do(radix.FlatCmd(nil, "DEL", al.spec.name, records))
}
