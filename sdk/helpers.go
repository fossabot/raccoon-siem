package sdk

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net"
	"os"
	"sort"
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

// Sort eventSpecs by EndTime DESC
func sortEventsByEndTime(events []*Event, asc bool) {
	sort.SliceStable(events, func(i, j int) bool {
		if asc {
			return events[i].EndTime.UnixNano() < events[j].EndTime.UnixNano()
		} else {
			return events[i].EndTime.UnixNano() > events[j].EndTime.UnixNano()
		}
	})
}

func GetUUID() string {
	return uuid.NewV4().String()
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

func PrintErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
