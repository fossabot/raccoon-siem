package globals

import (
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/dictionary"
)

var (
	DictionaryStorage *dictionary.Storage
	ALStorage         *activeLists.Storage
)
