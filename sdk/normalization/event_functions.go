package normalization

//
// THIS FILE IS GENERATED. DO NOT EDIT!
//

import (
	"github.com/francoispqt/gojay"
	"gopkg.in/vmihailenco/msgpack.v4"
	"strings"
)

func (r *Event) GetAnyField(field string) interface{} {
	switch field {
	case "ID":
		return r.ID
	case "Tag":
		return r.Tag
	case "Timestamp":
		return r.Timestamp
	case "BaseEventCount":
		return r.BaseEventCount
	case "AggregatedEventCount":
		return r.AggregatedEventCount
	case "AggregationRuleName":
		return r.AggregationRuleName
	case "CollectorIPAddress":
		return r.CollectorIPAddress
	case "CollectorDNSName":
		return r.CollectorDNSName
	case "CorrelationRuleName":
		return r.CorrelationRuleName
	case "CorrelatorIPAddress":
		return r.CorrelatorIPAddress
	case "CorrelatorDNSName":
		return r.CorrelatorDNSName
	case "SourceID":
		return r.SourceID
	case "FieldsNormalized":
		return r.FieldsNormalized
	case "BaseEventIDs":
		return r.BaseEventIDs
	case "Incident":
		return r.Incident
	case "Correlated":
		return r.Correlated
	case "Score":
		return r.Score
	case "Severity":
		return r.Severity
	case "Customer":
		return r.Customer
	case "Code":
		return r.Code
	case "Message":
		return r.Message
	case "Details":
		return r.Details
	case "Trace":
		return r.Trace
	case "OriginEventID":
		return r.OriginEventID
	case "OriginTimestamp":
		return r.OriginTimestamp
	case "OriginEnvironment":
		return r.OriginEnvironment
	case "OriginSeverity":
		return r.OriginSeverity
	case "OriginServiceName":
		return r.OriginServiceName
	case "OriginServiceVersion":
		return r.OriginServiceVersion
	case "OriginProcessName":
		return r.OriginProcessName
	case "OriginFileName":
		return r.OriginFileName
	case "OriginDNSName":
		return r.OriginDNSName
	case "OriginDomain":
		return r.OriginDomain
	case "OriginIPAddress":
		return r.OriginIPAddress
	case "RequestID":
		return r.RequestID
	case "RequestApplicationProtocol":
		return r.RequestApplicationProtocol
	case "RequestTransportProtocol":
		return r.RequestTransportProtocol
	case "RequestURL":
		return r.RequestURL
	case "RequestReferrer":
		return r.RequestReferrer
	case "RequestMethod":
		return r.RequestMethod
	case "RequestUserAgent":
		return r.RequestUserAgent
	case "RequestStatus":
		return r.RequestStatus
	case "RequestTook":
		return r.RequestTook
	case "RequestBytesIn":
		return r.RequestBytesIn
	case "RequestBytesOut":
		return r.RequestBytesOut
	case "RequestResults":
		return r.RequestResults
	case "RequestUser":
		return r.RequestUser
	case "RequestUnit":
		return r.RequestUnit
	case "RequestOrganization":
		return r.RequestOrganization
	case "SourceIPAddress":
		return r.SourceIPAddress
	case "SourceMACAddress":
		return r.SourceMACAddress
	case "SourceDomain":
		return r.SourceDomain
	case "SourceDNSName":
		return r.SourceDNSName
	case "SourcePort":
		return r.SourcePort
	case "DestinationIPAddress":
		return r.DestinationIPAddress
	case "DestinationMACAddress":
		return r.DestinationMACAddress
	case "DestinationDomain":
		return r.DestinationDomain
	case "DestinationDNSName":
		return r.DestinationDNSName
	case "DestinationPort":
		return r.DestinationPort
	case "UserString1":
		return r.UserString1
	case "UserString1Label":
		return r.UserString1Label
	case "UserString2":
		return r.UserString2
	case "UserString2Label":
		return r.UserString2Label
	case "UserString3":
		return r.UserString3
	case "UserString3Label":
		return r.UserString3Label
	case "UserString4":
		return r.UserString4
	case "UserString4Label":
		return r.UserString4Label
	case "UserString5":
		return r.UserString5
	case "UserString5Label":
		return r.UserString5Label
	case "UserString6":
		return r.UserString6
	case "UserString6Label":
		return r.UserString6Label
	case "UserString7":
		return r.UserString7
	case "UserString7Label":
		return r.UserString7Label
	case "UserString8":
		return r.UserString8
	case "UserString8Label":
		return r.UserString8Label
	case "UserInt1":
		return r.UserInt1
	case "UserInt1Label":
		return r.UserInt1Label
	case "UserInt2":
		return r.UserInt2
	case "UserInt2Label":
		return r.UserInt2Label
	case "UserInt3":
		return r.UserInt3
	case "UserInt3Label":
		return r.UserInt3Label
	case "UserInt4":
		return r.UserInt4
	case "UserInt4Label":
		return r.UserInt4Label
	case "UserInt5":
		return r.UserInt5
	case "UserInt5Label":
		return r.UserInt5Label
	case "UserInt6":
		return r.UserInt6
	case "UserInt6Label":
		return r.UserInt6Label
	case "UserInt7":
		return r.UserInt7
	case "UserInt7Label":
		return r.UserInt7Label
	case "UserInt8":
		return r.UserInt8
	case "UserInt8Label":
		return r.UserInt8Label
	case "UserFloat1":
		return r.UserFloat1
	case "UserFloat1Label":
		return r.UserFloat1Label
	case "UserFloat2":
		return r.UserFloat2
	case "UserFloat2Label":
		return r.UserFloat2Label
	case "UserFloat3":
		return r.UserFloat3
	case "UserFloat3Label":
		return r.UserFloat3Label
	case "UserFloat4":
		return r.UserFloat4
	case "UserFloat4Label":
		return r.UserFloat4Label
	case "UserFloat5":
		return r.UserFloat5
	case "UserFloat5Label":
		return r.UserFloat5Label
	case "UserFloat6":
		return r.UserFloat6
	case "UserFloat6Label":
		return r.UserFloat6Label
	case "UserFloat7":
		return r.UserFloat7
	case "UserFloat7Label":
		return r.UserFloat7Label
	case "UserFloat8":
		return r.UserFloat8
	case "UserFloat8Label":
		return r.UserFloat8Label
	case "UserTimestamp1":
		return r.UserTimestamp1
	case "UserTimestamp1Label":
		return r.UserTimestamp1Label
	case "UserTimestamp2":
		return r.UserTimestamp2
	case "UserTimestamp2Label":
		return r.UserTimestamp2Label
	case "UserTimestamp3":
		return r.UserTimestamp3
	case "UserTimestamp3Label":
		return r.UserTimestamp3Label
	case "UserTimestamp4":
		return r.UserTimestamp4
	case "UserTimestamp4Label":
		return r.UserTimestamp4Label
	default:
		return nil
	}
}

func (r *Event) GetIntField(field string) int64 {
	switch field {
	case "Timestamp":
		return r.Timestamp
	case "FieldsNormalized":
		return r.FieldsNormalized
	case "Score":
		return r.Score
	case "Severity":
		return r.Severity
	case "OriginTimestamp":
		return r.OriginTimestamp
	case "RequestTook":
		return r.RequestTook
	case "RequestBytesIn":
		return r.RequestBytesIn
	case "RequestBytesOut":
		return r.RequestBytesOut
	case "RequestResults":
		return r.RequestResults
	case "UserInt1":
		return r.UserInt1
	case "UserInt2":
		return r.UserInt2
	case "UserInt3":
		return r.UserInt3
	case "UserInt4":
		return r.UserInt4
	case "UserInt5":
		return r.UserInt5
	case "UserInt6":
		return r.UserInt6
	case "UserInt7":
		return r.UserInt7
	case "UserInt8":
		return r.UserInt8
	case "UserTimestamp1":
		return r.UserTimestamp1
	case "UserTimestamp2":
		return r.UserTimestamp2
	case "UserTimestamp3":
		return r.UserTimestamp3
	case "UserTimestamp4":
		return r.UserTimestamp4
	default:
		return 0
	}
}

func (r *Event) GetFloatField(field string) float64 {
	switch field {
	case "UserFloat1":
		return r.UserFloat1
	case "UserFloat2":
		return r.UserFloat2
	case "UserFloat3":
		return r.UserFloat3
	case "UserFloat4":
		return r.UserFloat4
	case "UserFloat5":
		return r.UserFloat5
	case "UserFloat6":
		return r.UserFloat6
	case "UserFloat7":
		return r.UserFloat7
	case "UserFloat8":
		return r.UserFloat8
	default:
		return 0
	}
}

func (r *Event) GetBoolField(field string) bool {
	switch field {
	case "Incident":
		return r.Incident
	case "Correlated":
		return r.Correlated
	default:
		return false
	}
}

func (r *Event) SetAnyField(field string, value string) bool {
	if len(value) == 0 {
		return false
	}

	switch field {
	case "Incident":
		r.Incident = StringToBool(value)
	case "Correlated":
		r.Correlated = StringToBool(value)
	case "Score":
		r.Score = StringToInt(value)
	case "Severity":
		r.Severity = StringToInt(value)
	case "Customer":
		r.Customer = strings.TrimSpace(value)
	case "Code":
		r.Code = strings.TrimSpace(value)
	case "Message":
		r.Message = strings.TrimSpace(value)
	case "Details":
		r.Details = strings.TrimSpace(value)
	case "Trace":
		r.Trace = strings.TrimSpace(value)
	case "OriginEventID":
		r.OriginEventID = strings.TrimSpace(value)
	case "OriginTimestamp":
		r.OriginTimestamp = StringToTime(value)
	case "OriginEnvironment":
		r.OriginEnvironment = strings.TrimSpace(value)
	case "OriginSeverity":
		r.OriginSeverity = strings.TrimSpace(value)
	case "OriginServiceName":
		r.OriginServiceName = strings.TrimSpace(value)
	case "OriginServiceVersion":
		r.OriginServiceVersion = strings.TrimSpace(value)
	case "OriginProcessName":
		r.OriginProcessName = strings.TrimSpace(value)
	case "OriginFileName":
		r.OriginFileName = strings.TrimSpace(value)
	case "OriginDNSName":
		r.OriginDNSName = strings.TrimSpace(value)
	case "OriginDomain":
		r.OriginDomain = strings.TrimSpace(value)
	case "OriginIPAddress":
		r.OriginIPAddress = strings.TrimSpace(value)
	case "RequestID":
		r.RequestID = strings.TrimSpace(value)
	case "RequestApplicationProtocol":
		r.RequestApplicationProtocol = strings.TrimSpace(value)
	case "RequestTransportProtocol":
		r.RequestTransportProtocol = strings.TrimSpace(value)
	case "RequestURL":
		r.RequestURL = strings.TrimSpace(value)
	case "RequestReferrer":
		r.RequestReferrer = strings.TrimSpace(value)
	case "RequestMethod":
		r.RequestMethod = strings.TrimSpace(value)
	case "RequestUserAgent":
		r.RequestUserAgent = strings.TrimSpace(value)
	case "RequestStatus":
		r.RequestStatus = strings.TrimSpace(value)
	case "RequestTook":
		r.RequestTook = StringToInt(value)
	case "RequestBytesIn":
		r.RequestBytesIn = StringToInt(value)
	case "RequestBytesOut":
		r.RequestBytesOut = StringToInt(value)
	case "RequestResults":
		r.RequestResults = StringToInt(value)
	case "RequestUser":
		r.RequestUser = strings.TrimSpace(value)
	case "RequestUnit":
		r.RequestUnit = strings.TrimSpace(value)
	case "RequestOrganization":
		r.RequestOrganization = strings.TrimSpace(value)
	case "SourceIPAddress":
		r.SourceIPAddress = strings.TrimSpace(value)
	case "SourceMACAddress":
		r.SourceMACAddress = strings.TrimSpace(value)
	case "SourceDomain":
		r.SourceDomain = strings.TrimSpace(value)
	case "SourceDNSName":
		r.SourceDNSName = strings.TrimSpace(value)
	case "SourcePort":
		r.SourcePort = strings.TrimSpace(value)
	case "DestinationIPAddress":
		r.DestinationIPAddress = strings.TrimSpace(value)
	case "DestinationMACAddress":
		r.DestinationMACAddress = strings.TrimSpace(value)
	case "DestinationDomain":
		r.DestinationDomain = strings.TrimSpace(value)
	case "DestinationDNSName":
		r.DestinationDNSName = strings.TrimSpace(value)
	case "DestinationPort":
		r.DestinationPort = strings.TrimSpace(value)
	case "UserString1":
		r.UserString1 = strings.TrimSpace(value)
	case "UserString1Label":
		r.UserString1Label = strings.TrimSpace(value)
	case "UserString2":
		r.UserString2 = strings.TrimSpace(value)
	case "UserString2Label":
		r.UserString2Label = strings.TrimSpace(value)
	case "UserString3":
		r.UserString3 = strings.TrimSpace(value)
	case "UserString3Label":
		r.UserString3Label = strings.TrimSpace(value)
	case "UserString4":
		r.UserString4 = strings.TrimSpace(value)
	case "UserString4Label":
		r.UserString4Label = strings.TrimSpace(value)
	case "UserString5":
		r.UserString5 = strings.TrimSpace(value)
	case "UserString5Label":
		r.UserString5Label = strings.TrimSpace(value)
	case "UserString6":
		r.UserString6 = strings.TrimSpace(value)
	case "UserString6Label":
		r.UserString6Label = strings.TrimSpace(value)
	case "UserString7":
		r.UserString7 = strings.TrimSpace(value)
	case "UserString7Label":
		r.UserString7Label = strings.TrimSpace(value)
	case "UserString8":
		r.UserString8 = strings.TrimSpace(value)
	case "UserString8Label":
		r.UserString8Label = strings.TrimSpace(value)
	case "UserInt1":
		r.UserInt1 = StringToInt(value)
	case "UserInt1Label":
		r.UserInt1Label = strings.TrimSpace(value)
	case "UserInt2":
		r.UserInt2 = StringToInt(value)
	case "UserInt2Label":
		r.UserInt2Label = strings.TrimSpace(value)
	case "UserInt3":
		r.UserInt3 = StringToInt(value)
	case "UserInt3Label":
		r.UserInt3Label = strings.TrimSpace(value)
	case "UserInt4":
		r.UserInt4 = StringToInt(value)
	case "UserInt4Label":
		r.UserInt4Label = strings.TrimSpace(value)
	case "UserInt5":
		r.UserInt5 = StringToInt(value)
	case "UserInt5Label":
		r.UserInt5Label = strings.TrimSpace(value)
	case "UserInt6":
		r.UserInt6 = StringToInt(value)
	case "UserInt6Label":
		r.UserInt6Label = strings.TrimSpace(value)
	case "UserInt7":
		r.UserInt7 = StringToInt(value)
	case "UserInt7Label":
		r.UserInt7Label = strings.TrimSpace(value)
	case "UserInt8":
		r.UserInt8 = StringToInt(value)
	case "UserInt8Label":
		r.UserInt8Label = strings.TrimSpace(value)
	case "UserFloat1":
		r.UserFloat1 = StringToFloat(value)
	case "UserFloat1Label":
		r.UserFloat1Label = strings.TrimSpace(value)
	case "UserFloat2":
		r.UserFloat2 = StringToFloat(value)
	case "UserFloat2Label":
		r.UserFloat2Label = strings.TrimSpace(value)
	case "UserFloat3":
		r.UserFloat3 = StringToFloat(value)
	case "UserFloat3Label":
		r.UserFloat3Label = strings.TrimSpace(value)
	case "UserFloat4":
		r.UserFloat4 = StringToFloat(value)
	case "UserFloat4Label":
		r.UserFloat4Label = strings.TrimSpace(value)
	case "UserFloat5":
		r.UserFloat5 = StringToFloat(value)
	case "UserFloat5Label":
		r.UserFloat5Label = strings.TrimSpace(value)
	case "UserFloat6":
		r.UserFloat6 = StringToFloat(value)
	case "UserFloat6Label":
		r.UserFloat6Label = strings.TrimSpace(value)
	case "UserFloat7":
		r.UserFloat7 = StringToFloat(value)
	case "UserFloat7Label":
		r.UserFloat7Label = strings.TrimSpace(value)
	case "UserFloat8":
		r.UserFloat8 = StringToFloat(value)
	case "UserFloat8Label":
		r.UserFloat8Label = strings.TrimSpace(value)
	case "UserTimestamp1":
		r.UserTimestamp1 = StringToTime(value)
	case "UserTimestamp1Label":
		r.UserTimestamp1Label = strings.TrimSpace(value)
	case "UserTimestamp2":
		r.UserTimestamp2 = StringToTime(value)
	case "UserTimestamp2Label":
		r.UserTimestamp2Label = strings.TrimSpace(value)
	case "UserTimestamp3":
		r.UserTimestamp3 = StringToTime(value)
	case "UserTimestamp3Label":
		r.UserTimestamp3Label = strings.TrimSpace(value)
	case "UserTimestamp4":
		r.UserTimestamp4 = StringToTime(value)
	case "UserTimestamp4Label":
		r.UserTimestamp4Label = strings.TrimSpace(value)
	default:
		return false
	}

	return true
}

func (r *Event) SetIntField(field string, value int64) {
	switch field {
	case "Score":
		r.Score = value
	case "Severity":
		r.Severity = value
	case "OriginTimestamp":
		r.OriginTimestamp = value
	case "RequestTook":
		r.RequestTook = value
	case "RequestBytesIn":
		r.RequestBytesIn = value
	case "RequestBytesOut":
		r.RequestBytesOut = value
	case "RequestResults":
		r.RequestResults = value
	case "UserInt1":
		r.UserInt1 = value
	case "UserInt2":
		r.UserInt2 = value
	case "UserInt3":
		r.UserInt3 = value
	case "UserInt4":
		r.UserInt4 = value
	case "UserInt5":
		r.UserInt5 = value
	case "UserInt6":
		r.UserInt6 = value
	case "UserInt7":
		r.UserInt7 = value
	case "UserInt8":
		r.UserInt8 = value
	case "UserTimestamp1":
		r.UserTimestamp1 = value
	case "UserTimestamp2":
		r.UserTimestamp2 = value
	case "UserTimestamp3":
		r.UserTimestamp3 = value
	case "UserTimestamp4":
		r.UserTimestamp4 = value
	}
}

func (r *Event) SetFloatField(field string, value float64) {
	switch field {
	case "UserFloat1":
		r.UserFloat1 = value
	case "UserFloat2":
		r.UserFloat2 = value
	case "UserFloat3":
		r.UserFloat3 = value
	case "UserFloat4":
		r.UserFloat4 = value
	case "UserFloat5":
		r.UserFloat5 = value
	case "UserFloat6":
		r.UserFloat6 = value
	case "UserFloat7":
		r.UserFloat7 = value
	case "UserFloat8":
		r.UserFloat8 = value
	}
}

func (r *Event) SetBoolField(field string, value bool) {
	switch field {
	case "Incident":
		r.Incident = value
	case "Correlated":
		r.Correlated = value
	}
}

func (r *Event) EncodeMsgpack(enc *msgpack.Encoder) error {
	if err := enc.EncodeString(r.ID); err != nil {
		return err
	}

	if err := enc.EncodeString(r.Tag); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.Timestamp); err != nil {
		return err
	}

	if err := enc.EncodeInt64(int64(r.BaseEventCount)); err != nil {
		return err
	}

	if err := enc.EncodeInt64(int64(r.AggregatedEventCount)); err != nil {
		return err
	}

	if err := enc.EncodeString(r.AggregationRuleName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.CollectorIPAddress); err != nil {
		return err
	}

	if err := enc.EncodeString(r.CollectorDNSName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.CorrelationRuleName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.CorrelatorIPAddress); err != nil {
		return err
	}

	if err := enc.EncodeString(r.CorrelatorDNSName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.SourceID); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.FieldsNormalized); err != nil {
		return err
	}

	if err := enc.EncodeArrayLen(len(r.BaseEventIDs)); err != nil {
		return err
	}

	for i := range r.BaseEventIDs {
		if err := enc.EncodeString(r.BaseEventIDs[i]); err != nil {
			return err
		}
	}

	if err := enc.EncodeBool(r.Incident); err != nil {
		return err
	}

	if err := enc.EncodeBool(r.Correlated); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.Score); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.Severity); err != nil {
		return err
	}

	if err := enc.EncodeString(r.Customer); err != nil {
		return err
	}

	if err := enc.EncodeString(r.Code); err != nil {
		return err
	}

	if err := enc.EncodeString(r.Message); err != nil {
		return err
	}

	if err := enc.EncodeString(r.Details); err != nil {
		return err
	}

	if err := enc.EncodeString(r.Trace); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginEventID); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.OriginTimestamp); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginEnvironment); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginSeverity); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginServiceName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginServiceVersion); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginProcessName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginFileName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginDNSName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginDomain); err != nil {
		return err
	}

	if err := enc.EncodeString(r.OriginIPAddress); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestID); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestApplicationProtocol); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestTransportProtocol); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestURL); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestReferrer); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestMethod); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestUserAgent); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestStatus); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.RequestTook); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.RequestBytesIn); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.RequestBytesOut); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.RequestResults); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestUser); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestUnit); err != nil {
		return err
	}

	if err := enc.EncodeString(r.RequestOrganization); err != nil {
		return err
	}

	if err := enc.EncodeString(r.SourceIPAddress); err != nil {
		return err
	}

	if err := enc.EncodeString(r.SourceMACAddress); err != nil {
		return err
	}

	if err := enc.EncodeString(r.SourceDomain); err != nil {
		return err
	}

	if err := enc.EncodeString(r.SourceDNSName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.SourcePort); err != nil {
		return err
	}

	if err := enc.EncodeString(r.DestinationIPAddress); err != nil {
		return err
	}

	if err := enc.EncodeString(r.DestinationMACAddress); err != nil {
		return err
	}

	if err := enc.EncodeString(r.DestinationDomain); err != nil {
		return err
	}

	if err := enc.EncodeString(r.DestinationDNSName); err != nil {
		return err
	}

	if err := enc.EncodeString(r.DestinationPort); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString1); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString1Label); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString2); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString2Label); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString3); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString3Label); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString4); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString4Label); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString5); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString5Label); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString6); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString6Label); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString7); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString7Label); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString8); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserString8Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserInt1); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserInt1Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserInt2); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserInt2Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserInt3); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserInt3Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserInt4); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserInt4Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserInt5); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserInt5Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserInt6); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserInt6Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserInt7); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserInt7Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserInt8); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserInt8Label); err != nil {
		return err
	}

	if err := enc.EncodeFloat64(r.UserFloat1); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserFloat1Label); err != nil {
		return err
	}

	if err := enc.EncodeFloat64(r.UserFloat2); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserFloat2Label); err != nil {
		return err
	}

	if err := enc.EncodeFloat64(r.UserFloat3); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserFloat3Label); err != nil {
		return err
	}

	if err := enc.EncodeFloat64(r.UserFloat4); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserFloat4Label); err != nil {
		return err
	}

	if err := enc.EncodeFloat64(r.UserFloat5); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserFloat5Label); err != nil {
		return err
	}

	if err := enc.EncodeFloat64(r.UserFloat6); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserFloat6Label); err != nil {
		return err
	}

	if err := enc.EncodeFloat64(r.UserFloat7); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserFloat7Label); err != nil {
		return err
	}

	if err := enc.EncodeFloat64(r.UserFloat8); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserFloat8Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserTimestamp1); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserTimestamp1Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserTimestamp2); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserTimestamp2Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserTimestamp3); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserTimestamp3Label); err != nil {
		return err
	}

	if err := enc.EncodeInt64(r.UserTimestamp4); err != nil {
		return err
	}

	if err := enc.EncodeString(r.UserTimestamp4Label); err != nil {
		return err
	}

	return nil
}

func (r *Event) DecodeMsgpack(dec *msgpack.Decoder) (err error) {
	if r.ID, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.Tag, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.Timestamp, err = dec.DecodeInt64(); err != nil {
		return err
	}

	BaseEventCount, err := dec.DecodeInt64()
	if err != nil {
		return err
	}
	r.BaseEventCount = int(BaseEventCount)

	AggregatedEventCount, err := dec.DecodeInt64()
	if err != nil {
		return err
	}
	r.AggregatedEventCount = int(AggregatedEventCount)

	if r.AggregationRuleName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.CollectorIPAddress, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.CollectorDNSName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.CorrelationRuleName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.CorrelatorIPAddress, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.CorrelatorDNSName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.SourceID, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.FieldsNormalized, err = dec.DecodeInt64(); err != nil {
		return err
	}

	l, err := dec.DecodeArrayLen()
	if err != nil {
		return err
	}

	r.BaseEventIDs = make(strSlice, l)
	for i := 0; i < l; i++ {
		r.BaseEventIDs[i], err = dec.DecodeString()
		if err != nil {
			return err
		}
	}

	if r.Incident, err = dec.DecodeBool(); err != nil {
		return err
	}

	if r.Correlated, err = dec.DecodeBool(); err != nil {
		return err
	}

	if r.Score, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.Severity, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.Customer, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.Code, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.Message, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.Details, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.Trace, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginEventID, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginTimestamp, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.OriginEnvironment, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginSeverity, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginServiceName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginServiceVersion, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginProcessName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginFileName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginDNSName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginDomain, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.OriginIPAddress, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestID, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestApplicationProtocol, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestTransportProtocol, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestURL, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestReferrer, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestMethod, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestUserAgent, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestStatus, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestTook, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.RequestBytesIn, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.RequestBytesOut, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.RequestResults, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.RequestUser, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestUnit, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.RequestOrganization, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.SourceIPAddress, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.SourceMACAddress, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.SourceDomain, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.SourceDNSName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.SourcePort, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.DestinationIPAddress, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.DestinationMACAddress, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.DestinationDomain, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.DestinationDNSName, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.DestinationPort, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString1, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString1Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString2, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString2Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString3, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString3Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString4, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString4Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString5, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString5Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString6, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString6Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString7, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString7Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString8, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserString8Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserInt1, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserInt1Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserInt2, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserInt2Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserInt3, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserInt3Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserInt4, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserInt4Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserInt5, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserInt5Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserInt6, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserInt6Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserInt7, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserInt7Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserInt8, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserInt8Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserFloat1, err = dec.DecodeFloat64(); err != nil {
		return err
	}

	if r.UserFloat1Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserFloat2, err = dec.DecodeFloat64(); err != nil {
		return err
	}

	if r.UserFloat2Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserFloat3, err = dec.DecodeFloat64(); err != nil {
		return err
	}

	if r.UserFloat3Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserFloat4, err = dec.DecodeFloat64(); err != nil {
		return err
	}

	if r.UserFloat4Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserFloat5, err = dec.DecodeFloat64(); err != nil {
		return err
	}

	if r.UserFloat5Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserFloat6, err = dec.DecodeFloat64(); err != nil {
		return err
	}

	if r.UserFloat6Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserFloat7, err = dec.DecodeFloat64(); err != nil {
		return err
	}

	if r.UserFloat7Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserFloat8, err = dec.DecodeFloat64(); err != nil {
		return err
	}

	if r.UserFloat8Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserTimestamp1, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserTimestamp1Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserTimestamp2, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserTimestamp2Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserTimestamp3, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserTimestamp3Label, err = dec.DecodeString(); err != nil {
		return err
	}

	if r.UserTimestamp4, err = dec.DecodeInt64(); err != nil {
		return err
	}

	if r.UserTimestamp4Label, err = dec.DecodeString(); err != nil {
		return err
	}

	return nil
}

func (r *Event) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty("ID", r.ID)
	enc.StringKeyOmitEmpty("Tag", r.Tag)
	enc.Int64KeyOmitEmpty("Timestamp", r.Timestamp)
	enc.IntKeyOmitEmpty("BaseEventCount", r.BaseEventCount)
	enc.IntKeyOmitEmpty("AggregatedEventCount", r.AggregatedEventCount)
	enc.StringKeyOmitEmpty("AggregationRuleName", r.AggregationRuleName)
	enc.StringKeyOmitEmpty("CollectorIPAddress", r.CollectorIPAddress)
	enc.StringKeyOmitEmpty("CollectorDNSName", r.CollectorDNSName)
	enc.StringKeyOmitEmpty("CorrelationRuleName", r.CorrelationRuleName)
	enc.StringKeyOmitEmpty("CorrelatorIPAddress", r.CorrelatorIPAddress)
	enc.StringKeyOmitEmpty("CorrelatorDNSName", r.CorrelatorDNSName)
	enc.StringKeyOmitEmpty("SourceID", r.SourceID)
	enc.Int64KeyOmitEmpty("FieldsNormalized", r.FieldsNormalized)
	enc.ArrayKeyOmitEmpty("BaseEventIDs", r.BaseEventIDs)
	enc.BoolKeyOmitEmpty("Incident", r.Incident)
	enc.BoolKeyOmitEmpty("Correlated", r.Correlated)
	enc.Int64KeyOmitEmpty("Score", r.Score)
	enc.Int64KeyOmitEmpty("Severity", r.Severity)
	enc.StringKeyOmitEmpty("Customer", r.Customer)
	enc.StringKeyOmitEmpty("Code", r.Code)
	enc.StringKeyOmitEmpty("Message", r.Message)
	enc.StringKeyOmitEmpty("Details", r.Details)
	enc.StringKeyOmitEmpty("Trace", r.Trace)
	enc.StringKeyOmitEmpty("OriginEventID", r.OriginEventID)
	enc.Int64KeyOmitEmpty("OriginTimestamp", r.OriginTimestamp)
	enc.StringKeyOmitEmpty("OriginEnvironment", r.OriginEnvironment)
	enc.StringKeyOmitEmpty("OriginSeverity", r.OriginSeverity)
	enc.StringKeyOmitEmpty("OriginServiceName", r.OriginServiceName)
	enc.StringKeyOmitEmpty("OriginServiceVersion", r.OriginServiceVersion)
	enc.StringKeyOmitEmpty("OriginProcessName", r.OriginProcessName)
	enc.StringKeyOmitEmpty("OriginFileName", r.OriginFileName)
	enc.StringKeyOmitEmpty("OriginDNSName", r.OriginDNSName)
	enc.StringKeyOmitEmpty("OriginDomain", r.OriginDomain)
	enc.StringKeyOmitEmpty("OriginIPAddress", r.OriginIPAddress)
	enc.StringKeyOmitEmpty("RequestID", r.RequestID)
	enc.StringKeyOmitEmpty("RequestApplicationProtocol", r.RequestApplicationProtocol)
	enc.StringKeyOmitEmpty("RequestTransportProtocol", r.RequestTransportProtocol)
	enc.StringKeyOmitEmpty("RequestURL", r.RequestURL)
	enc.StringKeyOmitEmpty("RequestReferrer", r.RequestReferrer)
	enc.StringKeyOmitEmpty("RequestMethod", r.RequestMethod)
	enc.StringKeyOmitEmpty("RequestUserAgent", r.RequestUserAgent)
	enc.StringKeyOmitEmpty("RequestStatus", r.RequestStatus)
	enc.Int64KeyOmitEmpty("RequestTook", r.RequestTook)
	enc.Int64KeyOmitEmpty("RequestBytesIn", r.RequestBytesIn)
	enc.Int64KeyOmitEmpty("RequestBytesOut", r.RequestBytesOut)
	enc.Int64KeyOmitEmpty("RequestResults", r.RequestResults)
	enc.StringKeyOmitEmpty("RequestUser", r.RequestUser)
	enc.StringKeyOmitEmpty("RequestUnit", r.RequestUnit)
	enc.StringKeyOmitEmpty("RequestOrganization", r.RequestOrganization)
	enc.StringKeyOmitEmpty("SourceIPAddress", r.SourceIPAddress)
	enc.StringKeyOmitEmpty("SourceMACAddress", r.SourceMACAddress)
	enc.StringKeyOmitEmpty("SourceDomain", r.SourceDomain)
	enc.StringKeyOmitEmpty("SourceDNSName", r.SourceDNSName)
	enc.StringKeyOmitEmpty("SourcePort", r.SourcePort)
	enc.StringKeyOmitEmpty("DestinationIPAddress", r.DestinationIPAddress)
	enc.StringKeyOmitEmpty("DestinationMACAddress", r.DestinationMACAddress)
	enc.StringKeyOmitEmpty("DestinationDomain", r.DestinationDomain)
	enc.StringKeyOmitEmpty("DestinationDNSName", r.DestinationDNSName)
	enc.StringKeyOmitEmpty("DestinationPort", r.DestinationPort)
	enc.StringKeyOmitEmpty("UserString1", r.UserString1)
	enc.StringKeyOmitEmpty("UserString1Label", r.UserString1Label)
	enc.StringKeyOmitEmpty("UserString2", r.UserString2)
	enc.StringKeyOmitEmpty("UserString2Label", r.UserString2Label)
	enc.StringKeyOmitEmpty("UserString3", r.UserString3)
	enc.StringKeyOmitEmpty("UserString3Label", r.UserString3Label)
	enc.StringKeyOmitEmpty("UserString4", r.UserString4)
	enc.StringKeyOmitEmpty("UserString4Label", r.UserString4Label)
	enc.StringKeyOmitEmpty("UserString5", r.UserString5)
	enc.StringKeyOmitEmpty("UserString5Label", r.UserString5Label)
	enc.StringKeyOmitEmpty("UserString6", r.UserString6)
	enc.StringKeyOmitEmpty("UserString6Label", r.UserString6Label)
	enc.StringKeyOmitEmpty("UserString7", r.UserString7)
	enc.StringKeyOmitEmpty("UserString7Label", r.UserString7Label)
	enc.StringKeyOmitEmpty("UserString8", r.UserString8)
	enc.StringKeyOmitEmpty("UserString8Label", r.UserString8Label)
	enc.Int64KeyOmitEmpty("UserInt1", r.UserInt1)
	enc.StringKeyOmitEmpty("UserInt1Label", r.UserInt1Label)
	enc.Int64KeyOmitEmpty("UserInt2", r.UserInt2)
	enc.StringKeyOmitEmpty("UserInt2Label", r.UserInt2Label)
	enc.Int64KeyOmitEmpty("UserInt3", r.UserInt3)
	enc.StringKeyOmitEmpty("UserInt3Label", r.UserInt3Label)
	enc.Int64KeyOmitEmpty("UserInt4", r.UserInt4)
	enc.StringKeyOmitEmpty("UserInt4Label", r.UserInt4Label)
	enc.Int64KeyOmitEmpty("UserInt5", r.UserInt5)
	enc.StringKeyOmitEmpty("UserInt5Label", r.UserInt5Label)
	enc.Int64KeyOmitEmpty("UserInt6", r.UserInt6)
	enc.StringKeyOmitEmpty("UserInt6Label", r.UserInt6Label)
	enc.Int64KeyOmitEmpty("UserInt7", r.UserInt7)
	enc.StringKeyOmitEmpty("UserInt7Label", r.UserInt7Label)
	enc.Int64KeyOmitEmpty("UserInt8", r.UserInt8)
	enc.StringKeyOmitEmpty("UserInt8Label", r.UserInt8Label)
	enc.Float64KeyOmitEmpty("UserFloat1", r.UserFloat1)
	enc.StringKeyOmitEmpty("UserFloat1Label", r.UserFloat1Label)
	enc.Float64KeyOmitEmpty("UserFloat2", r.UserFloat2)
	enc.StringKeyOmitEmpty("UserFloat2Label", r.UserFloat2Label)
	enc.Float64KeyOmitEmpty("UserFloat3", r.UserFloat3)
	enc.StringKeyOmitEmpty("UserFloat3Label", r.UserFloat3Label)
	enc.Float64KeyOmitEmpty("UserFloat4", r.UserFloat4)
	enc.StringKeyOmitEmpty("UserFloat4Label", r.UserFloat4Label)
	enc.Float64KeyOmitEmpty("UserFloat5", r.UserFloat5)
	enc.StringKeyOmitEmpty("UserFloat5Label", r.UserFloat5Label)
	enc.Float64KeyOmitEmpty("UserFloat6", r.UserFloat6)
	enc.StringKeyOmitEmpty("UserFloat6Label", r.UserFloat6Label)
	enc.Float64KeyOmitEmpty("UserFloat7", r.UserFloat7)
	enc.StringKeyOmitEmpty("UserFloat7Label", r.UserFloat7Label)
	enc.Float64KeyOmitEmpty("UserFloat8", r.UserFloat8)
	enc.StringKeyOmitEmpty("UserFloat8Label", r.UserFloat8Label)
	enc.Int64KeyOmitEmpty("UserTimestamp1", r.UserTimestamp1)
	enc.StringKeyOmitEmpty("UserTimestamp1Label", r.UserTimestamp1Label)
	enc.Int64KeyOmitEmpty("UserTimestamp2", r.UserTimestamp2)
	enc.StringKeyOmitEmpty("UserTimestamp2Label", r.UserTimestamp2Label)
	enc.Int64KeyOmitEmpty("UserTimestamp3", r.UserTimestamp3)
	enc.StringKeyOmitEmpty("UserTimestamp3Label", r.UserTimestamp3Label)
	enc.Int64KeyOmitEmpty("UserTimestamp4", r.UserTimestamp4)
	enc.StringKeyOmitEmpty("UserTimestamp4Label", r.UserTimestamp4Label)
}

func (r *Event) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "ID":
		return dec.String(&r.ID)
	case "Tag":
		return dec.String(&r.Tag)
	case "Timestamp":
		return dec.Int64(&r.Timestamp)
	case "BaseEventCount":
		return dec.Int(&r.BaseEventCount)
	case "AggregatedEventCount":
		return dec.Int(&r.AggregatedEventCount)
	case "AggregationRuleName":
		return dec.String(&r.AggregationRuleName)
	case "CollectorIPAddress":
		return dec.String(&r.CollectorIPAddress)
	case "CollectorDNSName":
		return dec.String(&r.CollectorDNSName)
	case "CorrelationRuleName":
		return dec.String(&r.CorrelationRuleName)
	case "CorrelatorIPAddress":
		return dec.String(&r.CorrelatorIPAddress)
	case "CorrelatorDNSName":
		return dec.String(&r.CorrelatorDNSName)
	case "SourceID":
		return dec.String(&r.SourceID)
	case "FieldsNormalized":
		return dec.Int64(&r.FieldsNormalized)
	case "BaseEventIDs":
		return dec.Array(&r.BaseEventIDs)
	case "Incident":
		return dec.Bool(&r.Incident)
	case "Correlated":
		return dec.Bool(&r.Correlated)
	case "Score":
		return dec.Int64(&r.Score)
	case "Severity":
		return dec.Int64(&r.Severity)
	case "Customer":
		return dec.String(&r.Customer)
	case "Code":
		return dec.String(&r.Code)
	case "Message":
		return dec.String(&r.Message)
	case "Details":
		return dec.String(&r.Details)
	case "Trace":
		return dec.String(&r.Trace)
	case "OriginEventID":
		return dec.String(&r.OriginEventID)
	case "OriginTimestamp":
		return dec.Int64(&r.OriginTimestamp)
	case "OriginEnvironment":
		return dec.String(&r.OriginEnvironment)
	case "OriginSeverity":
		return dec.String(&r.OriginSeverity)
	case "OriginServiceName":
		return dec.String(&r.OriginServiceName)
	case "OriginServiceVersion":
		return dec.String(&r.OriginServiceVersion)
	case "OriginProcessName":
		return dec.String(&r.OriginProcessName)
	case "OriginFileName":
		return dec.String(&r.OriginFileName)
	case "OriginDNSName":
		return dec.String(&r.OriginDNSName)
	case "OriginDomain":
		return dec.String(&r.OriginDomain)
	case "OriginIPAddress":
		return dec.String(&r.OriginIPAddress)
	case "RequestID":
		return dec.String(&r.RequestID)
	case "RequestApplicationProtocol":
		return dec.String(&r.RequestApplicationProtocol)
	case "RequestTransportProtocol":
		return dec.String(&r.RequestTransportProtocol)
	case "RequestURL":
		return dec.String(&r.RequestURL)
	case "RequestReferrer":
		return dec.String(&r.RequestReferrer)
	case "RequestMethod":
		return dec.String(&r.RequestMethod)
	case "RequestUserAgent":
		return dec.String(&r.RequestUserAgent)
	case "RequestStatus":
		return dec.String(&r.RequestStatus)
	case "RequestTook":
		return dec.Int64(&r.RequestTook)
	case "RequestBytesIn":
		return dec.Int64(&r.RequestBytesIn)
	case "RequestBytesOut":
		return dec.Int64(&r.RequestBytesOut)
	case "RequestResults":
		return dec.Int64(&r.RequestResults)
	case "RequestUser":
		return dec.String(&r.RequestUser)
	case "RequestUnit":
		return dec.String(&r.RequestUnit)
	case "RequestOrganization":
		return dec.String(&r.RequestOrganization)
	case "SourceIPAddress":
		return dec.String(&r.SourceIPAddress)
	case "SourceMACAddress":
		return dec.String(&r.SourceMACAddress)
	case "SourceDomain":
		return dec.String(&r.SourceDomain)
	case "SourceDNSName":
		return dec.String(&r.SourceDNSName)
	case "SourcePort":
		return dec.String(&r.SourcePort)
	case "DestinationIPAddress":
		return dec.String(&r.DestinationIPAddress)
	case "DestinationMACAddress":
		return dec.String(&r.DestinationMACAddress)
	case "DestinationDomain":
		return dec.String(&r.DestinationDomain)
	case "DestinationDNSName":
		return dec.String(&r.DestinationDNSName)
	case "DestinationPort":
		return dec.String(&r.DestinationPort)
	case "UserString1":
		return dec.String(&r.UserString1)
	case "UserString1Label":
		return dec.String(&r.UserString1Label)
	case "UserString2":
		return dec.String(&r.UserString2)
	case "UserString2Label":
		return dec.String(&r.UserString2Label)
	case "UserString3":
		return dec.String(&r.UserString3)
	case "UserString3Label":
		return dec.String(&r.UserString3Label)
	case "UserString4":
		return dec.String(&r.UserString4)
	case "UserString4Label":
		return dec.String(&r.UserString4Label)
	case "UserString5":
		return dec.String(&r.UserString5)
	case "UserString5Label":
		return dec.String(&r.UserString5Label)
	case "UserString6":
		return dec.String(&r.UserString6)
	case "UserString6Label":
		return dec.String(&r.UserString6Label)
	case "UserString7":
		return dec.String(&r.UserString7)
	case "UserString7Label":
		return dec.String(&r.UserString7Label)
	case "UserString8":
		return dec.String(&r.UserString8)
	case "UserString8Label":
		return dec.String(&r.UserString8Label)
	case "UserInt1":
		return dec.Int64(&r.UserInt1)
	case "UserInt1Label":
		return dec.String(&r.UserInt1Label)
	case "UserInt2":
		return dec.Int64(&r.UserInt2)
	case "UserInt2Label":
		return dec.String(&r.UserInt2Label)
	case "UserInt3":
		return dec.Int64(&r.UserInt3)
	case "UserInt3Label":
		return dec.String(&r.UserInt3Label)
	case "UserInt4":
		return dec.Int64(&r.UserInt4)
	case "UserInt4Label":
		return dec.String(&r.UserInt4Label)
	case "UserInt5":
		return dec.Int64(&r.UserInt5)
	case "UserInt5Label":
		return dec.String(&r.UserInt5Label)
	case "UserInt6":
		return dec.Int64(&r.UserInt6)
	case "UserInt6Label":
		return dec.String(&r.UserInt6Label)
	case "UserInt7":
		return dec.Int64(&r.UserInt7)
	case "UserInt7Label":
		return dec.String(&r.UserInt7Label)
	case "UserInt8":
		return dec.Int64(&r.UserInt8)
	case "UserInt8Label":
		return dec.String(&r.UserInt8Label)
	case "UserFloat1":
		return dec.Float64(&r.UserFloat1)
	case "UserFloat1Label":
		return dec.String(&r.UserFloat1Label)
	case "UserFloat2":
		return dec.Float64(&r.UserFloat2)
	case "UserFloat2Label":
		return dec.String(&r.UserFloat2Label)
	case "UserFloat3":
		return dec.Float64(&r.UserFloat3)
	case "UserFloat3Label":
		return dec.String(&r.UserFloat3Label)
	case "UserFloat4":
		return dec.Float64(&r.UserFloat4)
	case "UserFloat4Label":
		return dec.String(&r.UserFloat4Label)
	case "UserFloat5":
		return dec.Float64(&r.UserFloat5)
	case "UserFloat5Label":
		return dec.String(&r.UserFloat5Label)
	case "UserFloat6":
		return dec.Float64(&r.UserFloat6)
	case "UserFloat6Label":
		return dec.String(&r.UserFloat6Label)
	case "UserFloat7":
		return dec.Float64(&r.UserFloat7)
	case "UserFloat7Label":
		return dec.String(&r.UserFloat7Label)
	case "UserFloat8":
		return dec.Float64(&r.UserFloat8)
	case "UserFloat8Label":
		return dec.String(&r.UserFloat8Label)
	case "UserTimestamp1":
		return dec.Int64(&r.UserTimestamp1)
	case "UserTimestamp1Label":
		return dec.String(&r.UserTimestamp1Label)
	case "UserTimestamp2":
		return dec.Int64(&r.UserTimestamp2)
	case "UserTimestamp2Label":
		return dec.String(&r.UserTimestamp2Label)
	case "UserTimestamp3":
		return dec.Int64(&r.UserTimestamp3)
	case "UserTimestamp3Label":
		return dec.String(&r.UserTimestamp3Label)
	case "UserTimestamp4":
		return dec.Int64(&r.UserTimestamp4)
	case "UserTimestamp4Label":
		return dec.String(&r.UserTimestamp4Label)
	default:
		return nil
	}
}
