package resources

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var deleteCmd = &cobra.Command{
	Use:       "delete <resource kind> <resource ID>",
	Short:     "delete particular resource",
	Args:      cobra.ExactArgs(2),
	ValidArgs: validResourceKinds,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cobra.OnlyValidArgs(cmd, args[:1]); err != nil {
			return err
		}
		return delete(args[0], args[1])
	},
}

func delete(resourceKind string, resourceID string) error {
	url := fmt.Sprintf("%s/%s/%s", coreURL, resourceKind, resourceID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = sendRequest(req)
	if err != nil {
		return err
	}

	fmt.Println("OK")
	return nil
}
