package rfc5424

import "strconv"

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
    var recentSDKey string
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
            output["time"] = string(data[valueOffset:p])
        }

        action setHostname {
            output["host"] = string(data[valueOffset:p])
        }

        action setAppName {
            output["app"] = string(data[valueOffset:p])
        }

        action setProcID {
            output["pid"] = string(data[valueOffset:p])
        }

        action setMsgID {
            output["mid"] = string(data[valueOffset:p])
        }

        action setMsg {
            output["msg"] = string(data[valueOffset:pe])
        }

        action setSDKey {
            recentSDKey = string(data[valueOffset:p])
        }

        action setSDValue {
            output[recentSDKey] = string(data[valueOffset:p])
        }

        action fail {
            success = false;
            fbreak;
        }

        include RFC5424 "rfc5424.rl";

        write init;
        write exec;
    }%%

    return output, success
}
