package globals

import (
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
)

var (
	Dictionaries *dictionaries.Storage
	ActiveLists  *activeLists.Container
)
