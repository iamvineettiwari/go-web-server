package http

import (
	"errors"
)

type Request struct {
	data    []byte
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Query   map[string]any
	Body    []byte
}

func (r *Request) ParseRequest() error {
	if r.data == nil || len(r.data) == 0 {
		return errors.New("Empty body")
	}

	r.Body = r.data

	if _, presentContentType := r.Headers["Content-Type"]; !presentContentType {
		if r.Method == GET {
			r.Headers["Content-Type"] = OCTET_STREAM
		} else {
			r.Headers["Content-Type"] = APPLICATION_URLENCODED
		}
	}

	return nil
}
