package http

import "github.com/iamvineettiwari/go-web-server/http/status"

func Encode(data []byte, status status.HttpStatus) []byte {
	response := []byte{}

	response = append(response, VERSION...)
	response = append(response, SPACE...)
	response = append(response, []byte(status)...)
	response = append(response, CRLF...)
	response = append(response, CRLF...)
	response = append(response, data...)
	response = append(response, CRLF...)

	return response
}
