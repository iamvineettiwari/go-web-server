package http

var (
	CR      byte   = '\r'
	LF      byte   = '\n'
	CRLF    []byte = []byte{CR, LF}
	SPACE   []byte = []byte(" ")
	VERSION []byte = []byte("HTTP/1.1")
)

const (
	APPLICATION_JSON       string = "application/json"
	APPLICATION_URLENCODED string = "application/x-www-form-urlencoded"
	TEXT_HTML              string = "text/html"
	TEXT_PLAIN             string = "text/plain"
	OCTET_STREAM           string = "application/octet-stream"
)

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)
