package get

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/resources/helpers"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:       "get <resource kind> <resource ID>",
	Short:     "get particular resource",
	Args:      cobra.ExactArgs(2),
	ValidArgs: helpers.ValidResourceKinds,
	RunE:      helpers.CheckArgsAndExec(run),
}

var flags cmdFlags

func init() {
	// Core URL
	Cmd.Flags().StringVar(&flags.CoreURL, "core", "http://localhost:7220", "raccoon core URL")
}

func run(args []string) error {
	resourceKind := args[0]
	resourceID := args[1]

	url := fmt.Sprintf("%s/%s/%s", flags.CoreURL, resourceKind, resourceID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	data, err := helpers.SendCoreRequest(req)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}
