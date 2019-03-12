package sdk

import (
	"fmt"
	"github.com/satori/go.uuid"
	"log"
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

func CopyBytes(data []byte) []byte {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	return dataCopy
}

func PrintConfiguration(resources ...interface{}) {
	for _, r := range resources {
		switch r.(type) {
		case []IConnector:
			fmt.Printf("Connectors:\n")
			for i, v := range r.([]IConnector) {
				fmt.Printf("\t%d.%v\n", i+1, v.ID())
			}
		case []IParser:
			fmt.Printf("Parsers:\n")
			for i, v := range r.([]IParser) {
				fmt.Printf("\t%d.%v\n", i+1, v.ID())
			}
		case []IFilter:
			fmt.Printf("Filters:\n")
			for i, v := range r.([]IFilter) {
				fmt.Printf("\t%d.%v\n", i+1, v.ID())
			}
		case []IAggregationRule:
			fmt.Printf("Aggregation rules:\n")
			for i, v := range r.([]IAggregationRule) {
				fmt.Printf("\t%d.%v\n", i+1, v.ID())
			}
		case []IDestination:
			fmt.Printf("Destinations:\n")
			for i, v := range r.([]IDestination) {
				fmt.Printf("\t%d.%v\n", i+1, v.ID())
			}
		case []CorrelationRule:
			fmt.Printf("Correlation rules:\n")
			for i, v := range r.([]CorrelationRule) {
				fmt.Printf("\t%d.%v\n", i+1, v.ID())
			}
		case []ActiveListSettings:
			fmt.Printf("Active lists:\n")
			for i, v := range r.([]ActiveListSettings) {
				fmt.Printf("\t%d.%v\n", i+1, v.ID())
			}
		}
	}
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
