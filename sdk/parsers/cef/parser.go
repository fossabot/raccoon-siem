
//line parser.rl:1
package cef
import "fmt"


//line parser.go:8
const cef_start int = 1
const cef_first_final int = 23
const cef_error int = 0

const cef_en_main int = 1


//line parser.rl:7


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

    
//line parser.go:44
	{
	cs = cef_start
	}

//line parser.go:49
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 23:
		goto st_case_23
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 24:
		goto st_case_24
	case 22:
		goto st_case_22
	case 25:
		goto st_case_25
	}
	goto st_out
	st_case_1:
		if data[p] == 67 {
			goto st2
		}
		goto tr0
tr0:
//line parser.rl:76

            success = false;
            {p++; cs = 0; goto _out }
        
	goto st0
//line parser.go:121
st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 69 {
			goto st3
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 70 {
			goto st4
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 58 {
			goto st5
		}
		goto tr0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= data[p] && data[p] <= 57 {
			goto st6
		}
		goto tr0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 124 {
			goto st7
		}
		goto tr0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 124 {
			goto tr0
		}
		goto tr7
tr7:
//line parser.rl:34

            valueOffset = p
        
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parser.go:191
		if data[p] == 124 {
			goto tr9
		}
		goto st8
tr9:
//line parser.rl:38

            output["deviceVendor"] = string(data[valueOffset:p])
        
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line parser.go:207
		if data[p] == 124 {
			goto tr0
		}
		goto tr10
tr10:
//line parser.rl:34

            valueOffset = p
        
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line parser.go:223
		if data[p] == 124 {
			goto tr12
		}
		goto st10
tr12:
//line parser.rl:42

            output["deviceProduct"] = string(data[valueOffset:p])
        
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line parser.go:239
		if data[p] == 124 {
			goto tr0
		}
		goto tr13
tr13:
//line parser.rl:34

            valueOffset = p
        
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parser.go:255
		if data[p] == 124 {
			goto tr15
		}
		goto st12
tr15:
//line parser.rl:46

            output["deviceVersion"] = string(data[valueOffset:p])
        
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parser.go:271
		if data[p] == 124 {
			goto tr0
		}
		goto tr16
tr16:
//line parser.rl:34

            valueOffset = p
        
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parser.go:287
		if data[p] == 124 {
			goto tr18
		}
		goto st14
tr18:
//line parser.rl:50

            output["deviceEventClassID"] = string(data[valueOffset:p])
        
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parser.go:303
		if data[p] == 124 {
			goto tr0
		}
		goto tr19
tr19:
//line parser.rl:34

            valueOffset = p
        
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line parser.go:319
		if data[p] == 124 {
			goto tr21
		}
		goto st16
tr21:
//line parser.rl:54

            output["name"] = string(data[valueOffset:p])
        
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parser.go:335
		if data[p] == 124 {
			goto tr0
		}
		goto tr22
tr22:
//line parser.rl:34

            valueOffset = p
        
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line parser.go:351
		if data[p] == 124 {
			goto tr24
		}
		goto st18
tr24:
//line parser.rl:58

            output["severity"] = string(data[valueOffset:p])
        
	goto st23
tr31:
//line parser.rl:67

            output[recentExtKey] = string(data[valueOffset:p])
            fmt.Printf("val: %s\n", string(data[valueOffset:p]))
        
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parser.go:374
		switch data[p] {
		case 32:
			goto st0
		case 61:
			goto st0
		}
		goto tr34
tr34:
//line parser.rl:34

            valueOffset = p
        
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parser.go:393
		switch data[p] {
		case 32:
			goto st0
		case 61:
			goto tr27
		}
		goto st19
tr27:
//line parser.rl:62

            recentExtKey = string(data[valueOffset:p])
            fmt.Printf("key: %s\n", recentExtKey)
        
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line parser.go:413
		switch data[p] {
		case 32:
			goto st0
		case 61:
			goto st0
		}
		goto tr28
tr28:
//line parser.rl:34

            valueOffset = p
        
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:432
		switch data[p] {
		case 32:
			goto tr31
		case 61:
			goto st0
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr30
		}
		goto st21
tr30:
//line parser.rl:67

            output[recentExtKey] = string(data[valueOffset:p])
            fmt.Printf("val: %s\n", string(data[valueOffset:p]))
        
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line parser.go:455
		switch data[p] {
		case 32:
			goto tr31
		case 61:
			goto st0
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr36
		}
		goto tr35
tr35:
//line parser.rl:34

            valueOffset = p
        
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line parser.go:477
		switch data[p] {
		case 32:
			goto tr31
		case 61:
			goto tr27
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr33
		}
		goto st22
tr33:
//line parser.rl:67

            output[recentExtKey] = string(data[valueOffset:p])
            fmt.Printf("val: %s\n", string(data[valueOffset:p]))
        
	goto st25
tr36:
//line parser.rl:67

            output[recentExtKey] = string(data[valueOffset:p])
            fmt.Printf("val: %s\n", string(data[valueOffset:p]))
        
//line parser.rl:34

            valueOffset = p
        
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line parser.go:511
		switch data[p] {
		case 32:
			goto tr31
		case 61:
			goto tr27
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr36
		}
		goto tr35
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 21, 22, 24, 25:
//line parser.rl:72

            fmt.Println("test")
        
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
//line parser.rl:76

            success = false;
            {p++; cs = 0; goto _out }
        
//line parser.go:562
		}
	}

	_out: {}
	}

//line parser.rl:85


    return output, success
}
