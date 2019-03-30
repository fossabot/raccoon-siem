package normalization

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
	case "CorrelatorEventSpecID":
		return r.CorrelatorEventSpecID
	case "SourceID":
		return r.SourceID
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
	case "SourceDNSName":
		return r.SourceDNSName
	case "SourcePort":
		return r.SourcePort
	case "DestinationIPAddress":
		return r.DestinationIPAddress
	case "DestinationMACAddress":
		return r.DestinationMACAddress
	case "DestinationDNSName":
		return r.DestinationDNSName
	case "DestinationPort":
		return r.DestinationPort
	default:
		return nil
	}
}

func (r *Event) GetIntField(field string) int64 {
	switch field {
	case "Timestamp":
		return r.Timestamp
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
	default:
		return 0
	}
}

func (r *Event) GetFloatField(field string) float64 {
	switch field {
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

func (r *Event) SetAnyField(field string, value string) {
	if len(value) == 0 {
		return
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
		r.Customer = value
	case "Code":
		r.Code = value
	case "Message":
		r.Message = value
	case "Details":
		r.Details = value
	case "Trace":
		r.Trace = value
	case "OriginEventID":
		r.OriginEventID = value
	case "OriginTimestamp":
		r.OriginTimestamp = StringToTime(value)
	case "OriginEnvironment":
		r.OriginEnvironment = value
	case "OriginSeverity":
		r.OriginSeverity = value
	case "OriginServiceName":
		r.OriginServiceName = value
	case "OriginServiceVersion":
		r.OriginServiceVersion = value
	case "OriginProcessName":
		r.OriginProcessName = value
	case "OriginFileName":
		r.OriginFileName = value
	case "OriginDNSName":
		r.OriginDNSName = value
	case "OriginIPAddress":
		r.OriginIPAddress = value
	case "RequestID":
		r.RequestID = value
	case "RequestApplicationProtocol":
		r.RequestApplicationProtocol = value
	case "RequestTransportProtocol":
		r.RequestTransportProtocol = value
	case "RequestURL":
		r.RequestURL = value
	case "RequestReferrer":
		r.RequestReferrer = value
	case "RequestMethod":
		r.RequestMethod = value
	case "RequestUserAgent":
		r.RequestUserAgent = value
	case "RequestStatus":
		r.RequestStatus = value
	case "RequestTook":
		r.RequestTook = StringToInt(value)
	case "RequestBytesIn":
		r.RequestBytesIn = StringToInt(value)
	case "RequestBytesOut":
		r.RequestBytesOut = StringToInt(value)
	case "RequestResults":
		r.RequestResults = StringToInt(value)
	case "RequestUser":
		r.RequestUser = value
	case "RequestUnit":
		r.RequestUnit = value
	case "RequestOrganization":
		r.RequestOrganization = value
	case "SourceIPAddress":
		r.SourceIPAddress = value
	case "SourceMACAddress":
		r.SourceMACAddress = value
	case "SourceDNSName":
		r.SourceDNSName = value
	case "SourcePort":
		r.SourcePort = value
	case "DestinationIPAddress":
		r.DestinationIPAddress = value
	case "DestinationMACAddress":
		r.DestinationMACAddress = value
	case "DestinationDNSName":
		r.DestinationDNSName = value
	case "DestinationPort":
		r.DestinationPort = value
	}
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
	}
}

func (r *Event) SetFloatField(field string, value float64) {
	switch field {
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

func (r *Event) SetFloatField(field string, value float64) {
	switch field {
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
