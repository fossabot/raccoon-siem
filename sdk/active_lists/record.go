package activeLists

import (
	"github.com/francoispqt/gojay"
)

type record struct {
	ExpiresAt int64
	Version   int64
	Fields    recordFields
}

func newRecord() record {
	return record{Fields: make(recordFields)}
}

func (r *record) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Int64KeyOmitEmpty("ExpiresAt", r.ExpiresAt)
	enc.Int64KeyOmitEmpty("Version", r.Version)
	enc.ObjectKeyOmitEmpty("Fields", r.Fields)
}

func (r *record) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "ExpiresAt":
		return dec.Int64(&r.ExpiresAt)
	case "Version":
		return dec.Int64(&r.Version)
	case "Fields":
		return dec.Object(r.Fields)
	default:
		return nil
	}
}

func (r *record) NKeys() int {
	return 0
}

func (r *record) IsNil() bool {
	return r == nil
}

type recordFields map[string]string

func (r recordFields) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range r {
		enc.StringKeyOmitEmpty(k, v)
	}
}

func (r recordFields) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	str := ""
	if err := dec.String(&str); err != nil {
		return err
	}
	r[k] = str
	return nil
}

func (r recordFields) IsNil() bool {
	return r == nil
}

func (r recordFields) NKeys() int {
	return 0
}
