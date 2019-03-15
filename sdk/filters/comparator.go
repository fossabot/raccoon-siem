package filters

import (
	"time"
)

const (
	opEQ     = "=="
	opNEQ    = "!="
	opGTorEQ = ">="
	opGT     = ">"
	opLTorEQ = "<="
	opLT     = "<"
)

type comparator struct{}

func (r *comparator) compareValues(src interface{}, srcType byte, dst interface{}, op string) bool {
	if op == opEQ {
		return src == dst
	}

	if op == opNEQ {
		return src != dst
	}

	switch src.(type) {
	case int64:
		return r.compareInt(src.(int64), dst, op)
	case float64:
		return r.compareFloat(src.(float64), dst, op)
	case time.Time:
		return r.compareTime(src.(time.Time), dst, op)
	case time.Duration:
		return r.compareDuration(src.(time.Duration), dst, op)
	case string:
		return r.compareString(src.(string), dst, op)
	default:
		return false
	}
}

func (r *comparator) compareInt(src int64, dst interface{}, op string) bool {
	dstVal := dst.(int64)
	switch op {
	case opGT:
		return src > dstVal
	case opGTorEQ:
		return src >= dstVal
	case opLT:
		return src < dstVal
	case opLTorEQ:
		return src <= dstVal
	}
	return false
}

func (r *comparator) compareFloat(src float64, dst interface{}, op string) bool {
	dstVal := dst.(float64)
	switch op {
	case opGT:
		return src > dstVal
	case opGTorEQ:
		return src >= dstVal
	case opLT:
		return src < dstVal
	case opLTorEQ:
		return src <= dstVal
	}
	return false
}

func (r *comparator) compareString(src string, dst interface{}, op string) bool {
	dstVal := dst.(string)
	switch op {
	case opGT:
		return src > dstVal
	case opGTorEQ:
		return src >= dstVal
	case opLT:
		return src < dstVal
	case opLTorEQ:
		return src <= dstVal
	}
	return false
}

func (r *comparator) compareTime(src time.Time, dst interface{}, op string) bool {
	dstVal := dst.(time.Time)
	switch op {
	case opGT:
		return src.UnixNano() > dstVal.UnixNano()
	case opGTorEQ:
		return src.UnixNano() >= dstVal.UnixNano()
	case opLT:
		return src.UnixNano() < dstVal.UnixNano()
	case opLTorEQ:
		return src.UnixNano() <= dstVal.UnixNano()
	}
	return false
}

func (r *comparator) compareDuration(src time.Duration, dst interface{}, op string) bool {
	dstVal := dst.(time.Duration)
	switch op {
	case opGT:
		return src > dstVal
	case opGTorEQ:
		return src >= dstVal
	case opLT:
		return src < dstVal
	case opLTorEQ:
		return src <= dstVal
	}
	return false
}
