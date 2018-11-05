package resources

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var getCmd = &cobra.Command{
	Use:       "get <resource kind> <resource ID>",
	Short:     "get particular resource configuration",
	Args:      cobra.ExactArgs(2),
	ValidArgs: validResourceKinds,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cobra.OnlyValidArgs(cmd, args[:1]); err != nil {
			return err
		}
		return get(args[0], args[1])
	},
}

func get(resourceKind string, resourceID string) error {
	url := fmt.Sprintf("%s/%s/%s", coreURL, resourceKind, resourceID)

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
