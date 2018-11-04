package install

import (
	"github.com/spf13/cobra"
)

var (
	coreCmd = &cobra.Command{
		Use:   "core",
		Short: "install raccoon core",
		Args:  cobra.ExactArgs(0),
		RunE:  runCore,
	}

	// String flags variables
	flagCoreListenAddress string
)

func init() {
	// Listen address
	Cmd.Flags().StringVarP(
		&flagCoreListenAddress,
		"listen",
		"l",
		":7220",
		"listen address")
}

func runCore(_ *cobra.Command, _ []string) error {
	installer, err := newInstaller()
	if err != nil {
		return err
	}
	return installer.Install(componentCore)
}
