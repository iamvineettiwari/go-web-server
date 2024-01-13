package http

import (
	"bytes"
)

func Decode(content []byte) *Request {
	data := bytes.Split(content, CRLF)
	headers := make(map[string]string)

	requestLine := bytes.Split(data[0], SPACE)
	bodyStart := 0

	for i, item := range data[1:] {
		headerData := bytes.SplitN(item, SPACE, 2)

		if len(headerData) == 1 {
			bodyStart = i + 2
			break
		}

		key := string(headerData[0])
		val := string(headerData[1])

		headers[key[:len(key)-1]] = val
	}

	bodyContent := []byte{}

	for _, body := range data[bodyStart:] {
		body = append(body, CRLF...)
		bodyContent = append(bodyContent, body...)
	}

	path, queries := decodePath(requestLine[1])

	return &Request{
		Method:  string(requestLine[0]),
		Path:    path,
		Version: string(requestLine[2]),
		Headers: headers,
		data:    bodyContent,
		Query:   queries,
	}
}

func decodePath(data []byte) (string, map[string]any) {
	splitedPath := bytes.Split(data, []byte("?"))

	queryMap := make(map[string]any)
	path := string(splitedPath[0])

	if len(splitedPath) > 1 {
		queryMap = getURLEncodedValue(splitedPath[1])
	}

	return path, queryMap
}
