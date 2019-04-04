package filters

import "net"

const (
	OpEQ          = "="
	OpNEQ         = "!="
	OpGTorEQ      = ">="
	OpGT          = ">"
	OpLTorEQ      = "<="
	OpLT          = "<"
	OpInSubnet    = "inSubnet"
	OpNotInSubnet = "!inSubnet"
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
	case float64:
		return r.compareFloat(src.(float64), dst, op)
	case string:
		return r.compareString(src.(string), dst, op)
	default:
		return false
	}
}

func (r *comparator) compareInt(src int64, dst interface{}, op string) bool {
	dstVal, ok := dst.(int64)
	if !ok {
		return false
	}

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
	dstVal, ok := dst.(float64)
	if !ok {
		return false
	}

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
	dstVal, ok := dst.(string)
	if !ok {
		return false
	}

	switch op {
	case OpGT:
		return src > dstVal
	case OpGTorEQ:
		return src >= dstVal
	case OpLT:
		return src < dstVal
	case OpLTorEQ:
		return src <= dstVal
	case OpInSubnet:
		return r.inSubnet(src, dstVal)
	case OpNotInSubnet:
		return !r.inSubnet(src, dstVal)
	}

	return false
}

func (r *comparator) inSubnet(ip, cidr string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}

	return network.Contains(parsedIP)
}
