package sdk

import (
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net"
	"os"
)

func GetHostName() string {
	hn, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hn
}

func GetIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	panic("can not get IP address")
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func CopyBytes(data []byte) []byte {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	return dataCopy
}

func DebugError(err error) {
	if Debug {
		fmt.Println(err)
	}
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func FatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
