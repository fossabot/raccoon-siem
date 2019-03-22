package normalization

import (
	"encoding/gob"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/vmihailenco/msgpack.v4"
	"reflect"
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
	DefaultEventFieldsHash = "_"
)

var EventFieldTypeByName = make(map[string]byte)

func init() {
	event := Event{}
	gob.Register(event)

	v := reflect.ValueOf(event)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldType := v.Field(i).Type().Name()

		var numFieldType byte

		switch fieldType {
		case "string":
			numFieldType = FieldTypeString
		case "int64":
			numFieldType = FieldTypeInt
		case "float64":
			numFieldType = FieldTypeFloat
		case "bool":
			numFieldType = FieldTypeBool
		case "Time":
			numFieldType = FieldTypeTime
		case "Duration":
			numFieldType = FieldTypeDuration
		default:
			continue
		}

		EventFieldTypeByName[fieldName] = numFieldType
	}
}

type Event struct {
	_msgpack                   struct{} `msgpack:",asArray"`
	ID                         string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	Incident                   bool     `storage_type:"boolean"`
	Correlated                 bool     `storage_type:"boolean"`
	Tag                        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	ParentID                   string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	Customer                   string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	Code                       string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	Timestamp                  int64    `json:",omitempty" msgpack:",omitempty" storage_type:"date"`
	Severity                   int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long"`
	Score                      int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long"`
	BaseEventCount             int      `json:",omitempty" msgpack:",omitempty" storage_type:"integer"`
	AggregatedEventCount       int      `json:",omitempty" msgpack:",omitempty" storage_type:"integer"`
	AggregationRuleName        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	CollectorIPAddress         string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip"`
	CollectorDNSName           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	CorrelationRuleName        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	CorrelatorIPAddress        string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip"`
	CorrelatorDNSName          string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	CorrelatorEventSpecID      string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	SourceID                   string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	BaseEventIDs               []string `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	BaseEvents                 []*Event `json:"-" msgpack:"-"`
	Message                    string   `json:",omitempty" msgpack:",omitempty" storage_type:"text"`
	Details                    string   `json:",omitempty" msgpack:",omitempty" storage_type:"text"`
	Trace                      string   `json:",omitempty" msgpack:",omitempty" storage_type:"text"`
	OriginEventID              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	OriginTimestamp            int64    `json:",omitempty" msgpack:",omitempty" storage_type:"date"`
	OriginEnvironment          string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	OriginSeverity             string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	OriginServiceName          string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	OriginServiceVersion       string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	OriginProcessName          string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	OriginFileName             string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	OriginDNSName              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	OriginIPAddress            string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip"`
	RequestID                  string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestApplicationProtocol string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestTransportProtocol   string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestURL                 string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestReferrer            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestMethod              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestUserAgent           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestStatus              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestTook                int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long"`
	RequestBytesIn             int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long"`
	RequestBytesOut            int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long"`
	RequestResults             int64    `json:",omitempty" msgpack:",omitempty" storage_type:"long"`
	RequestUser                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestUnit                string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	RequestOrganization        string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	SourceIPAddress            string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip"`
	SourceMACAddress           string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	SourceDNSName              string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	SourcePort                 string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	DestinationIPAddress       string   `json:",omitempty" msgpack:",omitempty" storage_type:"ip"`
	DestinationMACAddress      string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	DestinationDNSName         string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
	DestinationPort            string   `json:",omitempty" msgpack:",omitempty" storage_type:"keyword"`
}

func (r *Event) SetField(name string, value interface{}, timeUnit byte) {
	targetFieldType := EventFieldTypeByName[name]

	finalValue := ConvertValue(value, targetFieldType, timeUnit)

	reflectValue := reflect.ValueOf(finalValue)
	reflect.ValueOf(r).Elem().FieldByName(name).Set(reflectValue)
}

func (r *Event) SetFieldNoConversion(name string, value interface{}) {
	reflectValue := reflect.ValueOf(value)
	reflect.ValueOf(r).Elem().FieldByName(name).Set(reflectValue)
}

func (r *Event) GetField(name string) (value interface{}, fieldType byte) {
	value = reflect.ValueOf(r).Elem().FieldByName(name).Interface()
	fieldType = EventFieldTypeByName[name]
	return
}

func (r *Event) GetFieldNoType(name string) interface{} {
	return reflect.ValueOf(r).Elem().FieldByName(name).Interface()
}

func (r *Event) FieldEmpty(name string) bool {
	return reflect.ValueOf(r).Elem().FieldByName(name).IsValid()
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
	return json.Marshal(r)
}

func (r *Event) FromJSON(input []byte) error {
	return json.Unmarshal(input, r)
}

func (r *Event) ToBSON() ([]byte, error) {
	return bson.Marshal(r)
}

func (r *Event) FromBSON(input []byte) error {
	return bson.Unmarshal(input, r)
}

func (r *Event) ToMsgPack() ([]byte, error) {
	return msgpack.Marshal(r)
}

func (r *Event) FromMsgPack(input []byte) error {
	return msgpack.Unmarshal(input, r)
}

func (r *Event) String() string {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return ""
	}
	return string(data)
}
