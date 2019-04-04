package rfc5424

import (
    "strconv"
    "github.com/tephrocactus/raccoon-siem/sdk/normalization/normalizers/parsers"
)

%%{
    machine syslog;
    write data;
}%%

func Parse(data []byte, callback parsers.Callback) bool {
    var cs, p, pe, eof, priNum, facilityNum, valueOffset int
    var recentSDKey string
    var priErr error
	pe = len(data)
	success := true

    %%{
        action mark {
            valueOffset = p
        }

         action setPriority {
             priNum, priErr = strconv.Atoi(string(data[valueOffset:p]))
             if priErr == nil {
                 facilityNum = priNum / 8
                 callback("facility", []byte(strconv.Itoa(facilityNum)))
                 callback("severity", []byte(strconv.Itoa(priNum - (facilityNum * 8))))
             }
         }

        action setTimestamp {
            callback("time", data[valueOffset:p])
        }

        action setHostname {
            callback("host", data[valueOffset:p])
        }

        action setAppName {
            callback("app", data[valueOffset:p])
        }

        action setProcID {
            callback("pid", data[valueOffset:p])
        }

        action setMsgID {
            callback("mid", data[valueOffset:p])
        }

        action setMsg {
            callback("msg", data[valueOffset:pe])
        }

        action setSDKey {
            recentSDKey = string(data[valueOffset:p])
        }

        action setSDValue {
            callback(recentSDKey, data[valueOffset:p])
        }

        action fail {
            success = false;
            fbreak;
        }

        include RFC5424 "rfc5424.rl";

        write init;
        write exec;
    }%%

    return success
}
