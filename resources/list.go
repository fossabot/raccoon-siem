package resources

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var listCmd = &cobra.Command{
	Use:       "list <resource kind>",
	Short:     "list resources of kind",
	Args:      cobra.ExactArgs(1),
	ValidArgs: validResourceKinds,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cobra.OnlyValidArgs(cmd, args); err != nil {
			return err
		}
		return list(args[0])
	},
}

func list(resourceKind string) error {
	url := fmt.Sprintf("%s/%s", coreURL, resourceKind)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	data, err := sendRequest(req)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}
