package context

type Context struct {
	RawInput    []byte
	ParsedInput map[string][]byte
	Event
}
