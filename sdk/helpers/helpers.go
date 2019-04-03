package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func CopyBytes(data []byte) []byte {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	return dataCopy
}

func BytesToString(input []byte) string {
	return *(*string)(unsafe.Pointer(&input))
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

func ReadConfigFromCore(baseURL string, component string, id string, dst interface{}) error {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", baseURL, component, id))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("core replied with %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, dst)
}

func ReadConfigFromFile(path string, dstPointer interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dstPointer)
}

func NowUnixMillis() int64 {
	return time.Now().UnixNano() / 1000 / 1000
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

func StringToSingleByte(s string) (byte, error) {
	bs := []byte(s)
	if len(bs) != 1 {
		return 0, fmt.Errorf("expected single byte ASCII character, got: %s", s)
	}
	return bs[0], nil
}

func MakeKey(keyFields []string, event *normalization.Event) string {
	key := strings.Builder{}
	for _, field := range keyFields {
		key.WriteString(normalization.ToString(event.GetAnyField(field)))
	}
	return key.String()
}

func AreEventFieldTypesEqual(lf, rf string) bool {
	event := new(normalization.Event)
	lv := event.GetAnyField(lf)
	rv := event.GetAnyField(rf)
	return reflect.TypeOf(lv).Kind() == reflect.TypeOf(rv).Kind()
}

func EventFieldHasGetter(field string) bool {
	event := normalization.Event{}
	rt := reflect.TypeOf(event)
	_, ok := rt.FieldByName(field)
	return ok
}

func EventFieldHasSetter(field string) bool {
	event := normalization.Event{}
	rt := reflect.TypeOf(event)
	f, ok := rt.FieldByName(field)
	if !ok {
		return false
	}
	return f.Tag.Get("set") != ""
}
