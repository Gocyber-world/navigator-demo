package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func SendHttpRequest(url string, method string, bodyData interface{}, queryData interface{}, token string) (int, string, error) {
	var body string
	var err error
	if bodyData != nil {
		var jsonData []byte
		jsonData, err = json.Marshal(bodyData)
		if err != nil {
			return -1, "", err
		}
		body = string(jsonData)
	}
	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	// using token if it is not empty
	if token != "" {
		req.Header.Set("Authorization", "Token "+token)
	}

	// set querystring if exists
	if queryData != nil {
		q := req.URL.Query()
		queryDataMap := queryData.(map[string]interface{})
		for key, value := range queryDataMap {
			q.Add(key, value.(string))
		}
		req.URL.RawQuery = q.Encode()
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
