package http

import (
	"fmt"

	"github.com/iamvineettiwari/go-web-server/http/status"
)

func Encode(data []byte, status status.HttpStatus, headers map[string]string) []byte {
	response := []byte{}

	response = append(response, VERSION...)
	response = append(response, SPACE...)
	response = append(response, []byte(status)...)
	appendHeaders(&response, headers)

	response = append(response, CRLF...)
	response = append(response, CRLF...)
	response = append(response, data...)
	response = append(response, CRLF...)

	return response
}

func appendHeaders(response *[]byte, headers map[string]string) {
	for key, value := range headers {
		data := fmt.Sprintf("%s: %s", key, value)
		(*response) = append((*response), CRLF...)
		(*response) = append((*response), []byte(data)...)
	}
}
