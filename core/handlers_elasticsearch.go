package core

import (
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"html/template"
	"reflect"
)

const (
	storageTypeTag = "storage_type"
)

var storageMappingTemplate = template.Must(template.New("raccoon-events").Parse(
	`{
  "index_patterns": ["raccoon-events-*"],
  "settings": { "number_of_shards": 1 },
  "mappings": {
    "_doc": {
      "dynamic": "false",
      "properties": {
       	{{range .Fields}}
		  "{{.Name}}": { "type": "{{.Type}}" }{{if not .Last }},{{end}}
		{{end}}
      }
    }
  }
}`))

type storageTemplateParams struct {
	Fields []storageField
}

type storageField struct {
	Name string
	Type string
	Last bool
}

func GenerateStorageMapping(ctx *gin.Context) {
	e := new(sdk.Event)
	rv := reflect.ValueOf(e).Elem()
	rt := rv.Type()

	templateParams := &storageTemplateParams{}
	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		storageType := field.Tag.Get(storageTypeTag)
		if storageType != "" {
			templateParams.Fields = append(templateParams.Fields, storageField{Name: field.Name, Type: storageType})
		}
	}

	// Mark last field to skip unnecessary punctuation
	templateParams.Fields[len(templateParams.Fields)-1].Last = true

	if err := storageMappingTemplate.Execute(ctx.Writer, templateParams); err != nil {
		reply(ctx, err)
	}
}
