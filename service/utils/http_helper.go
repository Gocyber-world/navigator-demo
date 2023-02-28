package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func SendHttpRequest(url string, method string, data interface{}, token string) (int, string, error) {
	var body string
	var err error
	if data != nil {
		var jsonData []byte
		jsonData, err = json.Marshal(data)
		if err != nil {
			return -1, "", err
		}
		body = string(jsonData)
	}
	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	// using token if it is not empty
	if token != "" {
		req.Header.Set("Authorization", "Token "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, "", err
	}
	defer resp.Body.Close()
	var result []byte
	result, err = io.ReadAll(resp.Body)
	if err != nil {
		return -1, "", err
	}
	return resp.StatusCode, string(result), nil
}
