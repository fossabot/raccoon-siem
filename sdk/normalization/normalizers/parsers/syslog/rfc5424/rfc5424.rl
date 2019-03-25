# Format
# https://tools.ietf.org/html/rfc5424#section-6

# Sample
# <165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] An application event log entry...

%%{
    machine RFC5424;

    NIL = '-';
    PRINT_US_ASCII = '!'..'~';

    ### PRIORITY ###

    PRIVAL = digit{1,3} >mark %setPriority;
    PRI = '<' PRIVAL '>' $err(fail);

    ### VERSION ###

    VERSION = digit{0,2} $err(fail);

    ### TIMESTAMP ###

    DATE_FULLYEAR = digit{4};
    DATE_MONTH = digit{2};
    DATE_MDAY = digit{2};
    FULL_DATE = DATE_FULLYEAR '-' DATE_MONTH '-' DATE_MDAY;
    TIME_HOUR = digit{2};
    TIME_MINUTE = digit{2};
    TIME_SECOND = digit{2};
    TIME_SECFRAC = '.' digit{1,6};
    TIME_NUMOFFSET = ('+' | '-') TIME_HOUR ':' TIME_MINUTE;
    TIME_OFFSET = ('Z' | TIME_NUMOFFSET);
    PARTIAL_TIME = TIME_HOUR ':' TIME_MINUTE ':' TIME_SECOND TIME_SECFRAC?;
    FULL_TIME = PARTIAL_TIME TIME_OFFSET;
    TIMESTAMP = (NIL | FULL_DATE 'T' FULL_TIME) >mark %setTimestamp $err(fail);

    ### Hostname ###

    HOSTNAME = (NIL | PRINT_US_ASCII{1,255}) >mark %setHostname $err(fail);

    ### APPNAME ###

    APPNAME = (NIL | PRINT_US_ASCII{1,48}) >mark %setAppName $err(fail);

    ### PROCID ###

    PROCID = (NIL | PRINT_US_ASCII{1,128}) >mark %setProcID $err(fail);

    ### MSGID ###

    MSGID = (NIL | PRINT_US_ASCII{1,32}) >mark %setMsgID $err(fail);

    ### STRUCTURED DATA ###
    SD_KEY = (PRINT_US_ASCII - ('=' | space | ']' | '"')){1,32};
    SD_VAL = '"' %mark any+ '"' >setSDValue;
    SD_ID = SD_KEY;
    SD_KV = SD_KEY >mark %setSDKey '=' SD_VAL;
    SD_CONTENT = '[' SD_ID space (SD_KV space?)+ ']';
    SD = (NIL | SD_CONTENT);

    ### MSG ###

    SPACE_BEFORE_MSG = space %mark %setMsg;
    MSG = any+;

    ### MAIN MACHINE ###
    main := PRI VERSION space TIMESTAMP space HOSTNAME space APPNAME space PROCID space MSGID (space SD)? SPACE_BEFORE_MSG MSG;
}%%
