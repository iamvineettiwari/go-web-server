package http

var (
	CR      byte   = '\r'
	LF      byte   = '\n'
	CRLF    []byte = []byte{CR, LF}
	SPACE   []byte = []byte(" ")
	VERSION []byte = []byte("HTTP/1.1")
)
