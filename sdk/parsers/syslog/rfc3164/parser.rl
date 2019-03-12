package rfc3164

import (
   "strconv"
   "github.com/tephrocactus/raccoon-siem/sdk/helpers"
)

%%{
    machine syslog;
    write data;
}%%

type Parser struct{
    name string
}

func (r *Parser) ID() string {
	return r.name
}

func (r *Parser) Parse(data []byte) (map[string]string, bool) {
    var cs, p, pe, eof, priNum, facilityNum, valueOffset int
    var priErr error
	pe = len(data)
	success := true
	output := make(map[string]string)

    %%{
        action mark {
            valueOffset = p
        }

        action setPriority {
            priNum, priErr = strconv.Atoi(string(data[valueOffset:p]))
            if priErr == nil {
                facilityNum = priNum / 8
                output["facility"] = strconv.Itoa(facilityNum)
                output["severity"] = strconv.Itoa(priNum - (facilityNum * 8))
            }
        }

        action setTimestamp {
            output["time"] = helpers.BytesToString(data[valueOffset:p])
        }

        action setHostname {
            output["host"] = helpers.BytesToString(data[valueOffset:p])
        }

        action setAppName {
            output["app"] = helpers.BytesToString(data[valueOffset:p])
        }

        action setProcID {
            output["pid"] = helpers.BytesToString(data[valueOffset:p])
        }

        action setMsgID {
            output["mid"] = helpers.BytesToString(data[valueOffset:p])
        }

        action setMsg {
            output["msg"] = helpers.BytesToString(data[valueOffset:pe])
        }

        action fail {
            success = false;
            fbreak;
        }

        include RFC3164 "rfc3164.rl";

        write init;
        write exec;
    }%%

    return output, success
}
