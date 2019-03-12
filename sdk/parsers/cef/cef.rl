# Format
# CEF:Version|Device Vendor|Device Product|Device Version|Device Event Class ID|Name|Severity|[Extension]

# The entire message should be UTF-8 encoded.
# Spaces used in the header are valid. Do not encode a space character by using <space>.

# If a pipe (|) is used in the header, it hasto be escaped with a backslash (\). But note that pipesin the
# extension do not need escaping. For example:

# CEF:0|security|threatmanager|1.0|100|detected a \| in message|10|src=10.0.0.1 act=blocked a | dst=1.1.1.1

# If a backslash (\) is used in the header or the extension, it hasto be escaped with another backslash
# (\). For example:

# CEF:0|security|threatmanager|1.0|100|detected a \\ in packet|10|src=10.0.0.1 act=blocked a \\ dst=1.1.1.1

# If an equal sign (=) is used in the extensions, it hasto be escaped with a backslash (\). Equalsignsin
# the header need no escaping. For example:

# CEF:0|security|threatmanager|1.0|100|detected a = in message|10|src=10.0.0.1 act=blocked a \= dst=1.1.1.1

# Multi-line fields can be sent by CEF by encoding the newline character as\n or\r. Note that multiple
# lines are only allowed in the value part of the extensions. For example:

# CEF:0|security|threatmanager|1.0|100|Detected a threat. No action needed.|10|src=10.0.0.1 msg=Detected a threat.\n No action needed.

%%{
    machine CEF0;

    header_chars = [^|]+;
    ext_chars = [^= ]+;

    ID = 'CEF:' $err(fail);
    VERSION = digit{1} $err(fail);
    DEVICE_VENDOR = header_chars >mark %setDeviceVendor $err(fail);
    DEVICE_PRODUCT = header_chars >mark %setDeviceProduct $err(fail);
    DEVICE_VERSION = header_chars >mark %setDeviceVersion $err(fail);
    DEVICE_EVENT_CLASS_ID = header_chars >mark %setDeviceEventClassID $err(fail);
    NAME = header_chars >mark %setName $err(fail);
    SEVERITY = header_chars >mark %setSeverity $err(fail);
    EXT_KEY = ext_chars >mark %setRecentExtKey;
    EXT_VAL = ext_chars >mark %setExtVal;
    EXT_KV = EXT_KEY '=' EXT_VAL %eof(test_fn) space;

    main := ID VERSION '|' DEVICE_VENDOR '|' DEVICE_PRODUCT '|'
        DEVICE_VERSION '|' DEVICE_EVENT_CLASS_ID '|' NAME '|' SEVERITY '|' EXT_KV*;
}%%
