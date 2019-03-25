package resources

import (
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"os"
	"reflect"
	"text/template"
)

var generateCmd = &cobra.Command{
	Use:       "generate <resource kind>",
	Short:     "generate resource configuration",
	Args:      cobra.ExactArgs(1),
	ValidArgs: validGenerateSubjects,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cobra.OnlyValidArgs(cmd, args); err != nil {
			return err
		}
		return generate(args[0])
	},
}

func generate(resourceKind string) error {
	switch resourceKind {
	case "elasticsearch":
		return generateElasticsearchTemplate()
	}
	return nil
}

func generateElasticsearchTemplate() error {
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
