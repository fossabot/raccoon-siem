package sdk

import (
	"encoding/gob"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	fieldTypeString = iota
	fieldTypeInt
	fieldTypeFloat
	fieldTypeBool
	fieldTypeTime
	fieldTypeDuration
)

const (
	defaultEventFieldsHash = "_"
)

var eventFieldTypeByName = make(map[string]byte)

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
			numFieldType = fieldTypeString
		case "int64":
			numFieldType = fieldTypeInt
		case "float64":
			numFieldType = fieldTypeFloat
		case "bool":
			numFieldType = fieldTypeBool
		case "Time":
			numFieldType = fieldTypeTime
		case "Duration":
			numFieldType = fieldTypeDuration
		default:
			continue
		}

		eventFieldTypeByName[fieldName] = numFieldType
	}
}

type Event struct {
	ID                         string `json:",omitempty"`
	Incident                   bool
	Correlated                 bool
	baseEvents                 []*Event
	ParentID                   string        `json:",omitempty"`
	Customer                   string        `json:",omitempty"`
	Code                       string        `json:",omitempty"`
	StartTime                  time.Time     `json:",omitempty"`
	EndTime                    time.Time     `json:",omitempty"`
	Message                    string        `json:",omitempty"`
	Details                    string        `json:",omitempty"`
	Trace                      string        `json:",omitempty"`
	Severity                   string        `json:",omitempty"`
	BaseEventCount             int           `json:",omitempty"`
	AggregatedEventCount       int           `json:",omitempty"`
	AggregationRuleName        string        `json:",omitempty"`
	CorrelationRuleName        string        `json:",omitempty"`
	OriginEventID              string        `json:",omitempty"`
	OriginTimestamp            time.Time     `json:",omitempty"`
	OriginEnvironment          string        `json:",omitempty"`
	OriginSeverity             string        `json:",omitempty"`
	OriginServiceName          string        `json:",omitempty"`
	OriginServiceVersion       string        `json:",omitempty"`
	OriginProcessName          string        `json:",omitempty"`
	OriginFileName             string        `json:",omitempty"`
	OriginDNSName              string        `json:",omitempty"`
	OriginZone                 string        `json:",omitempty"`
	CollectorIPAddress         string        `json:",omitempty"`
	CollectorMACAddress        string        `json:",omitempty"`
	CollectorDNSName           string        `json:",omitempty"`
	CorrelatorIPAddress        string        `json:",omitempty"`
	CorrelatorMACAddress       string        `json:",omitempty"`
	CorrelatorDNSName          string        `json:",omitempty"`
	CorrelatorEventSpecID      string        `json:",omitempty"`
	StorageTimestamp           time.Time     `json:",omitempty"`
	RequestID                  string        `json:",omitempty"`
	RequestApplicationProtocol string        `json:",omitempty"`
	RequestTransportProtocol   string        `json:",omitempty"`
	RequestURL                 string        `json:",omitempty"`
	RequestReferrer            string        `json:",omitempty"`
	RequestMethod              string        `json:",omitempty"`
	RequestUserAgent           string        `json:",omitempty"`
	RequestStatus              int64         `json:",omitempty"`
	RequestTook                time.Duration `json:",omitempty"`
	RequestBytesIn             int64         `json:",omitempty"`
	RequestBytesOut            int64         `json:",omitempty"`
	RequestResults             int64         `json:",omitempty"`
	RequestUser                string        `json:",omitempty"`
	RequestUnit                string        `json:",omitempty"`
	RequestOrganization        string        `json:",omitempty"`
	RequestDBMSOperations      string        `json:",omitempty"`
	ResponseCached             bool          `json:",omitempty"`
	SourceIPAddress            string        `json:",omitempty"`
	SourceMACAddress           string        `json:",omitempty"`
	SourceDNSName              string        `json:",omitempty"`
	SourcePort                 int64         `json:",omitempty"`
	DestinationIPAddress       string        `json:",omitempty"`
	DestinationMACAddress      string        `json:",omitempty"`
	DestinationDNSName         string        `json:",omitempty"`
	DestinationPort            int64         `json:",omitempty"`
}

func (e *Event) SetID(id string) {

}

func (e *Event) SetField(name string, value interface{}, timeUnit byte) {
	targetFieldType := eventFieldTypeByName[name]

	finalValue := convertValue(value, targetFieldType, timeUnit)

	reflectValue := reflect.ValueOf(finalValue)
	reflect.ValueOf(e).Elem().FieldByName(name).Set(reflectValue)
}

func (e *Event) SetFieldNoConversion(name string, value interface{}) {
	reflectValue := reflect.ValueOf(value)
	reflect.ValueOf(e).Elem().FieldByName(name).Set(reflectValue)
}

func (e *Event) GetField(name string) (value interface{}, fieldType byte) {
	value = reflect.ValueOf(e).Elem().FieldByName(name).Interface()
	fieldType = eventFieldTypeByName[name]
	return
}

func (e *Event) GetFieldNoType(name string) interface{} {
	return reflect.ValueOf(e).Elem().FieldByName(name).Interface()
}

func (e *Event) FieldEmpty(name string) bool {
	return reflect.ValueOf(e).Elem().FieldByName(name).IsValid()
}

func (e *Event) HashFields(fieldNames []string) string {
	if len(fieldNames) == 0 {
		return defaultEventFieldsHash
	}

	builder := strings.Builder{}

	for i := range fieldNames {
		val, typeName := e.GetField(fieldNames[i])
		switch typeName {
		case fieldTypeString:
			builder.WriteString(val.(string))
		case fieldTypeInt:
			builder.WriteString(strconv.FormatInt(val.(int64), 10))
		case fieldTypeFloat:
			builder.WriteString(strconv.FormatFloat(val.(float64), 'f', -1, 64))
		case fieldTypeTime:
			builder.WriteString(val.(time.Time).String())
		case fieldTypeDuration:
			builder.WriteString(val.(time.Duration).String())
		case fieldTypeBool:
			builder.WriteString(strconv.FormatBool(val.(bool)))
		}
	}

	return builder.String()
}

func (e *Event) ToJSON() ([]byte, error) {
	e.setStorageTS(time.Now())
	return json.Marshal(e)
}

func (e *Event) FromJSON(input []byte) error {
	return json.Unmarshal(input, e)
}

func (e *Event) ToBSON() ([]byte, error) {
	e.setStorageTS(time.Now())
	return bson.Marshal(e)
}

func (e *Event) FromBSON(input []byte) error {
	return bson.Unmarshal(input, e)
}

func (e *Event) String() string {
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return ""
	}
	return string(data)
}

func (e *Event) setStorageTS(ts time.Time) {
	e.StorageTimestamp = ts
	for _, be := range e.baseEvents {
		be.setStorageTS(ts)
	}
}
