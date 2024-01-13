package http

import "bytes"

func getURLEncodedValue(data []byte) map[string]any {
	bodyData := make(map[string]any)

	splitedData := bytes.Split(data, []byte("&"))

	for _, item := range splitedData {
		itemSplited := bytes.Split(item, []byte("="))

		if len(itemSplited) == 2 {
			bodyData[string(itemSplited[0])] = string(itemSplited[1])
		}
	}

	return bodyData
}
