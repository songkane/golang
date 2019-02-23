// Package rpc json-rpc client
// Created by chenguolin 2019-02-22
package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client json-rpc client
type Client struct {
	Server   string `json:"server"`   //required
	User     string `json:"usre"`     //optional
	Password string `json:"password"` //optional
}

// Request json-rpc request struct
type Request struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

// Response json-rpc response struct
type Response struct {
	JSONRPC string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Result  interface{}    `json:"result,omitempty"`
	Error   *responseError `json:"error,omitempty"`
}

// responseError response error
type responseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

// NewClient new client
func NewClient(server string) *Client {
	return &Client{
		Server: server,
	}
}

// genHTTPReq generate http request
func (c *Client) genHTTPReq(data []byte) (*http.Request, error) {
	// generate http request
	httpReq, err := http.NewRequest("POST", c.Server, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create req error: %s", err)
	}
	if c.User != "" {
		httpReq.SetBasicAuth(c.User, c.Password)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	return httpReq, nil
}

// SetAuth set rpc auth user and password
func (c *Client) SetAuth(user, password string) {
	c.User = user
	c.Password = password
}

// Call json-rpc call
func (c *Client) Call(req *Request) (*Response, error) {
	if req == nil {
		return nil, errors.New("Args req is nil")
	}

	// set request
	req.JSONRPC = "2.0"
	if req.ID <= 0 {
		req.ID = 1
	}
	// not found param need set 2 []
	if req.Params == nil {
		req.Params = make([]string, 0)
	}

	// marshal request
	bs, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal req error: %s", err)
	}

	// generate http request
	httpReq, err := c.genHTTPReq(bs)
	if err != nil {
		return nil, err
	}

	// http Do
	client := http.DefaultClient
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("HTTP client Do error: %s", err)
	}

	// read http response
	body, err := ioutil.ReadAll(httpRes.Body)
	defer httpRes.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll HTTP response body error: %s", err)
	}

	// json.Unmarshal
	res := &Response{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %s, body: %s", err, string(body))
	}

	return res, nil
}

// BatchCall json-rpc call
func (c *Client) BatchCall(reqs []*Request) ([]*Response, error) {
	if reqs == nil {
		return nil, errors.New("Args Reqs is nil")
	}

	// set request
	for i, req := range reqs {
		req.JSONRPC = "2.0"
		if req.ID <= 0 {
			req.ID = i + 1
		}
	}

	// marshal request
	bs, err := json.Marshal(reqs)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal reqs error: %s", err)
	}

	// generate http request
	httpReq, err := c.genHTTPReq(bs)
	if err != nil {
		return nil, err
	}

	// http Do
	client := http.DefaultClient
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("HTTP client Do error: %s", err)
	}

	// read http response
	body, err := ioutil.ReadAll(httpRes.Body)
	defer httpRes.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll HTTP response body error: %s", err)
	}

	// json.Unmarshal
	ress := make([]*Response, 0)
	err = json.Unmarshal(body, &ress)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %s, body: %s", err, string(body))
	}

	return ress, nil
}
