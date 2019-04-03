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
	ID                         string   `storage_type:"keyword"`
	Tag                        string   `storage_type:"keyword"`
	Timestamp                  int64    `storage_type:"date"`
	BaseEventCount             int      `storage_type:"integer"`
	AggregatedEventCount       int      `storage_type:"integer"`
	AggregationRuleName        string   `storage_type:"keyword"`
	CollectorIPAddress         string   `storage_type:"ip" `
	CollectorDNSName           string   `storage_type:"keyword"`
	CorrelationRuleName        string   `storage_type:"keyword"`
	CorrelatorIPAddress        string   `storage_type:"ip"`
	CorrelatorDNSName          string   `storage_type:"keyword"`
	SourceID                   string   `storage_type:"keyword"`
	BaseEventIDs               strSlice `storage_type:"keyword"`
	Incident                   bool     `storage_type:"boolean" set:"y"`
	Correlated                 bool     `storage_type:"boolean" set:"y"`
	Score                      int64    `storage_type:"integer" set:"y"`
	Severity                   int64    `storage_type:"integer" set:"y"`
	Customer                   string   `storage_type:"keyword" set:"y"`
	Code                       string   `storage_type:"keyword" set:"y"`
	Message                    string   `storage_type:"text" set:"y"`
	Details                    string   `storage_type:"text" set:"y"`
	Trace                      string   `storage_type:"text" set:"y"`
	OriginEventID              string   `storage_type:"keyword" set:"y"`
	OriginTimestamp            int64    `storage_type:"date" set:"y"`
	OriginEnvironment          string   `storage_type:"keyword" set:"y"`
	OriginSeverity             string   `storage_type:"keyword" set:"y"`
	OriginServiceName          string   `storage_type:"keyword" set:"y"`
	OriginServiceVersion       string   `storage_type:"keyword" set:"y"`
	OriginProcessName          string   `storage_type:"keyword" set:"y"`
	OriginFileName             string   `storage_type:"keyword" set:"y"`
	OriginDNSName              string   `storage_type:"keyword" set:"y"`
	OriginDomain               string   `storage_type:"keyword" set:"y"`
	OriginIPAddress            string   `storage_type:"ip" set:"y"`
	RequestID                  string   `storage_type:"keyword" set:"y"`
	RequestApplicationProtocol string   `storage_type:"keyword" set:"y"`
	RequestTransportProtocol   string   `storage_type:"keyword" set:"y"`
	RequestURL                 string   `storage_type:"keyword" set:"y"`
	RequestReferrer            string   `storage_type:"keyword" set:"y"`
	RequestMethod              string   `storage_type:"keyword" set:"y"`
	RequestUserAgent           string   `storage_type:"keyword" set:"y"`
	RequestStatus              string   `storage_type:"keyword" set:"y"`
	RequestTook                int64    `storage_type:"long" set:"y"`
	RequestBytesIn             int64    `storage_type:"long" set:"y"`
	RequestBytesOut            int64    `storage_type:"long" set:"y"`
	RequestResults             int64    `storage_type:"long" set:"y"`
	RequestUser                string   `storage_type:"keyword" set:"y"`
	RequestUnit                string   `storage_type:"keyword" set:"y"`
	RequestOrganization        string   `storage_type:"keyword" set:"y"`
	SourceIPAddress            string   `storage_type:"ip" set:"y"`
	SourceMACAddress           string   `storage_type:"keyword" set:"y"`
	SourceDomain               string   `storage_type:"keyword" set:"y"`
	SourceDNSName              string   `storage_type:"keyword" set:"y"`
	SourcePort                 string   `storage_type:"keyword" set:"y"`
	DestinationIPAddress       string   `storage_type:"ip" set:"y"`
	DestinationMACAddress      string   `storage_type:"keyword" set:"y"`
	DestinationDomain          string   `storage_type:"keyword" set:"y"`
	DestinationDNSName         string   `storage_type:"keyword" set:"y"`
	DestinationPort            string   `storage_type:"keyword" set:"y"`
	UserString1                string   `storage_type:"keyword" set:"y"`
	UserString1Label           string   `storage_type:"keyword" set:"y"`
	UserString2                string   `storage_type:"keyword" set:"y"`
	UserString2Label           string   `storage_type:"keyword" set:"y"`
	UserString3                string   `storage_type:"keyword" set:"y"`
	UserString3Label           string   `storage_type:"keyword" set:"y"`
	UserString4                string   `storage_type:"keyword" set:"y"`
	UserString4Label           string   `storage_type:"keyword" set:"y"`
	UserString5                string   `storage_type:"keyword" set:"y"`
	UserString5Label           string   `storage_type:"keyword" set:"y"`
	UserString6                string   `storage_type:"keyword" set:"y"`
	UserString6Label           string   `storage_type:"keyword" set:"y"`
	UserString7                string   `storage_type:"keyword" set:"y"`
	UserString7Label           string   `storage_type:"keyword" set:"y"`
	UserString8                string   `storage_type:"keyword" set:"y"`
	UserString8Label           string   `storage_type:"keyword" set:"y"`
	UserInt1                   int64    `storage_type:"long" set:"y"`
	UserInt1Label              string   `storage_type:"keyword" set:"y"`
	UserInt2                   int64    `storage_type:"long" set:"y"`
	UserInt2Label              string   `storage_type:"keyword" set:"y"`
	UserInt3                   int64    `storage_type:"long" set:"y"`
	UserInt3Label              string   `storage_type:"keyword" set:"y"`
	UserInt4                   int64    `storage_type:"long" set:"y"`
	UserInt4Label              string   `storage_type:"keyword" set:"y"`
	UserInt5                   int64    `storage_type:"long" set:"y"`
	UserInt5Label              string   `storage_type:"keyword" set:"y"`
	UserInt6                   int64    `storage_type:"long" set:"y"`
	UserInt6Label              string   `storage_type:"keyword" set:"y"`
	UserInt7                   int64    `storage_type:"long" set:"y"`
	UserInt7Label              string   `storage_type:"keyword" set:"y"`
	UserInt8                   int64    `storage_type:"long" set:"y"`
	UserInt8Label              string   `storage_type:"keyword" set:"y"`
	UserFloat1                 float64  `storage_type:"double" set:"y"`
	UserFloat1Label            string   `storage_type:"keyword" set:"y"`
	UserFloat2                 float64  `storage_type:"double" set:"y"`
	UserFloat2Label            string   `storage_type:"keyword" set:"y"`
	UserFloat3                 float64  `storage_type:"double" set:"y"`
	UserFloat3Label            string   `storage_type:"keyword" set:"y"`
	UserFloat4                 float64  `storage_type:"double" set:"y"`
	UserFloat4Label            string   `storage_type:"keyword" set:"y"`
	UserFloat5                 float64  `storage_type:"double" set:"y"`
	UserFloat5Label            string   `storage_type:"keyword" set:"y"`
	UserFloat6                 float64  `storage_type:"double" set:"y"`
	UserFloat6Label            string   `storage_type:"keyword" set:"y"`
	UserFloat7                 float64  `storage_type:"double" set:"y"`
	UserFloat7Label            string   `storage_type:"keyword" set:"y"`
	UserFloat8                 float64  `storage_type:"double" set:"y"`
	UserFloat8Label            string   `storage_type:"keyword" set:"y"`
	UserTimestamp1             int64    `storage_type:"date" set:"y"`
	UserTimestamp1Label        string   `storage_type:"keyword" set:"y"`
	UserTimestamp2             int64    `storage_type:"date" set:"y"`
	UserTimestamp2Label        string   `storage_type:"keyword" set:"y"`
	UserTimestamp3             int64    `storage_type:"date" set:"y"`
	UserTimestamp3Label        string   `storage_type:"keyword" set:"y"`
	UserTimestamp4             int64    `storage_type:"date" set:"y"`
	UserTimestamp4Label        string   `storage_type:"keyword" set:"y"`
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
