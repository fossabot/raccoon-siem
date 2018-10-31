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
	ID                         string        `json:",omitempty" storage_type:"keyword"`
	Incident                   bool          `storage_type:"boolean"`
	Correlated                 bool          `storage_type:"boolean"`
	ParentID                   string        `json:",omitempty" storage_type:"keyword"`
	Customer                   string        `json:",omitempty" storage_type:"keyword"`
	Code                       string        `json:",omitempty" storage_type:"keyword"`
	StartTime                  time.Time     `json:",omitempty" storage_type:"date"`
	EndTime                    time.Time     `json:",omitempty" storage_type:"date"`
	Message                    string        `json:",omitempty" storage_type:"text"`
	Details                    string        `json:",omitempty" storage_type:"text"`
	Trace                      string        `json:",omitempty" storage_type:"text"`
	Severity                   string        `json:",omitempty" storage_type:"keyword"`
	BaseEventCount             int           `json:",omitempty" storage_type:"integer"`
	AggregatedEventCount       int           `json:",omitempty" storage_type:"integer"`
	AggregationRuleName        string        `json:",omitempty" storage_type:"keyword"`
	CorrelationRuleName        string        `json:",omitempty" storage_type:"keyword"`
	OriginEventID              string        `json:",omitempty" storage_type:"keyword"`
	OriginTimestamp            time.Time     `json:",omitempty" storage_type:"date"`
	OriginEnvironment          string        `json:",omitempty" storage_type:"keyword"`
	OriginSeverity             string        `json:",omitempty" storage_type:"keyword"`
	OriginServiceName          string        `json:",omitempty" storage_type:"keyword"`
	OriginServiceVersion       string        `json:",omitempty" storage_type:"keyword"`
	OriginProcessName          string        `json:",omitempty" storage_type:"keyword"`
	OriginFileName             string        `json:",omitempty" storage_type:"keyword"`
	OriginDNSName              string        `json:",omitempty" storage_type:"keyword"`
	OriginIPAddress            string        `json:",omitempty" storage_type:"ip"`
	CollectorIPAddress         string        `json:",omitempty" storage_type:"ip"`
	CollectorMACAddress        string        `json:",omitempty" storage_type:"keyword"`
	CollectorDNSName           string        `json:",omitempty" storage_type:"keyword"`
	CorrelatorIPAddress        string        `json:",omitempty" storage_type:"ip"`
	CorrelatorMACAddress       string        `json:",omitempty" storage_type:"keyword"`
	CorrelatorDNSName          string        `json:",omitempty" storage_type:"keyword"`
	CorrelatorEventSpecID      string        `json:",omitempty" storage_type:"keyword"`
	StorageTimestamp           time.Time     `json:",omitempty" storage_type:"date"`
	RequestID                  string        `json:",omitempty" storage_type:"keyword"`
	RequestApplicationProtocol string        `json:",omitempty" storage_type:"keyword"`
	RequestTransportProtocol   string        `json:",omitempty" storage_type:"keyword"`
	RequestURL                 string        `json:",omitempty" storage_type:"keyword"`
	RequestReferrer            string        `json:",omitempty" storage_type:"keyword"`
	RequestMethod              string        `json:",omitempty" storage_type:"keyword"`
	RequestUserAgent           string        `json:",omitempty" storage_type:"keyword"`
	RequestStatus              int64         `json:",omitempty" storage_type:"integer"`
	RequestTook                time.Duration `json:",omitempty" storage_type:"long"`
	RequestBytesIn             int64         `json:",omitempty" storage_type:"long"`
	RequestBytesOut            int64         `json:",omitempty" storage_type:"long"`
	RequestResults             int64         `json:",omitempty" storage_type:"long"`
	RequestUser                string        `json:",omitempty" storage_type:"keyword"`
	RequestUnit                string        `json:",omitempty" storage_type:"keyword"`
	RequestOrganization        string        `json:",omitempty" storage_type:"keyword"`
	RequestDBMSOperations      string        `json:",omitempty" storage_type:"keyword"`
	ResponseCached             bool          `json:",omitempty" storage_type:"boolean"`
	SourceIPAddress            string        `json:",omitempty" storage_type:"ip"`
	SourceMACAddress           string        `json:",omitempty" storage_type:"keyword"`
	SourceDNSName              string        `json:",omitempty" storage_type:"keyword"`
	SourcePort                 int64         `json:",omitempty" storage_type:"integer"`
	DestinationIPAddress       string        `json:",omitempty" storage_type:"ip"`
	DestinationMACAddress      string        `json:",omitempty" storage_type:"keyword"`
	DestinationDNSName         string        `json:",omitempty" storage_type:"keyword"`
	DestinationPort            int64         `json:",omitempty" storage_type:"integer"`
	baseEvents                 []*Event
}

func (e *Event) SetID(id string) {
	e.ID = id
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
