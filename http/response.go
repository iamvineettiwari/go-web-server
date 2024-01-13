package http

import (
	"net"

	_status "github.com/iamvineettiwari/go-web-server/http/status"
)

type Response struct {
	conn    net.Conn
	status  _status.HttpStatus
	headers map[string]string
}

func GetResponseWriter(conn net.Conn) *Response {
	return &Response{
		conn:    conn,
		headers: make(map[string]string),
		status:  _status.HTTP_200_OK,
	}
}

func (r *Response) SetStatus(status _status.HttpStatus) {
	r.status = status
}

func (r *Response) AddHeader(key, value string) {
	r.headers[key] = value
}

func (r *Response) RemoveHeader(key string) {
	if _, present := r.headers[key]; !present {
		return
	}

	delete(r.headers, key)
}

func (r *Response) Write(content []byte) {
	defer r.conn.Close()

	data := Encode(content, r.status, r.headers)
	r.conn.Write(data)
}

func (r *Response) Send(content []byte, status _status.HttpStatus) {
	r.SetStatus(status)
	r.Write(content)
}
