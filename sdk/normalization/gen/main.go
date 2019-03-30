package main

import (
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"os"
	"reflect"
	"strings"
	"text/template"
)

type eventFieldMeta struct {
	Name string
	Kind string
	Set  bool
	Last bool
}

func main() {
	getters()
}

func scanEvent() (result []eventFieldMeta) {
	e := normalization.Event{}
	rt := reflect.TypeOf(e)

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if strings.Index(f.Name, "_") == 0 {
			continue
		}

		result = append(result, eventFieldMeta{
			Name: f.Name,
			Kind: f.Type.Kind().String(),
			Set:  f.Tag.Get("set") != "",
		})
	}

	if len(result) > 0 {
		result[len(result)-1].Last = true
	}

	return
}

var gettersTemplate = `
package normalization

func (r *Event) GetAnyField(field string) interface{} {
	switch field {
	{{- range .}}
	case "{{.Name}}": return r.{{.Name}}
	{{- end}}
	default: return nil
	}
}

func (r *Event) GetIntField(field string) int64 {
	switch field {
	{{- range .}}
	{{- if eq .Kind "int64"}}
	case "{{.Name}}": return r.{{.Name}}
	{{- end}}
	{{- end}}
	default: return 0
	}
}

func (r *Event) GetFloatField(field string) float64 {
	switch field {
	{{- range .}}
	{{- if eq .Kind "float64"}}
	case "{{.Name}}": return r.{{.Name}}
	{{- end}}
	{{- end}}
	default: return 0
	}
}

func (r *Event) GetBoolField(field string) bool {
	switch field {
	{{- range .}}
	{{- if eq .Kind "bool"}}
	case "{{.Name}}": return r.{{.Name}}
	{{- end}}
	{{- end}}
	default: return 0
	}
}
`

func getters() {
	tpl, err := template.New("gettersTemplate").Parse(gettersTemplate)
	if err != nil {
		panic(err)
	}

	if err := tpl.Execute(os.Stdout, scanEvent()); err != nil {
		panic(err)
	}
}
