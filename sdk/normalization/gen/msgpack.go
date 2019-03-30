package main

var encodeMsgpackTemplate = `
func (r *Event) EncodeMsgpack(enc *msgpack.Encoder) error {
	{{- range .}}
	{{- if eq .Kind "strSlice" }}
	if err := enc.EncodeArrayLen(len(r.{{.Name}})); err != nil {
		return err
	}

	for i := range r.{{.Name}} {
		if err := enc.EncodeString(r.{{.Name}}[i]); err != nil {
			return err
		}	
	}
	{{end}}

	{{- if eq .Kind "string" }}
	if err := enc.EncodeString(r.{{.Name}}); err != nil {
		return err
	}
	{{end}}

	{{- if eq .Kind "int64" }}
	if err := enc.EncodeInt64(r.{{.Name}}); err != nil {
		return err
	}
	{{end}}

	{{- if eq .Kind "int" }}
	if err := enc.EncodeInt64(int64(r.{{.Name}})); err != nil {
		return err
	}
	{{end}}

	{{- if eq .Kind "float64" }}
	if err := enc.EncodeFloat64(r.{{.Name}}); err != nil {
		return err
	}
	{{end}}

	{{- if eq .Kind "bool" }}
	if err := enc.EncodeBool(r.{{.Name}}); err != nil {
		return err
	}
	{{end}}

	{{- end}}
	return nil
}
`

var decodeMsgpackTemplate = `
func (r *Event) DecodeMsgpack(dec *msgpack.Decoder) (err error) {
	{{- range .}}
	{{- if eq .Kind "strSlice" }}
	l, err := dec.DecodeArrayLen()
	if err != nil {
		return err
	}

	r.{{.Name}} = make(strSlice, l)
	for i := 0; i < l; i++ {
		r.{{.Name}}[i], err = dec.DecodeString()
		if err != nil {
			return err
		}
	}
	{{end}}

	{{- if eq .Kind "string" }}
	if r.{{.Name}}, err = dec.DecodeString(); err != nil {
		return err
	}
	{{end}}

	{{- if eq .Kind "int64" }}
	if r.{{.Name}}, err = dec.DecodeInt64(); err != nil {
		return err
	}
	{{end}}

	{{- if eq .Kind "int" }}
	{{.Name}}, err := dec.DecodeInt64()
	if err != nil {
		return err
	}
	r.{{.Name}} = int({{.Name}})
	{{end}}

	{{- if eq .Kind "float64" }}
	if r.{{.Name}}, err = dec.DecodeFloat64(); err != nil {
		return err
	}
	{{end}}

	{{- if eq .Kind "bool" }}
	if r.{{.Name}}, err = dec.DecodeBool(); err != nil {
		return err
	}
	{{end}}

	{{- end}}
	return nil
}
`
