package elastic

import (
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"os"
	"reflect"
	"text/template"
)

var Cmd = &cobra.Command{
	Use:   "elasticsearch",
	Short: "generate elasticsearch index template",
	Args:  cobra.ExactArgs(0),
	RunE:  run,
}

func run(_ *cobra.Command, _ []string) error {
	tpl, err := template.New("elasticsearch").Parse(elasticsearchTemplate)
	if err != nil {
		return err
	}

	tplData := make([]*elasticsearchMappingProperty, 0)
	e := normalization.Event{}
	rt := reflect.TypeOf(e)

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		storageType := f.Tag.Get("storage_type")
		if storageType != "" {
			tplData = append(tplData, &elasticsearchMappingProperty{
				Name: f.Name,
				Type: storageType,
			})
		}
	}

	if len(tplData) > 0 {
		tplData[len(tplData)-1].Last = true
	}

	return tpl.Execute(os.Stdout, tplData)
}
