package sdk

import (
	"fmt"
)

const (
	dictionaryActionGetString = "get"
	dictionaryActionGet       = iota
)

var knownDictionaryActions = map[string]byte{
	dictionaryActionGetString: dictionaryActionGet,
}

var dictionariesByName = make(map[string]*dictionary)

type dictionary struct {
	name string
	data DictionaryData
}

func (d *dictionary) compile(settings *DictionarySettings) (*dictionary, error) {
	d.name = settings.Name

	if d.name == "" {
		return d, fmt.Errorf("dictionary must have a name")
	}

	d.data = make(DictionaryData)

	if len(settings.Data) == 0 {
		return d, fmt.Errorf("dictionary '%s' is empty", d.name)
	}

	for k, v := range settings.Data {
		d.data[to64Bits(k)] = to64Bits(v)
	}

	return d, nil
}
