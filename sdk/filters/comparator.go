package filters

const (
	OpEQ     = "="
	OpNEQ    = "!="
	OpGTorEQ = ">="
	OpGT     = ">"
	OpLTorEQ = "<="
	OpLT     = "<"
)

type comparator struct{}

func (r *comparator) compareValues(src interface{}, dst interface{}, op string) bool {
	if op == OpEQ {
		return src == dst
	}

	if op == OpNEQ {
		return src != dst
	}

	switch src.(type) {
	case int64:
		return r.compareInt(src.(int64), dst, op)
	case int:
		return r.compareInt(int64(src.(int)), dst, op)
	case float64:
		return r.compareFloat(src.(float64), dst, op)
	case string:
		return r.compareString(src.(string), dst, op)
	default:
		return false
	}
}

func (r *comparator) compareInt(src int64, dst interface{}, op string) bool {
	dstVal := dst.(int64)
	switch op {
	case OpGT:
		return src > dstVal
	case OpGTorEQ:
		return src >= dstVal
	case OpLT:
		return src < dstVal
	case OpLTorEQ:
		return src <= dstVal
	}
	return false
}

func (r *comparator) compareFloat(src float64, dst interface{}, op string) bool {
	dstVal := dst.(float64)
	switch op {
	case OpGT:
		return src > dstVal
	case OpGTorEQ:
		return src >= dstVal
	case OpLT:
		return src < dstVal
	case OpLTorEQ:
		return src <= dstVal
	}
	return false
}

func (r *comparator) compareString(src string, dst interface{}, op string) bool {
	dstVal := dst.(string)
	switch op {
	case OpGT:
		return src > dstVal
	case OpGTorEQ:
		return src >= dstVal
	case OpLT:
		return src < dstVal
	case OpLTorEQ:
		return src <= dstVal
	}
	return false
}
