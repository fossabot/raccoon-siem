package install

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "install",
		Short: "install raccoon components",
		Args:  cobra.ExactArgs(1),
		ValidArgs: []string{
			componentCore,
			componentCollector,
			componentCorrelator,
			componentBus,
			componentAL,
			componentStorage,
		},
		RunE: cobra.OnlyValidArgs,
	}

	raccoonBinaryName      = "raccoon"
	raccoonCoreServiceName = "raccoon_core"

	componentCore       = "core"
	componentCollector  = "collector"
	componentCorrelator = "correlator"
	componentBus        = "bus"
	componentAL         = "al"
	componentStorage    = "storage"
)

func init() {
	Cmd.AddCommand(coreCmd)
}
