package helpers

import "unsafe"

func BytesToString(slice []byte) string {
	return *(*string)(unsafe.Pointer(&slice))
}

func CopyBytes(data []byte) []byte {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	return dataCopy
}
