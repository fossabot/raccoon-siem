package helpers

import "unsafe"

func BytesToString(slice []byte) string {
	return *(*string)(unsafe.Pointer(&slice))
}
