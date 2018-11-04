package resources

import (
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var (
	Cmd = &cobra.Command{
		Use:       "resources",
		Short:     "manage raccoon resources",
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"create", "delete", "get", "list"},
		RunE:      cobra.OnlyValidArgs,
	}

	validResourceKinds = []string{
		"collector",
		"correlator",
		"source",
		"destination",
		"filter",
		"parser",
		"aggregationRule",
		"correlationRule",
		"dictionary",
		"activeList",
	}

	// String flags variables
	coreURL, sourcePath string

	// Other variables
	httpClient = http.Client{Timeout: 16 * time.Second}
)

func init() {
	// Sub commands
	Cmd.AddCommand(createCmd, deleteCmd, getCmd, listCmd)

	// Raccoon core URL
	Cmd.PersistentFlags().StringVarP(
		&coreURL,
		"core",
		"c",
		"http://localhost:7220",
		"raccoon core URL")

	// Source directory or file
	Cmd.PersistentFlags().StringVarP(
		&sourcePath,
		"from",
		"f",
		"",
		"source directory or file")
}
