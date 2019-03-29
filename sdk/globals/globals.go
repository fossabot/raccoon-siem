package globals

import (
	"github.com/tephrocactus/raccoon-siem/sdk/active_lists"
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
	"github.com/tephrocactus/raccoon-siem/sdk/notifier"
)

var (
	Dictionaries *dictionaries.Storage
	ActiveLists  *activeLists.Container
	Notifier     *notifier.Notifier
)
