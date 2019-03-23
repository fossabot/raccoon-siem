package dictionary

import (
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
	"gotest.tools/assert"
	"testing"
)

const (
	raccoon = "raccoon"
	weird   = "weird"
)

func TestDictionary(t *testing.T) {
	dictionariesData := make(map[string]map[string]string)

	raccoonDict := make(map[string]string)
	raccoonDict["ERROR"] = "error"
	raccoonDict["DEBUG"] = "debug"
	raccoonDict["INFO"] = "info"

	weirdDict := make(map[string]string)
	weirdDict["1"] = "error"
	weirdDict["2"] = "debug"
	weirdDict["3"] = "info"

	dictionariesData[raccoon] = raccoonDict
	dictionariesData[weird] = weirdDict

	dictionaries := dictionaries.NewStorage(dictionaries.Config{Data: dictionariesData})

	assert.Equal(t, dictionaries.Get(raccoon, "ERROR"), "error")
	assert.Equal(t, dictionaries.Get(weird, "2"), "debug")
}
