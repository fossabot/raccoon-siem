package normalization

import (
	"encoding/json"
	"github.com/francoispqt/gojay"
	"gopkg.in/vmihailenco/msgpack.v4"
	"strconv"
	"strings"
	"time"
)

const (
	FieldTypeString = iota
	FieldTypeInt
	FieldTypeFloat
	FieldTypeBool
	FieldTypeTime
	FieldTypeDuration
)

const (
	SeverityInfo = iota
	SeverityWarn
	SeverityError
	SeverityCritical
)

const (
	DefaultEventFieldsHash = "_"
)

type Event struct {
	ID                         string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	Tag                        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	Timestamp                  int64    `json:",omitempty" msgpack:",omitempty" storage_type:"date"`
	BaseEventCount             int      `json:",omitempty" msgpack:",omitempty" storage_type:"integer"`
	AggregatedEventCount       int      `json:",omitempty" msgpack:",omitempty" storage_type:"integer"`
	AggregationRuleName        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	CollectorIPAddress         string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip" `
	CollectorDNSName           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	CorrelationRuleName        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	CorrelatorIPAddress        string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip"`
	CorrelatorDNSName          string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	SourceID                   string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	BaseEventIDs               strSlice `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	Incident                   bool     `json:",omitempty" msgpack:",omitempty" storage_type:"boolean" set:"y"`
	Correlated                 bool     `json:",omitempty" msgpack:",omitempty" storage_type:"boolean" set:"y"`
	Score                      int64    `json:",omitempty" msgpack:",omitempty" storage_type:"integer" set:"y"`
	Severity                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"integer" set:"y"`
	Customer                   string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	Code                       string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	Message                    string   `json:",omitempty" msgpack:",omitempty" storage_type:"text" set:"y"`
	Details                    string   `json:",omitempty" msgpack:",omitempty" storage_type:"text" set:"y"`
	Trace                      string   `json:",omitempty" msgpack:",omitempty" storage_type:"text" set:"y"`
	OriginEventID              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	OriginTimestamp            int64    `json:",omitempty" msgpack:",omitempty" storage_type:"date" set:"y"`
	OriginEnvironment          string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	OriginSeverity             string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	OriginServiceName          string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	OriginServiceVersion       string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	OriginProcessName          string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	OriginFileName             string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	OriginDNSName              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	OriginDomain               string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	OriginIPAddress            string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip" set:"y"`
	RequestID                  string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestApplicationProtocol string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestTransportProtocol   string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestURL                 string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestReferrer            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestMethod              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestUserAgent           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestStatus              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestTook                int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	RequestBytesIn             int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	RequestBytesOut            int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	RequestResults             int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	RequestUser                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestUnit                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	RequestOrganization        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	SourceIPAddress            string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip" set:"y"`
	SourceMACAddress           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	SourceDomain               string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	SourceDNSName              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	SourcePort                 string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	DestinationIPAddress       string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip" set:"y"`
	DestinationMACAddress      string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	DestinationDomain          string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	DestinationDNSName         string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	DestinationPort            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString1                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString1Label           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString2                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString2Label           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString3                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString3Label           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString4                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString4Label           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString5                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString5Label           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString6                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString6Label           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString7                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString7Label           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString8                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserString8Label           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserInt1                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	UserInt1Label              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserInt2                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	UserInt2Label              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserInt3                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	UserInt3Label              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserInt4                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	UserInt4Label              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserInt5                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	UserInt5Label              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserInt6                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	UserInt6Label              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserInt7                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	UserInt7Label              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserInt8                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long" set:"y"`
	UserInt8Label              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserFloat1                 float64  `json:",omitempty" msgpack:",omitempty" storage_type:"double" set:"y"`
	UserFloat1Label            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserFloat2                 float64  `json:",omitempty" msgpack:",omitempty" storage_type:"double" set:"y"`
	UserFloat2Label            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserFloat3                 float64  `json:",omitempty" msgpack:",omitempty" storage_type:"double" set:"y"`
	UserFloat3Label            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserFloat4                 float64  `json:",omitempty" msgpack:",omitempty" storage_type:"double" set:"y"`
	UserFloat4Label            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserFloat5                 float64  `json:",omitempty" msgpack:",omitempty" storage_type:"double" set:"y"`
	UserFloat5Label            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserFloat6                 float64  `json:",omitempty" msgpack:",omitempty" storage_type:"double" set:"y"`
	UserFloat6Label            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserFloat7                 float64  `json:",omitempty" msgpack:",omitempty" storage_type:"double" set:"y"`
	UserFloat7Label            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserFloat8                 float64  `json:",omitempty" msgpack:",omitempty" storage_type:"double" set:"y"`
	UserFloat8Label            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserTimestamp1             int64    `json:",omitempty" msgpack:",omitempty" storage_type:"date" set:"y"`
	UserTimestamp1Label        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserTimestamp2             int64    `json:",omitempty" msgpack:",omitempty" storage_type:"date" set:"y"`
	UserTimestamp2Label        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserTimestamp3             int64    `json:",omitempty" msgpack:",omitempty" storage_type:"date" set:"y"`
	UserTimestamp3Label        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
	UserTimestamp4             int64    `json:",omitempty" msgpack:",omitempty" storage_type:"date" set:"y"`
	UserTimestamp4Label        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword" set:"y"`
}

func (r *Event) SetID(id string) {
	r.ID = id
}

func (r *Event) SetAnyFieldBytes(field string, value []byte) {
	r.SetAnyField(field, BytesToString(value))
}

func (r *Event) HashFields(fieldNames []string) string {
	if len(fieldNames) == 0 {
		return DefaultEventFieldsHash
	}

	builder := strings.Builder{}
	for _, field := range fieldNames {
		val := r.GetAnyField(field)
		switch val.(type) {
		case string:
			builder.WriteString(val.(string))
		case int64:
			builder.WriteString(strconv.FormatInt(val.(int64), 10))
		case float64:
			builder.WriteString(strconv.FormatFloat(val.(float64), 'f', -1, 64))
		case time.Time:
			builder.WriteString(val.(time.Time).String())
		case time.Duration:
			builder.WriteString(val.(time.Duration).String())
		case bool:
			builder.WriteString(strconv.FormatBool(val.(bool)))
		}
	}

	return builder.String()
}

func (r *Event) Clone() Event {
	return *r
}

func (r *Event) ToJSON() ([]byte, error) {
	return gojay.Marshal(r)
}

func (r *Event) FromJSON(input []byte) error {
	return gojay.Unmarshal(input, r)
}

func (r *Event) ToMsgPack() ([]byte, error) {
	return msgpack.Marshal(r)
}

func (r *Event) FromMsgPack(input []byte) error {
	return msgpack.Unmarshal(input, r)
}

func (r *Event) NKeys() int {
	return 0
}

func (r *Event) IsNil() bool {
	return r == nil
}

func (r *Event) String() string {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return ""
	}
	return string(data)
}
