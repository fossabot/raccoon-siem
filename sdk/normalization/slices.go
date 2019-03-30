package normalization

import "github.com/francoispqt/gojay"

type strSlice []string

func (r strSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range r {
		enc.String(r[i])
	}
}

func (r *strSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	str := ""
	if err := dec.String(&str); err != nil {
		return err
	}
	*r = append(*r, str)
	return nil
}

func (r strSlice) IsNil() bool {
	return len(r) == 0
}
