package rfc3164

import (
   "strconv"
)

%%{
    machine syslog;
    write data;
}%%

func Parse(data []byte) (map[string][]byte, bool) {
    var cs, p, pe, eof, valueOffset, priNum, facilityNum int
    var priErr error
	pe = len(data)
	success := true
	output := make(map[string][]byte)

    %%{
        action mark {
            valueOffset = p
        }

        action setPriority {
            priNum, priErr = strconv.Atoi(string(data[valueOffset:p]))
            if priErr == nil {
                facilityNum = priNum / 8
                output["facility"] = []byte(strconv.Itoa(facilityNum))
                output["severity"] = []byte(strconv.Itoa(priNum - (facilityNum * 8)))
            }
        }

        action setTimestamp {
            output["time"] = data[valueOffset:p]
        }

        action setHostname {
            output["host"] = data[valueOffset:p]
        }

        action setAppName {
            output["app"] = data[valueOffset:p]
        }

        action setProcID {
            output["pid"] = data[valueOffset:p]
        }

        action setMsgID {
            output["mid"] = data[valueOffset:p]
        }

        action setMsg {
            output["msg"] = data[valueOffset:pe]
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

