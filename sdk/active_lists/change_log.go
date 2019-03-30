package activeLists

import "github.com/francoispqt/gojay"

type changeLog struct {
	CID     string
	ALName  string
	Op      string
	Key     string
	Version int64
	Record  record
}

func newChangeLog() changeLog {
	return changeLog{Record: newRecord()}
}

func (r *changeLog) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty("CID", r.CID)
	enc.StringKeyOmitEmpty("ALName", r.ALName)
	enc.StringKeyOmitEmpty("Op", r.Op)
	enc.StringKeyOmitEmpty("Key", r.Key)
	enc.Int64KeyOmitEmpty("Version", r.Version)
	enc.ObjectKeyOmitEmpty("Record", &r.Record)
}

func (r *changeLog) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "CID":
		return dec.String(&r.CID)
	case "ALName":
		return dec.String(&r.ALName)
	case "Op":
		return dec.String(&r.Op)
	case "Key":
		return dec.String(&r.Key)
	case "Version":
		return dec.Int64(&r.Version)
	case "Record":
		return dec.Object(&r.Record)
	default:
		return nil
	}
}

func (r *changeLog) NKeys() int {
	return 0
}

func (r *changeLog) IsNil() bool {
	return r == nil
}
