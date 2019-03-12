package cef
import "fmt"

%%{
    machine cef;
    write data;
}%%

/*var dict = map[string]string{
    "act": "deviceAction",
    "app": "applicationProtocol",
    "c6a1": "deviceCustomIPv6Address1",
    "c6a1Label": "deviceCustomIPv6Address1",
    "c6a2": "deviceCustomIPv6Address2",
    "c6a2Label":
}*/

type Parser struct{
    name string
}

func (r *Parser) ID() string {
	return r.name
}

func (r *Parser) Parse(data []byte) (map[string]string, bool) {
    var cs, p, eof, valueOffset int
    pe := len(data)
    var recentExtKey string
	success := true
	output := make(map[string]string)

    %%{
        action mark {
            valueOffset = p
        }

        action setDeviceVendor {
            output["deviceVendor"] = string(data[valueOffset:p])
        }

        action setDeviceProduct{
            output["deviceProduct"] = string(data[valueOffset:p])
        }

        action setDeviceVersion{
            output["deviceVersion"] = string(data[valueOffset:p])
        }

        action setDeviceEventClassID{
            output["deviceEventClassID"] = string(data[valueOffset:p])
        }

        action setName{
            output["name"] = string(data[valueOffset:p])
        }

        action setSeverity {
            output["severity"] = string(data[valueOffset:p])
        }

        action setRecentExtKey {
            recentExtKey = string(data[valueOffset:p])
            fmt.Printf("key: %s\n", recentExtKey)
        }

        action setExtVal {
            output[recentExtKey] = string(data[valueOffset:p])
            fmt.Printf("val: %s\n", string(data[valueOffset:p]))
        }

        action test_fn {
            fmt.Println("test")
        }

        action fail {
            success = false;
            fbreak;
        }

        include CEF0 "cef.rl";

        write init;
        write exec;
    }%%

    return output, success
}
