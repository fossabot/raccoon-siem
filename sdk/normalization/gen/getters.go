package main

var gettersTemplate = `
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
	default: return false
	}
}
`
