package normalization

func (r *Event) GetAnyField(field string) interface{} {
	switch field {
	case "Incident":
		return r.Incident
	case "Correlated":
		return r.Correlated
	case "CorrelationRuleName":
		return r.CorrelationRuleName
	case "Timestamp":
		return r.Timestamp
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
	case "RequestBytesIn":
		return r.RequestBytesIn
	case "RequestBytesOut":
		return r.RequestBytesOut
	case "RequestResults":
		return r.RequestResults
	case "Timestamp":
		return r.Timestamp
	case "OriginTimestamp":
		return r.OriginTimestamp
	case "RequestTook":
		return r.RequestTook
	default:
		return 0
	}
}

func (r *Event) GetFloatField(field string) float64 {
	return 0
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
