package normalization

func (r *Event) SetID(id string) {
	r.ID = id
}

func (r *Event) SetAnyField(field string, value string) {
	if len(value) == 0 {
		return
	}
	switch field {
	case "Incident":
		r.Incident = StringToBool(value)
	case "Severity":
		r.Severity = StringToInt(value)
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

func (r *Event) SetAnyFieldBytes(field string, value []byte) {
	r.SetAnyField(field, BytesToString(value))
}

func (r *Event) SetIntField(field string, value int64) {
	switch field {
	case "Severity":
		r.Severity = value
	case "RequestBytesIn":
		r.RequestBytesIn = value
	case "RequestBytesOut":
		r.RequestBytesOut = value
	case "RequestResults":
		r.RequestResults = value
	}
}

func (r *Event) SetFloatField(field string, value float64) {}

func (r *Event) SetBoolField(field string, value bool) {}
