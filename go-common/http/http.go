// Package http net/http 封装
// Created by chenguolin 2019-02-11
package http

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	golog "gitlab.local.com/golang/go-log"
)

// NewHTTPClient new http.Client
func NewHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

// Get http get request
// @httpClient required
// @url required
// @body optional
// @header optional
// @retry required
func Get(httpClient *http.Client, url string, body []byte,
	header map[string]string, retry int) (*http.Response, error) {

	if retry <= 0 {
		retry = 1
	}

	return request(httpClient, "GET", url, body, header, retry)
}

// Post http post request
// @httpClient required
// @url required
// @body optional
// @header optional
// @retry required
func Post(httpClient *http.Client, url string, body []byte,
	header map[string]string, retry int) (*http.Response, error) {

	if retry <= 0 {
		retry = 1
	}

	return request(httpClient, "POST", url, body, header, retry)
}

// Request new http request
// @url required
// @body optional
// @header optional
// @retry required
// @timeout required
func request(httpClient *http.Client, method, url string, body []byte,
	header map[string]string, retry int) (*http.Response, error) {

	// new request
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// request add header
	for k, v := range header {
		req.Header.Add(k, v)
	}

	// retry
	for i := 0; i < retry; i++ {
		response, err := httpClient.Do(req)
		if err != nil {
			golog.Warn("httpClient.Do error", golog.Object("request", req),
				golog.Object("error", err))
			continue
		}

		return response, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("http %s %s, request failed~", method, url))
}
