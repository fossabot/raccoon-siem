package cef

import (
	"github.com/tephrocactus/raccoon-siem/sdk/helpers"
)

const (
	space = ' '
	pipe  = '|'
	eq    = '='
	bs    = '\\'
)

var (
	entrySequence = []byte{'C', 'E', 'F', ':', '0', '|'}
	headerFields  = []string{
		"deviceVendor",
		"deviceProduct",
		"deviceVersion",
		"deviceEventClassId",
		"name",
		"severity",
	}
	dict = map[string]string{
		"act":                                 "deviceAction",
		"app":                                 "applicationProtocol",
		"c6a1":                                "deviceCustomIPv6Address1",
		"c6a1Label":                           "deviceCustomIPv6Address1Label",
		"c6a2":                                "deviceCustomIPv6Address2",
		"c6a2Label":                           "deviceCustomIPv6Address2Label",
		"c6a3":                                "deviceCustomIPv6Address3",
		"c6a3Label":                           "deviceCustomIPv6Address3Label",
		"c6a4":                                "deviceCustomIPv6Address4",
		"c6a4Label":                           "deviceCustomIPv6Address4Label",
		"cfp1":                                "deviceCustomFloatingPoint1",
		"cfp1Label":                           "deviceCustomFloatingPoint1Label",
		"cfp2":                                "deviceCustomFloatingPoint2",
		"cfp2Label":                           "deviceCustomFloatingPoint2Label",
		"cfp3":                                "deviceCustomFloatingPoint3",
		"cfp3Label":                           "deviceCustomFloatingPoint3Label",
		"cfp4":                                "deviceCustomFloatingPoint4",
		"cfp4Label":                           "deviceCustomFloatingPoint4Label",
		"cn1":                                 "deviceCustomNumber1",
		"cn1Label":                            "deviceCustomNumber1Label",
		"cn2":                                 "deviceCustomNumber2",
		"cn2Label":                            "deviceCustomNumber2Label",
		"cn3":                                 "deviceCustomNumber3",
		"cn3Label":                            "deviceCustomNumber3Label",
		"cnt":                                 "baseEventCount",
		"cs1":                                 "deviceCustomString1",
		"cs1Label":                            "deviceCustomString1Label",
		"cs2":                                 "deviceCustomString2",
		"cs2Label":                            "deviceCustomString2Label",
		"cs3":                                 "deviceCustomString3",
		"cs3Label":                            "deviceCustomString3Label",
		"cs4":                                 "deviceCustomString4",
		"cs4Label":                            "deviceCustomString4Label",
		"cs5":                                 "deviceCustomString5",
		"cs5Label":                            "deviceCustomString5Label",
		"cs6":                                 "deviceCustomString6",
		"cs6Label":                            "deviceCustomString6Label",
		"destinationDnsDomain":                "destinationDnsDomain",
		"destinationServiceName":              "destinationServiceName",
		"destinationTranslatedAddress":        "destinationTranslatedAddress",
		"destinationTranslatedPort":           "destinationTranslatedPort",
		"deviceCustomDate1":                   "deviceCustomDate1",
		"deviceCustomDate1Label":              "deviceCustomDate1Label",
		"deviceCustomDate2":                   "deviceCustomDate2",
		"deviceCustomDate2Label":              "deviceCustomDate2Label",
		"deviceDirection":                     "deviceDirection",
		"deviceDnsDomain":                     "deviceDnsDomain",
		"deviceExternalId":                    "deviceExternalId",
		"deviceFacility":                      "deviceFacility",
		"deviceInboundInterface":              "deviceInboundInterface",
		"deviceNtDomain":                      "deviceNtDomain",
		"deviceOutboundInterface":             "deviceOutboundInterface",
		"devicePayloadId":                     "devicePayloadId",
		"deviceProcessName":                   "deviceProcessName",
		"deviceTranslatedAddress":             "deviceTranslatedAddress",
		"dhost":                               "destinationHostName",
		"dmac":                                "destinationMacAddress",
		"dntdom":                              "destinationNtDomain",
		"dpid":                                "destinationProcessId",
		"dpriv":                               "destinationUserPrivileges",
		"dproc":                               "destinationProcessName",
		"dpt":                                 "destinationPort",
		"dst":                                 "destinationAddress",
		"dtz":                                 "deviceTimeZone",
		"duid":                                "destinationUserId",
		"duser":                               "destinationUserName",
		"dvc":                                 "deviceAddress",
		"dvchost":                             "deviceHostName",
		"dvcmac":                              "deviceMacAddress",
		"dvcpid":                              "deviceProcessId",
		"end":                                 "endTime",
		"externalId":                          "externalId",
		"fileCreateTime":                      "fileCreateTime",
		"fileHash":                            "fileHash",
		"fileId":                              "fileId",
		"fileModificationTime":                "fileModificationTime",
		"filePath":                            "filePath",
		"filePermission":                      "filePermission",
		"fileType":                            "fileType",
		"flexDate1":                           "flexDate1",
		"flexDate1Label":                      "flexDate1Label",
		"flexNumber1":                         "flexNumber1",
		"flexNumber1Label":                    "flexNumber1Label",
		"flexNumber2":                         "flexNumber2",
		"flexNumber2Label":                    "flexNumber2Label",
		"flexString1":                         "flexString1",
		"flexString1Label":                    "flexString1Label",
		"flexString2":                         "flexString2",
		"flexString2Label":                    "flexString2Label",
		"fname":                               "fileName",
		"fsize":                               "fileSize",
		"in":                                  "bytesIn",
		"msg":                                 "message",
		"oldFileCreateTime":                   "oldFileCreateTime",
		"oldFileHash":                         "oldFileHash",
		"oldFileId":                           "oldFileId",
		"oldFileModificationTime":             "oldFileModificationTime",
		"oldFileName":                         "oldFileName",
		"oldFilePath":                         "oldFilePath",
		"oldFilePermission":                   "oldFilePermission",
		"oldFileSize":                         "oldFileSize",
		"oldFileType":                         "oldFileType",
		"out":                                 "bytesOut",
		"outcome":                             "eventOutcome",
		"proto":                               "transportProtocol",
		"reason":                              "reason",
		"request":                             "requestUrl",
		"requestClientApplication":            "requestClientApplication",
		"requestContext":                      "requestContext",
		"requestCookies":                      "requestCookies",
		"requestMethod":                       "requestMethod",
		"rt":                                  "deviceReceiptTime",
		"shost":                               "sourceHostName",
		"smac":                                "sourceMacAddress",
		"sntdom":                              "sourceNtDomain",
		"sourceDnsDomain":                     "sourceDnsDomain",
		"sourceServiceName":                   "sourceServiceName",
		"sourceTranslatedAddress":             "sourceTranslatedAddress",
		"sourceTranslatedPort":                "sourceTranslatedPort",
		"spid":                                "sourceProcessId",
		"spriv":                               "sourceUserPrivileges",
		"sproc":                               "sourceProcessName",
		"spt":                                 "sourcePort",
		"src":                                 "sourceAddress",
		"start":                               "startTime",
		"suid":                                "sourceUserId",
		"suser":                               "sourceUserName",
		"type":                                "type",
		"agentDnsDomain":                      "agentDnsDomain",
		"agentNtDomain":                       "agentNtDomain",
		"agentTranslatedAddress":              "agentTranslatedAddress",
		"agentTranslatedZoneExternalID":       "agentTranslatedZoneExternalID",
		"agentTranslatedZoneURI":              "agentTranslatedZoneURI",
		"agentZoneExternalID":                 "agentZoneURI",
		"agt":                                 "agentAddress",
		"ahost":                               "agentHostName",
		"aid":                                 "agentId",
		"amac":                                "agentMacAddress",
		"art":                                 "agentReceiptTime",
		"at":                                  "agentType",
		"atz":                                 "agentTimeZone",
		"av":                                  "agentVersion",
		"cat":                                 "deviceEventCategory",
		"customerExternalID":                  "customerExternalID",
		"customerURI":                         "customerURI",
		"destinationTranslatedZoneExternalID": "destinationTranslatedZoneExternalID",
		"destinationTranslatedZoneURI":        "destinationTranslatedZoneURI",
		"destinationZoneExternalID":           "destinationZoneExternalID",
		"destinationZoneURI":                  "destinationZoneURI",
		"deviceTranslatedZoneExternalID":      "deviceTranslatedZoneExternalID",
		"deviceTranslatedZoneURI":             "deviceTranslatedZoneURI",
		"deviceZoneExternalID":                "deviceZoneExternalID",
		"deviceZoneURI":                       "deviceZoneURI",
		"dlat":                                "destinationGeoLatitude",
		"dlong":                               "destinationGeoLongitude",
		"eventId":                             "eventId",
		"rawEvent":                            "rawEvent",
		"slat":                                "sourceGeoLatitude",
		"slong":                               "sourceGeoLongitude",
		"sourceTranslatedZoneExternalID":      "sourceTranslatedZoneExternalID",
		"sourceTranslatedZoneURI":             "sourceTranslatedZoneURI",
		"sourceZoneExternalID":                "sourceZoneExternalID",
		"sourceZoneURI":                       "sourceZoneURI",
	}
)

type Parser struct {
	name string
}

func (r *Parser) ID() string {
	return r.name
}

// Sample:
// CEF:0|security|threatmanager|1.0|100|detected a \| in message|10|src=10.0.0.1 act=blocked a | dst=1.1.1.1
func (r *Parser) Parse(data []byte) (map[string]string, bool) {
	if len(data) < len(entrySequence) {
		return nil, false
	}

	//
	// Match ID: "CEF:0"
	//

	pos := 0
	for ; pos < len(entrySequence); pos++ {
		if data[pos] != entrySequence[pos] {
			return nil, false
		}
	}

	//
	// Match header: "Device Vendor|Device Product|Device Version|Device Event Class ID|Name|Severity|"
	//

	out := make(map[string]string)
	valueStart := pos
	headerFieldsIdx := 0
	headerOK := false
	for ; pos < len(data); pos++ {
		if data[pos] == pipe && data[pos-1] != bs {
			out[headerFields[headerFieldsIdx]] = helpers.BytesToString(data[valueStart:pos])

			headerFieldsIdx++
			if headerFieldsIdx == len(headerFields) {
				headerOK = true
				break
			}

			valueStart = pos + 1
		}
	}

	if !headerOK {
		return nil, false
	}

	//
	// Match ext: src=10.0.0.1 act=blocked message dst=1.1.1.1
	//

	var key string
	prevSpacePos := pos
	valueStart = -1
	for ; pos < len(data); pos++ {
		if data[pos] == space && data[pos-1] != bs {
			prevSpacePos = pos
			continue
		}

		if data[pos] == eq && data[pos-1] != bs {
			if valueStart != -1 {
				out[key] = helpers.BytesToString(data[valueStart:prevSpacePos])
			}
			key = dict[helpers.BytesToString(data[prevSpacePos+1:pos])]
			valueStart = pos + 1
			continue
		}
	}

	if valueStart != -1 {
		out[key] = helpers.BytesToString(data[valueStart:pos])
	}

	return out, len(out) > 0
}
