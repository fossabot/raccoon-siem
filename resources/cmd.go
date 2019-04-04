package resources

import (
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/resources/create"
	"github.com/tephrocactus/raccoon-siem/resources/delete"
	"github.com/tephrocactus/raccoon-siem/resources/generate"
	"github.com/tephrocactus/raccoon-siem/resources/get"
	"github.com/tephrocactus/raccoon-siem/resources/list"
)

var (
	Cmd = &cobra.Command{
		Use:   "resources",
		Short: "manage raccoon resources",
		Args:  cobra.ExactArgs(1),
	}
)

func init() {
	// Sub commands
	Cmd.AddCommand(
		create.Cmd,
		get.Cmd,
		list.Cmd,
		delete.Cmd,
		generate.Cmd,
	)
}
