package rfc3164

import (
    "github.com/tephrocactus/raccoon-siem/sdk/parsers"
    "strconv"
)

%%{
    machine syslog;
    write data;
}%%

type Config struct {
    parsers.BaseConfig
}

type parser struct{
    cfg Config
}

func NewParser(cfg Config) (*parser, error) {
    return &parser{cfg:cfg}, nil
}

func (r *parser) Parse(data []byte) (map[string]string, bool) {
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

