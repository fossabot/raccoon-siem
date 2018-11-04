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
)

func init() {

}

func runCore(_ *cobra.Command, _ []string) error {
	installer, err := newInstaller()
	if err != nil {
		return err
	}

	installer.Install(componentCore)

	return nil
}
