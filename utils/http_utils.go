package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetRequestPath(host, path string) string {
	return fmt.Sprintf("%s%s", host, path)
}

func PUT(host, path string, data interface{}, headers map[string][]string) (string, int, error) {
	return doHttpRequest("PUT", host, path, data, headers)
}

func GET(host, path string, data interface{}, headers map[string][]string) (string, int, error) {
	return doHttpRequest("GET", host, path, data, headers)
}

func POST(host, path string, data interface{}, headers map[string][]string) (string, int, error) {
	return doHttpRequest("POST", host, path, data, headers)
}

func DELETE(host, path string, data interface{}, headers map[string][]string) (string, int, error) {
	return doHttpRequest("DELETE", host, path, data, headers)
}

func doHttpRequest(method, host, path string, data interface{}, headers map[string][]string) (string, int, error) {

	params, err := EncodeData(data)
	if err != nil {
		return "", -1, err
	}

	req, err := http.NewRequest(method, GetRequestPath(host, path), params)
	if err != nil {
		return "", -1, err
	}

	req.Header.Set("User-Agent", "DD-server/v1")
	req.Header.Set("Content-Type", "application/json")

	// timeout 5 minutes
	client := &http.Client{Timeout: time.Duration(300) * time.Minute}

	log.Printf("http request to: %s/%s", host, path)
	log.Printf("http Header : %v", req.Header)
	// log.Printf("http req : %v", req)

	resp, err := client.Do(req)
	statusCode := -1
	if resp != nil {
		statusCode = resp.StatusCode
	}
	if err != nil {
		return "", statusCode, err
	}
	log.Printf("http response code : %d", statusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", statusCode, err
	}

	return string(body), statusCode, nil
}

func EncodeData(v interface{}) (*bytes.Buffer, error) {
	param := bytes.NewBuffer(nil)
	j, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	if _, err := param.Write(j); err != nil {
		return nil, err
	}
	return param, nil
}
