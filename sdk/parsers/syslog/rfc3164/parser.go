
//line parser.rl:1
package rfc3164

import (
    "github.com/tephrocactus/raccoon-siem/sdk/parsers"
    "strconv"
)


//line parser.go:12
const syslog_start int = 0
const syslog_first_final int = 1
const syslog_error int = -1

const syslog_en_main int = 0


//line parser.rl:11


type Config struct {
    parsers.BaseConfig
}

type parser struct{
    cfg Config
}

func (r *parser) ID() string {
	return r.cfg.Name
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

    
//line parser.go:47
	{
	cs = syslog_start
	}

//line parser.go:52
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
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
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	}
	goto st_out
	st_case_0:
		switch data[p] {
		case 60:
			goto st112
		case 65:
			goto tr3
		case 68:
			goto tr4
		case 70:
			goto tr5
		case 74:
			goto tr6
		case 77:
			goto tr7
		case 78:
			goto tr8
		case 79:
			goto tr9
		case 83:
			goto tr10
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st1
tr37:
//line parser.rl:37

            valueOffset = p
        
//line parser.rl:70

            output["msg"] = string(data[valueOffset:pe])
        
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line parser.go:338
		goto st1
tr1:
//line parser.rl:37

            valueOffset = p
        
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parser.go:351
		if data[p] == 58 {
			goto st47
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st1
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 58 {
			goto st47
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto st1
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 58 {
			goto st47
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto st1
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 45:
			goto st6
		case 58:
			goto st47
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st46
		}
		goto st1
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if 48 <= data[p] && data[p] <= 57 {
			goto st7
		}
		goto st1
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if 48 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto st1
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 45 {
			goto st9
		}
		goto st1
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if 48 <= data[p] && data[p] <= 51 {
			goto st10
		}
		goto st1
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= data[p] && data[p] <= 57 {
			goto st11
		}
		goto st1
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 32:
			goto st12
		case 84:
			goto st12
		case 116:
			goto st12
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st12
		}
		goto st1
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 50 {
			goto st45
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st13
		}
		goto st1
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if 48 <= data[p] && data[p] <= 57 {
			goto st14
		}
		goto st1
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 58 {
			goto st15
		}
		goto st1
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= data[p] && data[p] <= 53 {
			goto st16
		}
		goto st1
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= data[p] && data[p] <= 57 {
			goto st17
		}
		goto st1
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 58 {
			goto st18
		}
		goto st1
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if 48 <= data[p] && data[p] <= 53 {
			goto st19
		}
		goto st1
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if 48 <= data[p] && data[p] <= 57 {
			goto st20
		}
		goto st1
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 32:
			goto tr32
		case 43:
			goto st35
		case 45:
			goto st35
		case 46:
			goto st42
		case 58:
			goto tr35
		case 90:
			goto st44
		case 122:
			goto st44
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr32
		}
		goto st1
tr32:
//line parser.rl:50

            output["time"] = string(data[valueOffset:p])
        
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:571
		switch data[p] {
		case 58:
			goto tr40
		case 95:
			goto tr38
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr38
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr39
				}
			case data[p] >= 65:
				goto tr39
			}
		default:
			goto tr39
		}
		goto tr37
tr38:
//line parser.rl:37

            valueOffset = p
        
//line parser.rl:70

            output["msg"] = string(data[valueOffset:pe])
        
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line parser.go:611
		switch data[p] {
		case 58:
			goto st34
		case 95:
			goto st22
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st22
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st23
				}
			case data[p] >= 65:
				goto st23
			}
		default:
			goto st23
		}
		goto st1
tr39:
//line parser.rl:37

            valueOffset = p
        
//line parser.rl:70

            output["msg"] = string(data[valueOffset:pe])
        
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parser.go:651
		switch data[p] {
		case 32:
			goto tr44
		case 58:
			goto tr45
		case 95:
			goto st22
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto st22
				}
			case data[p] >= 9:
				goto tr44
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st23
				}
			case data[p] >= 65:
				goto st23
			}
		default:
			goto st23
		}
		goto st1
tr44:
//line parser.rl:54

            output["host"] = string(data[valueOffset:p])
        
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line parser.go:694
		switch data[p] {
		case 32:
			goto tr47
		case 91:
			goto st1
		case 93:
			goto st1
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr47
		}
		goto tr46
tr46:
//line parser.rl:37

            valueOffset = p
        
//line parser.rl:70

            output["msg"] = string(data[valueOffset:pe])
        
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line parser.go:722
		switch data[p] {
		case 32:
			goto st1
		case 58:
			goto tr49
		case 91:
			goto tr50
		case 93:
			goto st1
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st1
		}
		goto st25
tr49:
//line parser.rl:58

            output["app"] = string(data[valueOffset:p])
        
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line parser.go:748
		switch data[p] {
		case 32:
			goto st27
		case 58:
			goto tr49
		case 91:
			goto tr50
		case 93:
			goto st1
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st27
		}
		goto st25
tr47:
//line parser.rl:37

            valueOffset = p
        
//line parser.rl:70

            output["msg"] = string(data[valueOffset:pe])
        
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line parser.go:778
		goto tr37
tr50:
//line parser.rl:58

            output["app"] = string(data[valueOffset:p])
        
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line parser.go:791
		if 48 <= data[p] && data[p] <= 57 {
			goto tr52
		}
		goto st1
tr52:
//line parser.rl:37

            valueOffset = p
        
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line parser.go:807
		if data[p] == 93 {
			goto tr54
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st29
		}
		goto st1
tr54:
//line parser.rl:62

            output["pid"] = string(data[valueOffset:p])
        
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line parser.go:826
		if data[p] == 58 {
			goto st31
		}
		goto st1
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 32 {
			goto st27
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st27
		}
		goto st1
tr45:
//line parser.rl:54

            output["host"] = string(data[valueOffset:p])
        
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line parser.go:854
		switch data[p] {
		case 32:
			goto st24
		case 58:
			goto st33
		case 95:
			goto st22
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto st22
				}
			case data[p] >= 9:
				goto st24
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st23
				}
			case data[p] >= 65:
				goto st23
			}
		default:
			goto st23
		}
		goto st1
tr58:
//line parser.rl:54

            output["host"] = string(data[valueOffset:p])
        
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line parser.go:897
		switch data[p] {
		case 32:
			goto tr44
		case 58:
			goto tr58
		case 95:
			goto st22
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto st22
				}
			case data[p] >= 9:
				goto tr44
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st23
				}
			case data[p] >= 65:
				goto st23
			}
		default:
			goto st23
		}
		goto st1
tr40:
//line parser.rl:37

            valueOffset = p
        
//line parser.rl:70

            output["msg"] = string(data[valueOffset:pe])
        
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line parser.go:944
		switch data[p] {
		case 58:
			goto st33
		case 95:
			goto st22
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st22
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st23
				}
			case data[p] >= 65:
				goto st23
			}
		default:
			goto st23
		}
		goto st1
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if 48 <= data[p] && data[p] <= 57 {
			goto st36
		}
		goto st1
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if 48 <= data[p] && data[p] <= 57 {
			goto st37
		}
		goto st1
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 32:
			goto tr32
		case 58:
			goto tr62
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st38
			}
		case data[p] >= 9:
			goto tr32
		}
		goto st1
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if 48 <= data[p] && data[p] <= 57 {
			goto st39
		}
		goto st1
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 32:
			goto tr32
		case 58:
			goto tr35
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr32
		}
		goto st1
tr35:
//line parser.rl:50

            output["time"] = string(data[valueOffset:p])
        
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line parser.go:1042
		if data[p] == 32 {
			goto st21
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st21
		}
		goto st1
tr62:
//line parser.rl:50

            output["time"] = string(data[valueOffset:p])
        
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line parser.go:1061
		if data[p] == 32 {
			goto st21
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st38
			}
		case data[p] >= 9:
			goto st21
		}
		goto st1
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if 48 <= data[p] && data[p] <= 57 {
			goto st43
		}
		goto st1
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 32:
			goto tr32
		case 43:
			goto st35
		case 45:
			goto st35
		case 58:
			goto tr35
		case 90:
			goto st44
		case 122:
			goto st44
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st43
			}
		case data[p] >= 9:
			goto tr32
		}
		goto st1
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 32:
			goto tr32
		case 43:
			goto st35
		case 45:
			goto st35
		case 58:
			goto tr35
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr32
		}
		goto st1
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if 48 <= data[p] && data[p] <= 51 {
			goto st14
		}
		goto st1
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 58 {
			goto st47
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st46
		}
		goto st1
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 32 {
			goto st48
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st48
		}
		goto st1
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 65:
			goto tr3
		case 68:
			goto tr4
		case 70:
			goto tr5
		case 74:
			goto tr6
		case 77:
			goto tr7
		case 78:
			goto tr8
		case 79:
			goto tr9
		case 83:
			goto tr10
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr67
		}
		goto st1
tr67:
//line parser.rl:37

            valueOffset = p
        
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line parser.go:1201
		if 48 <= data[p] && data[p] <= 57 {
			goto st50
		}
		goto st1
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		if 48 <= data[p] && data[p] <= 57 {
			goto st51
		}
		goto st1
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		if 48 <= data[p] && data[p] <= 57 {
			goto st52
		}
		goto st1
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 45 {
			goto st6
		}
		goto st1
tr3:
//line parser.rl:37

            valueOffset = p
        
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line parser.go:1244
		switch data[p] {
		case 112:
			goto st54
		case 117:
			goto st75
		}
		goto st1
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if data[p] == 114 {
			goto st55
		}
		goto st1
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 32:
			goto st56
		case 105:
			goto st73
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 32:
			goto st57
		case 51:
			goto st72
		}
		switch {
		case data[p] < 49:
			if 9 <= data[p] && data[p] <= 13 {
				goto st57
			}
		case data[p] > 50:
			if 52 <= data[p] && data[p] <= 57 {
				goto st58
			}
		default:
			goto st71
		}
		goto st1
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if 49 <= data[p] && data[p] <= 57 {
			goto st58
		}
		goto st1
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 32 {
			goto st59
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st59
		}
		goto st1
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 50 {
			goto st70
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st60
		}
		goto st1
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if 48 <= data[p] && data[p] <= 57 {
			goto st61
		}
		goto st1
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 58 {
			goto st62
		}
		goto st1
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if 48 <= data[p] && data[p] <= 53 {
			goto st63
		}
		goto st1
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if 48 <= data[p] && data[p] <= 57 {
			goto st64
		}
		goto st1
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 58 {
			goto st65
		}
		goto st1
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if 48 <= data[p] && data[p] <= 53 {
			goto st66
		}
		goto st1
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		if 48 <= data[p] && data[p] <= 57 {
			goto st67
		}
		goto st1
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 32:
			goto tr32
		case 46:
			goto st68
		case 58:
			goto tr35
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr32
		}
		goto st1
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if 48 <= data[p] && data[p] <= 57 {
			goto st69
		}
		goto st1
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 32:
			goto tr32
		case 58:
			goto tr35
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st69
			}
		case data[p] >= 9:
			goto tr32
		}
		goto st1
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		if 48 <= data[p] && data[p] <= 51 {
			goto st61
		}
		goto st1
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		if data[p] == 32 {
			goto st59
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st58
			}
		case data[p] >= 9:
			goto st59
		}
		goto st1
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 32 {
			goto st59
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 49 {
				goto st58
			}
		case data[p] >= 9:
			goto st59
		}
		goto st1
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 108 {
			goto st74
		}
		goto st1
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		if data[p] == 32 {
			goto st56
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if data[p] == 103 {
			goto st76
		}
		goto st1
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 32:
			goto st56
		case 117:
			goto st77
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		if data[p] == 115 {
			goto st78
		}
		goto st1
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		if data[p] == 116 {
			goto st74
		}
		goto st1
tr4:
//line parser.rl:37

            valueOffset = p
        
	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line parser.go:1559
		if data[p] == 101 {
			goto st80
		}
		goto st1
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		if data[p] == 99 {
			goto st81
		}
		goto st1
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 32:
			goto st56
		case 101:
			goto st82
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		if data[p] == 109 {
			goto st83
		}
		goto st1
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		if data[p] == 98 {
			goto st84
		}
		goto st1
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 101 {
			goto st85
		}
		goto st1
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		if data[p] == 114 {
			goto st74
		}
		goto st1
tr5:
//line parser.rl:37

            valueOffset = p
        
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line parser.go:1635
		if data[p] == 101 {
			goto st87
		}
		goto st1
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		if data[p] == 98 {
			goto st88
		}
		goto st1
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 32:
			goto st56
		case 114:
			goto st89
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		if data[p] == 117 {
			goto st90
		}
		goto st1
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		if data[p] == 97 {
			goto st91
		}
		goto st1
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		if data[p] == 114 {
			goto st92
		}
		goto st1
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		if data[p] == 121 {
			goto st74
		}
		goto st1
tr6:
//line parser.rl:37

            valueOffset = p
        
	goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line parser.go:1711
		switch data[p] {
		case 97:
			goto st94
		case 117:
			goto st96
		}
		goto st1
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		if data[p] == 110 {
			goto st95
		}
		goto st1
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 32:
			goto st56
		case 117:
			goto st90
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 108:
			goto st97
		case 110:
			goto st98
		}
		goto st1
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 32:
			goto st56
		case 121:
			goto st74
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 32:
			goto st56
		case 101:
			goto st74
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
tr7:
//line parser.rl:37

            valueOffset = p
        
	goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line parser.go:1796
		if data[p] == 97 {
			goto st100
		}
		goto st1
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 32:
			goto st56
		case 114:
			goto st101
		case 121:
			goto st74
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 32:
			goto st56
		case 99:
			goto st102
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		if data[p] == 104 {
			goto st74
		}
		goto st1
tr8:
//line parser.rl:37

            valueOffset = p
        
	goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
//line parser.go:1853
		if data[p] == 111 {
			goto st104
		}
		goto st1
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		if data[p] == 118 {
			goto st81
		}
		goto st1
tr9:
//line parser.rl:37

            valueOffset = p
        
	goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line parser.go:1878
		if data[p] == 99 {
			goto st106
		}
		goto st1
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		if data[p] == 116 {
			goto st107
		}
		goto st1
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 32:
			goto st56
		case 111:
			goto st83
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
tr10:
//line parser.rl:37

            valueOffset = p
        
	goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line parser.go:1918
		if data[p] == 101 {
			goto st109
		}
		goto st1
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		if data[p] == 112 {
			goto st110
		}
		goto st1
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 32:
			goto st56
		case 116:
			goto st111
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st1
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		if data[p] == 101 {
			goto st82
		}
		goto st1
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr122
		}
		goto st1
tr122:
//line parser.rl:37

            valueOffset = p
        
	goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line parser.go:1976
		if data[p] == 62 {
			goto tr124
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st114
		}
		goto st1
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		if data[p] == 62 {
			goto tr124
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st115
		}
		goto st1
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		if data[p] == 62 {
			goto tr124
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st116
		}
		goto st1
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		if data[p] == 62 {
			goto tr124
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st117
		}
		goto st1
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		if data[p] == 62 {
			goto tr124
		}
		goto st1
tr124:
//line parser.rl:41

            priNum, priErr = strconv.Atoi(string(data[valueOffset:p]))
            if priErr == nil {
                facilityNum = priNum / 8
                output["facility"] = strconv.Itoa(facilityNum)
                output["severity"] = strconv.Itoa(priNum - (facilityNum * 8))
            }
        
	goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line parser.go:2045
		switch data[p] {
		case 65:
			goto tr3
		case 68:
			goto tr4
		case 70:
			goto tr5
		case 74:
			goto tr6
		case 77:
			goto tr7
		case 78:
			goto tr8
		case 79:
			goto tr9
		case 83:
			goto tr10
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st1
	st_out:
	_test_eof1: cs = 1; goto _test_eof
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
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 0:
//line parser.rl:74

            success = false;
            {p++; cs = 0; goto _out }
        
//line parser.go:2197
		}
	}

	_out: {}
	}

//line parser.rl:83


    return output, success
}

