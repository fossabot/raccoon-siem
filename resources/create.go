package resources

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"net/http"
)

var createCmd = &cobra.Command{
	Use:       "create",
	Short:     "create or update resource from files",
	Args:      cobra.ExactArgs(0),
	ValidArgs: validResourceKinds,
	RunE: func(cmd *cobra.Command, args []string) error {
		if sourcePath == "" {
			return errors.New("source file or directory must be specified")
		}
		return create()
	},
}

func create() error {
	resources, err := readResourcesFromInputFile(sourcePath)
	if err != nil {
		return err
	}

	for _, resource := range resources {
		url := fmt.Sprintf("%s/%s", coreURL, resource.kind)
		data := bytes.NewBuffer(resource.data)

		req, err := http.NewRequest(http.MethodPut, url, data)
		if err != nil {
			return err
		}

		if _, err := sendRequest(req); err != nil {
			return err
		}
	}

	fmt.Println("OK")
	return nil
}
