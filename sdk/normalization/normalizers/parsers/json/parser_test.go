package json

import (
	"gotest.tools/assert"
	"testing"
)

var sampleWithoutOpenCurlyBracket = []byte(`"key": 2}`)
var sampleWithoutCloseCurlyBracket = []byte(`{"key": 2`)

var sample = []byte(`{
  "first": "timestamp",

  "timestamp": "2018-05-24 23:15:07",

  "id": 0,

  "class": "connection",

  "event": "connect",

  "connection_id": 12,

  "медведи и балалайки": true,

  "account": {
    "user": "user",
    "host": "localh\"ost",
    "domain": null,
    "login": "undefined",
    "group": ""
  },

  "string_arr": [ "1", "2", "3[" ],

  "obj_arr": [ { "k": null }, { "k": "s}tr", "v": [ "a", "b", "c" ] } ],

  "login": {
    "user": "user",
    "os": "",
    "ip": "::1",
    "proxy": "",
    "org": null
  },

  "connection_data": {
    "connection_type": "tcp/ip",
    "status": 0,
    "db": "bank_db"
  },

  "level": -1
}`)

var SampleLarge = []byte(`{
  "@timestamp": "2019-03-29T10:46:08.122Z",
  "@metadata": {
    "beat": "winlogbeat",
    "type": "doc",
    "version": "6.7.0"
  },
  "version": 2,
  "activity_id": "{DB002E43-DF32-0003-622E-00DB32DFD401}",
  "computer_name": "WIN-COMPUTER",
  "thread_id": 1152,
  "level": "Information",
  "host": {
    "name": "WIN-VF9S2B8SGV6",
    "architecture": "x86_64",
    "os": {
      "platform": "windows",
      "version": "10.0",
      "family": "windows",
      "name": "Windows Server 2016 Standard",
      "build": "14393.2248"
    },
    "id": "427692e0-4b24-4f9b-805c-88cbdee78eb8"
  },
  "event_data": {
    "SubjectLogonId": "0x0",
    "TransmittedServices": "-",
    "ImpersonationLevel": "%%1833",
    "SubjectUserName": "-",
    "WorkstationName": "WORKSTATION",
    "SubjectDomainName": "-",
    "AuthenticationPackageName": "NTLM",
    "TargetUserName": "Administrator",
    "KeyLength": "128",
    "TargetOutboundUserName": "-",
    "TargetLinkedLogonId": "0x0",
    "ElevatedToken": "%%1842",
    "ProcessName": "-",
    "SubjectUserSid": "S-1-0-0",
    "TargetLogonId": "0xc4e54348",
    "LogonType": "3",
    "IpPort": "0",
    "TargetOutboundDomainName": "-",
    "RestrictedAdminMode": "-",
    "TargetDomainName": "WIN-DOMAIN",
    "TargetUserSid": "S-1-5-21-1827733825-3198728177-1094963921-500",
    "IpAddress": "10.16.32.75",
    "LogonGuid": "{00000000-0000-0000-0000-000000000000}",
    "LmPackageName": "NTLM V2",
    "VirtualAccount": "%%1843",
    "LogonProcessName": "NtLmSsp ",
    "ProcessId": "0x0"
  },
  "event_id": 4624,
  "opcode": "Info",
  "task": "Logon",
  "provider_guid": "{54849625-5478-4994-A5BA-3E3B0328C30D}",
  "process_id": 808,
  "log_name": "Security",
  "message": "An account was successfully logged on.\n\nSubject:\n\tSecurity ID:\t\tS-1-0-0\n\tAccount Name:\t\t-\n\tAccount Domain:\t\t-\n\tLogon ID:\t\t0x0\n\nLogon Information:\n\tLogon Type:\t\t3\n\tRestricted Admin Mode:\t-\n\tVirtual Account:\t\tNo\n\tElevated Token:\t\tYes\n\nImpersonation Level:\t\tImpersonation\n\nNew Logon:\n\tSecurity ID:\t\tS-1-5-21-1827733825-3198728177-1094963921-500\n\tAccount Name:\t\tAdministrator\n\tAccount Domain:\t\tWIN-VF9S2B8SGV6\n\tLogon ID:\t\t0xC4E54348\n\tLinked Logon ID:\t\t0x0\n\tNetwork Account Name:\t-\n\tNetwork Account Domain:\t-\n\tLogon GUID:\t\t{00000000-0000-0000-0000-000000000000}\n\nProcess Information:\n\tProcess ID:\t\t0x0\n\tProcess Name:\t\t-\n\nNetwork Information:\n\tWorkstation Name:\tPADALKO\n\tSource Network Address:\t10.16.32.75\n\tSource Port:\t\t0\n\nDetailed Authentication Information:\n\tLogon Process:\t\tNtLmSsp \n\tAuthentication Package:\tNTLM\n\tTransited Services:\t-\n\tPackage Name (NTLM only):\tNTLM V2\n\tKey Length:\t\t128\n\nThis event is generated when a logon session is created. It is generated on the computer that was accessed.\n\nThe subject fields indicate the account on the local system which requested the logon. This is most commonly a service such as the Server service, or a local process such as Winlogon.exe or Services.exe.\n\nThe logon type field indicates the kind of logon that occurred. The most common types are 2 (interactive) and 3 (network).\n\nThe New Logon fields indicate the account for whom the new logon was created, i.e. the account that was logged on.\n\nThe network fields indicate where a remote logon request originated. Workstation name is not always available and may be left blank in some cases.\n\nThe impersonation level field indicates the extent to which a process in the logon session can impersonate.\n\nThe authentication information fields provide detailed information about this specific logon request.\n\t- Logon GUID is a unique identifier that can be used to correlate this event with a KDC event.\n\t- Transited services indicate which intermediate services have participated in this logon request.\n\t- Package name indicates which sub-protocol was used among the NTLM protocols.\n\t- Key length indicates the length of the generated session key. This will be 0 if no session key was requested.",
  "type": "wineventlog",
  "record_number": "9610",
  "source_name": "Microsoft-Windows-Security-Auditing",
  "keywords": [
    "Audit Success"
  ],
  "beat": {
    "hostname": "WIN-VF9S2B8SGV6",
    "version": "6.7.0",
    "name": "WIN-VF9S2B8SGV6"
  }
}`)

func TestJSONParser(t *testing.T) {
	success := false
	result := make(map[string][]byte)
	callback := func(key string, value []byte) {
		result[key] = value
	}

	for i := 0; i < 5; i++ {
		success = Parse(sample, callback)
	}

	assert.Equal(t, success, true)
	assert.DeepEqual(t, result["first"], []byte("timestamp"))
	assert.DeepEqual(t, result["timestamp"], []byte("2018-05-24 23:15:07"))
	assert.DeepEqual(t, result["connection_id"], []byte("12"))
	assert.DeepEqual(t, result["account.user"], []byte("user"))
	assert.DeepEqual(t, result["login.org"], []byte("null"))
	assert.DeepEqual(t, result["connection_data.connection_type"], []byte("tcp/ip"))
	assert.DeepEqual(t, result["level"], []byte("-1"))

	success = Parse(sampleWithoutOpenCurlyBracket, callback)
	assert.Equal(t, success, false)

	success = Parse(sampleWithoutCloseCurlyBracket, callback)
	assert.Equal(t, success, false)
}

func BenchmarkJSONParser(b *testing.B) {
	cb := func(key string, value []byte) {}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Parse(SampleLarge, cb)
	}
}
