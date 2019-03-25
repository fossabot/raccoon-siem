package activeLists

import "gopkg.in/vmihailenco/msgpack.v4"

type Record struct {
	ExpiresAt int64
	Version   int64 `json:"-"`
	Fields    map[string]interface{}
}

func (r *Record) EncodeMsgpack(enc *msgpack.Encoder) error {
	if err := enc.EncodeMulti(r.ExpiresAt, r.Version); err != nil {
		return err
	}

	if err := enc.EncodeMapLen(len(r.Fields)); err != nil {
		return err
	}

	for k, v := range r.Fields {
		if err := enc.EncodeString(k); err != nil {
			return err
		}

		if err := enc.EncodeMulti(v); err != nil {
			return err
		}
	}

	return nil
}

func (r *Record) DecodeMsgpack(dec *msgpack.Decoder) error {
	if err := dec.DecodeMulti(&r.ExpiresAt, &r.Version); err != nil {
		return err
	}

	l, err := dec.DecodeMapLen()
	if err != nil {
		return err
	}

	if l == 0 {
		r.Fields = nil
		return nil
	}

	r.Fields = make(map[string]interface{})
	for i := 0; i < l; i++ {
		k, err := dec.DecodeString()
		if err != nil {
			return err
		}

		v, err := dec.DecodeInterface()
		if err != nil {
			return err
		}

		r.Fields[k] = v
	}

	return nil
}
