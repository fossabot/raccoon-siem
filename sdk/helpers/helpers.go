package helpers

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func CopyBytes(data []byte) []byte {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	return dataCopy
}

func SumEvents(dst *normalization.Event, src *normalization.Event, fields []string) {
	for _, field := range fields {
		dstValue := dst.GetAnyField(field)
		switch dstValue.(type) {
		case int64:
			srcValue := src.GetIntField(field)
			dst.SetIntField(field, dstValue.(int64)+srcValue)
		case float64:
			newValue := dstValue.(float64) + src.GetFloatField(field)
			dst.SetFloatField(field, newValue)
		case string:
			srcValue, ok := src.GetAnyField(field).(string)
			if ok {
				sb := strings.Builder{}
				if srcValue != "" {
					sb.WriteByte(',')
				}
				sb.WriteString(srcValue)
				dst.SetAnyField(field, sb.String())
			}
		default:
			continue
		}
	}
}

func CopyFields(dst *normalization.Event, src *normalization.Event, fields []string) {
	for _, field := range fields {
		srcValue := src.GetAnyField(field)
		switch srcValue.(type) {
		case string:
			dst.SetAnyField(field, srcValue.(string))
		case int64:
			dst.SetIntField(field, srcValue.(int64))
		case float64:
			dst.SetFloatField(field, srcValue.(float64))
		}
	}
}

func CoreQuery(url string, dst interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(body, dst); err != nil {
		if len(body) > 0 {
			return fmt.Errorf("%s", string(body))
		}
		return err
	}

	return nil
}

func NowUnixMillis() int64 {
	return time.Now().UnixNano() / 1000
}

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
