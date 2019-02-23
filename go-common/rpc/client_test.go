// Package rpc json-rpc client unit test
// Created by chenguolin 2019-02-22
package rpc

import (
	"testing"
)

func TestClient_genHTTPReq(t *testing.T) {
	client := NewClient("https://mainnet.infura.io")

	// case 1
	httpReq, err := client.genHTTPReq(nil)
	if err != nil {
		t.Fatal("TestClient_genHTTPReq case 1 err != nil")
	}
	if httpReq == nil {
		t.Fatal("TestClient_genHTTPReq case 1 httpReq == nil")
	}

	// case 2
	data := []byte(string("\"jsonrpc\":\"2.0\""))
	httpReq, err = client.genHTTPReq(data)
	if err != nil {
		t.Fatal("TestClient_genHTTPReq case 2 err != nil")
	}
	if httpReq == nil {
		t.Fatal("TestClient_genHTTPReq case 2 httpReq == nil")
	}
}

func TestClient_SetAuth(t *testing.T) {
	client := NewClient("https://mainnet.infura.io")
	client.SetAuth("root", "123456")

	if client.User != "root" {
		t.Fatal("TestClient_SetAuth client.User != \"root\"")
	}
	if client.Password != "123456" {
		t.Fatal("TestClient_SetAuth client.Password != \"123456\"")
	}
}

func TestClient_Call(t *testing.T) {
	client := NewClient("https://mainnet.infura.io")

	// case 1
	res, err := client.Call(nil)
	if err == nil {
		t.Fatal("TestClient_Call case 1 err == nil")
	}
	if res != nil {
		t.Fatal("TestClient_Call case 2 res != nil")
	}

	// case 2
	req := &Request{
		Method: "eth_getTransactionByHash",
		Params: []string{"0x7b8cf5a4da03e8ad9eebe93fc7f6b7441073c20e993bb86e956c8e8d4b1c1e9a"},
	}
	res, err = client.Call(req)
	if err != nil {
		t.Fatal("TestClient_Call case 2 err != nil")
	}
	if res.JSONRPC != "2.0" {
		t.Fatal("TestClient_Call case 2 res.JSONRPC != 2.0")
	}
	if res.ID != 1 {
		t.Fatal("TestClient_Call case 2 res.JSONRPC != 1")
	}
	if res.Result == nil {
		t.Fatal("TestClient_Call case 2 res.Result == nil")
	}
	if res.Error != nil {
		t.Fatal("TestClient_Call case 2 res.Error != nil")
	}

	// case 3
	req = &Request{
		ID:     100,
		Method: "eth_getTransactionByHash",
		Params: []string{"0x7b8cf5a4da03e8ad9eebe93fc7f6b7441073c20e993bb86e956c8e8d4b1c1e9a"},
	}
	res, err = client.Call(req)
	if err != nil {
		t.Fatal("TestClient_Call case 3 err != nil")
	}
	if res.JSONRPC != "2.0" {
		t.Fatal("TestClient_Call case 3 res.JSONRPC != 2.0")
	}
	if res.ID != 100 {
		t.Fatal("TestClient_Call case 3 res.JSONRPC != 100")
	}
	if res.Result == nil {
		t.Fatal("TestClient_Call case 3 res.Result == nil")
	}
	if res.Error != nil {
		t.Fatal("TestClient_Call case 3 res.Error != nil")
	}

	// case 4
	req = &Request{
		Method: "eth_blockNumber",
		Params: make([]string, 0),
	}
	res, err = client.Call(req)
	if err != nil {
		t.Fatal("TestClient_Call case 4 err != nil")
	}
	if res.JSONRPC != "2.0" {
		t.Fatal("TestClient_Call case 4 res.JSONRPC != 2.0")
	}
	if res.ID != 1 {
		t.Fatal("TestClient_Call case 4 res.JSONRPC != 1")
	}
	if res.Result == nil {
		t.Fatal("TestClient_Call case 4 res.Result == nil")
	}
	if res.Error != nil {
		t.Fatal("TestClient_Call case 4 res.Error != nil")
	}

	// case 5
	req = &Request{
		Method: "eth_blockNumber",
	}
	res, err = client.Call(req)
	if err != nil {
		t.Fatal("TestClient_Call case 5 err != nil")
	}
	if res.JSONRPC != "2.0" {
		t.Fatal("TestClient_Call case 5 res.JSONRPC != 2.0")
	}
	if res.ID != 1 {
		t.Fatal("TestClient_Call case 5 res.JSONRPC != 1")
	}
	if res.Result == nil {
		t.Fatal("TestClient_Call case 5 res.Result == nil")
	}
	if res.Error != nil {
		t.Fatal("TestClient_Call case 5 res.Error != nil")
	}

	// case 6
	// error
	req = &Request{
		ID:     10000,
		Method: "eth_getTransactionByHash",
		Params: []string{"0xb8cf5a4da03e8ad9eebe93fc7f6b7441073c20e993bb86e956c8e8d4b1c1e9a"},
	}
	res, err = client.Call(req)
	if err != nil {
		t.Fatal("TestClient_Call case 6 err != nil")
	}
	if res.JSONRPC != "2.0" {
		t.Fatal("TestClient_Call case 6 res.JSONRPC != 2.0")
	}
	if res.ID != 10000 {
		t.Fatal("TestClient_Call case 6 res.JSONRPC != 100")
	}
	if res.Result != nil {
		t.Fatal("TestClient_Call case 6 res.Result != nil")
	}
	if res.Error == nil {
		t.Fatal("TestClient_Call case 6 res.Error == nil")
	}
	if res.Error.Code != -32602 {
		t.Fatal("TestClient_Call case 6 res.Error.Code != -32602")
	}
	if res.Error.Message != "invalid argument 0: json: cannot unmarshal hex string of odd length into Go value of type common.Hash" {
		t.Fatal("TestClient_Call case 6 res.Error.Message error")
	}
}

func TestClient_BatchCall(t *testing.T) {
	client := NewClient("https://mainnet.infura.io")

	// case 1
	ress, err := client.BatchCall(nil)
	if err == nil {
		t.Error("TestClient_BatchCall case 1 err == nil")
	}
	if ress != nil {
		t.Error("TestClient_BatchCall case 1 ress != nil")
	}

	// case 2
	reqs := make([]*Request, 0)
	req1 := &Request{
		Method: "eth_getTransactionByHash",
		Params: []string{"0x7b8cf5a4da03e8ad9eebe93fc7f6b7441073c20e993bb86e956c8e8d4b1c1e9a"},
	}
	req2 := &Request{
		ID:     100,
		Method: "eth_getTransactionByHash",
		Params: []string{"0x7b8cf5a4da03e8ad9eebe93fc7f6b7441073c20e993bb86e956c8e8d4b1c1e9a"},
	}
	req3 := &Request{
		Method: "eth_blockNumber",
		Params: make([]string, 0),
	}
	req4 := &Request{
		Method: "eth_blockNumber",
	}
	req5 := &Request{
		Method: "eth_getTransactionByHash",
		Params: []string{"0xb8cf5a4da03e8ad9eebe93fc7f6b7441073c20e993bb86e956c8e8d4b1c1e9a"},
	}
	reqs = append(reqs, req1, req2, req3, req4, req5)

	ress, err = client.BatchCall(reqs)
	if err != nil {
		t.Error("TestClient_BatchCall case 2 err != nil")
	}
	if ress == nil {
		t.Error("TestClient_BatchCall case 2 ress == nil")
	}
	if len(ress) != 5 {
		t.Error("TestClient_BatchCall case 2 len(ress) != 3")
	}
	// ress 0~3 successed
	// ress 4 failed
	if ress[0].Result == nil {
		t.Error("TestClient_BatchCall case 2 ress[0].Result == nil")
	}
	if ress[1].Result == nil {
		t.Error("TestClient_BatchCall case 2 ress[1].Result == nil")
	}
	if ress[2].Result == nil {
		t.Error("TestClient_BatchCall case 2 ress[2].Result == nil")
	}
	if ress[3].Result == nil {
		t.Error("TestClient_BatchCall case 2 ress[3].Result == nil")
	}
	if ress[4].Error == nil {
		t.Error("TestClient_BatchCall case 2 ress[4].Error == nil")
	}
}
