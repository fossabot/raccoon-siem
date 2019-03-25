package activeLists

import "gopkg.in/vmihailenco/msgpack.v4"

type changeLog struct {
	CID     string
	ALName  string
	Op      string
	Key     string
	Version int64
	Record  Record
}

func (r *changeLog) EncodeMsgpack(enc *msgpack.Encoder) error {
	if err := enc.EncodeMulti(r.CID, r.ALName, r.Op, r.Key, r.Version); err != nil {
		return err
	}
	return r.Record.EncodeMsgpack(enc)
}

func (r *changeLog) DecodeMsgpack(dec *msgpack.Decoder) error {
	if err := dec.DecodeMulti(&r.CID, &r.ALName, &r.Op, &r.Key, &r.Version); err != nil {
		return err
	}
	return r.Record.DecodeMsgpack(dec)
}
