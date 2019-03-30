package main

var settersTemplate = `
func (r *Event) SetAnyField(field string, value string) bool {
	if len(value) == 0 {
		return false
	}

	switch field {
	{{- range .}}
	{{- if .Set}}
	{{- if eq .Kind "string"}}
	case "{{.Name}}": r.{{.Name}} = strings.TrimSpace(value)
	{{- end}}
	{{- if and (eq .Kind "int64") (not .Time)}}
	case "{{.Name}}": r.{{.Name}} = StringToInt(value)
	{{- end}}
	{{- if .Time}}
	case "{{.Name}}": r.{{.Name}} = StringToTime(value)
	{{- end}}
	{{- if eq .Kind "float64"}}
	case "{{.Name}}": r.{{.Name}} = StringToFloat(value)
	{{- end}}
	{{- if eq .Kind "bool"}}
	case "{{.Name}}": r.{{.Name}} = StringToBool(value)
	{{- end}}
	{{- end}}
	{{- end}}
	default: return false
	}

	return true
}

func (r *Event) SetIntField(field string, value int64) {
	switch field {
	{{- range .}}
	{{- if .Set}}
	{{- if eq .Kind "int64"}}
	case "{{.Name}}": r.{{.Name}} = value
	{{- end}}
	{{- end}}
	{{- end}}
	}
}

func (r *Event) SetFloatField(field string, value float64) {
	switch field {
	{{- range .}}
	{{- if .Set}}
	{{- if eq .Kind "float64"}}
	case "{{.Name}}": r.{{.Name}} = value
	{{- end}}
	{{- end}}
	{{- end}}
	}
}

func (r *Event) SetBoolField(field string, value bool) {
	switch field {
	{{- range .}}
	{{- if .Set}}
	{{- if eq .Kind "bool"}}
	case "{{.Name}}": r.{{.Name}} = value
	{{- end}}
	{{- end}}
	{{- end}}
	}
}
`
