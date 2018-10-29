package sdk

import "time"

const (
	opEQ = iota
	opNEQ
	opGTorEQ
	opGT
	opLTorEQ
	opLT
)

const (
	opEQString     = "=="
	opNEQString    = "!="
	opGTorEQString = ">="
	opGTString     = ">"
	opLTorEQString = "<="
	opLTString     = "<"
)

type comparator struct{}

func (ef *comparator) compareValues(src interface{}, srcType byte, dst interface{}, op byte) bool {
	if op == opEQ {
		return src == dst
	}

	if op == opNEQ {
		return src != dst
	}

	switch srcType {
	case fieldTypeInt:
		return ef.compareInt(src.(int64), dst, op)
	case fieldTypeFloat:
		return ef.compareFloat(src.(float64), dst, op)
	case fieldTypeTime:
		return ef.compareTime(src.(time.Time), dst, op)
	case fieldTypeDuration:
		return ef.compareDuration(src.(time.Duration), dst, op)
	case fieldTypeString:
		return ef.compareString(src.(string), dst, op)
	}

	return false
}

func (ef *comparator) compareInt(src int64, dst interface{}, op byte) bool {
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

func (ef *comparator) compareFloat(src float64, dst interface{}, op byte) bool {
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

func (ef *comparator) compareString(src string, dst interface{}, op byte) bool {
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

func (ef *comparator) compareTime(src time.Time, dst interface{}, op byte) bool {
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

func (ef *comparator) compareDuration(src time.Duration, dst interface{}, op byte) bool {
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
