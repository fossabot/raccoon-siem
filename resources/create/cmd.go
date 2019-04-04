package create

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/resources/helpers"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "create or update resource from files",
	Args:  cobra.ExactArgs(0),
	RunE:  run,
}

var flags cmdFlags

func init() {
	// Source file path
	Cmd.Flags().StringVar(&flags.SourcePath, "from", "", "source file or directory path")
	// Core URL
	Cmd.Flags().StringVar(&flags.CoreURL, "core", "http://localhost:7220", "raccoon core URL")
	// Required flags
	_ = Cmd.MarkFlagRequired("from")
}

func run(_ *cobra.Command, _ []string) error {
	resources, err := helpers.ReadResourcesFromInputFile(flags.SourcePath)
	if err != nil {
		return err
	}

	for _, resource := range resources {
		url := fmt.Sprintf("%s/%s", flags.CoreURL, resource.Kind)
		data := bytes.NewBuffer(resource.Data)

		req, err := http.NewRequest(http.MethodPut, url, data)
		if err != nil {
			return err
		}

		if _, err := helpers.SendCoreRequest(req); err != nil {
			return err
		}
	}

	fmt.Println("OK")
	return nil
}
