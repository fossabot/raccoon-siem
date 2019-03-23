# Sample
# <34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8

# This initial variant was shamelessly taken from a brilliant Elastic Beats (https://www.elastic.co/products/beats).

%%{
  machine RFC3164;

  brackets = "[" | "]";

  priority = digit{1,5} >mark %setPriority;
  prio =  "<" priority ">";

  month = ( "Jan" ("uary")? | "Feb" "ruary"? | "Mar" "ch"? | "Apr" "il"? | "Ma" "y"? | "Jun" "e"? | "Jul" "y"? | "Aug" "ust"? | "Sep" ("tember")? | "Oct" "ober"? | "Nov" "ember"? | "Dec" "ember"?);

  multiple_digits_day = (([12][0-9]) | ("3"[01]));
  single_digit_day = [1-9];
  day = (space? single_digit_day | multiple_digits_day);

  hour = ([01][0-9]|"2"[0-3]);
  minute = ([0-5][0-9]);
  second = ([0-5][0-9]);
  nanosecond = digit+;
  time = hour ":" minute ":" second ("." nanosecond)?;
  offset_marker = "Z" | "z";
  offset_direction = "-" | "+";
  offset_hour = digit{2};
  offset_minute = digit{2};
  timezone = (offset_marker | offset_marker? offset_direction offset_hour (":"? offset_minute)?);

  year = digit{4};
  month_numeric = digit{2};
  day_two_digits = ([0-3][0-9]);

  timestamp_rfc3164 = month space day space time;
  time_separator = "T" | "t";
  timestamp_rfc3339 = year "-" month_numeric "-" day_two_digits (time_separator | space) time timezone?;
  timestamp = (timestamp_rfc3339 | timestamp_rfc3164) >mark %setTimestamp $err(fail) ":"?;

  hostname = ([a-zA-Z0-9\.\-_:]*([a-zA-Z0-9] | "::"))+ >mark %setHostname $err(fail);
  hostVars = (hostname ":") | hostname;
  SPACE_BEFORE_MSG = space %mark %setMsg;
  header = timestamp space hostVars ":"? SPACE_BEFORE_MSG;

  program = (extend -space -brackets)+ >mark %setAppName $err(fail);
  pid = digit+ >mark %setProcID;
  syslogprog = program ("[" pid "]")? ":";
  message = any+;
  msg = syslogprog? SPACE_BEFORE_MSG message;
  sequence = digit+ ":" space;

  main := (prio)?(sequence)? (header msg | timestamp SPACE_BEFORE_MSG message | message);
}%%
