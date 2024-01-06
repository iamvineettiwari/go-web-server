package http

import (
	"bytes"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
	Query   map[string]string
}

func Decode(content []byte) *Request {
	data := bytes.Split(content, CRLF)
	headers := make(map[string]string)

	requestLine := bytes.Split(data[0], SPACE)
	bodyStart := 0

	for i, item := range data[1:] {
		headerData := bytes.Split(item, SPACE)

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
		Body:    bodyContent,
		Query:   queries,
	}
}

func decodePath(data []byte) (string, map[string]string) {
	splitedPath := bytes.Split(data, []byte("?"))

	queryMap := make(map[string]string)
	path := string(splitedPath[0])

	if len(splitedPath) > 1 {
		queries := bytes.Split(splitedPath[1], []byte("&"))

		for _, query := range queries {
			splitedQuery := bytes.Split(query, []byte("="))

			if len(splitedQuery) == 2 {
				queryMap[string(splitedQuery[0])] = string(splitedQuery[1])
			}
		}
	}

	return path, queryMap
}
