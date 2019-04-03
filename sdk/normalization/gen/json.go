package main

var encodeJSONTemplate = `
func (r *Event) MarshalJSONObject(enc *gojay.Encoder) {
	{{- range .}}
	{{- if eq .Kind "strSlice" }}
	enc.ArrayKeyOmitEmpty("{{.Name}}", r.{{.Name}})
	{{- end}}
	{{- if eq .Kind "string" }}
	enc.StringKeyOmitEmpty("{{.Name}}", r.{{.Name}})
	{{- end}}
	{{- if eq .Kind "int64" }}
	enc.Int64KeyOmitEmpty("{{.Name}}", r.{{.Name}})
	{{- end}}
	{{- if eq .Kind "int" }}
	enc.IntKeyOmitEmpty("{{.Name}}", r.{{.Name}})
	{{- end}}
	{{- if eq .Kind "float64" }}
	enc.Float64KeyOmitEmpty("{{.Name}}", r.{{.Name}})
	{{- end}}
	{{- if eq .Kind "bool" }}
	enc.BoolKeyOmitEmpty("{{.Name}}", r.{{.Name}})
	{{- end}}
	{{- end}}
}
`

var decodeJSONTemplate = `
func (r *Event) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	{{- range .}}
	case "{{.Name}}":
		{{- if eq .Kind "strSlice" }}
		return dec.Array(&r.{{.Name}})
		{{- end}}
		{{- if eq .Kind "string" }}
		return dec.String(&r.{{.Name}})
		{{- end}}
		{{- if eq .Kind "int64" }}
		return dec.Int64(&r.{{.Name}})
		{{- end}}
		{{- if eq .Kind "int" }}
		return dec.Int(&r.{{.Name}})
		{{- end}}
		{{- if eq .Kind "float64" }}
		return dec.Float64(&r.{{.Name}})
		{{- end}}
		{{- if eq .Kind "bool" }}
		return dec.Bool(&r.{{.Name}})
		{{- end}}
	{{- end}}
	default:
		return nil
	}
}
`
