// Package http unit test
// Created by chenguolin 2019-02-11
package http

import (
	"testing"
	"time"
)

func TestNewHTTPClient(t *testing.T) {
	client := NewHTTPClient(time.Duration(10 * time.Second))
	if client == nil {
		t.Fatal("TestNewHTTPClient failed ~")
	}
}

func TestGet(t *testing.T) {
	client := NewHTTPClient(time.Duration(10 * time.Second))
	url := "http://www.baidu.com"
	retry := 3

	// case 1
	resposne, err := Get(client, url, nil, nil, retry)
	if err != nil {
		t.Fatal("TestGet case 1 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestGet case 1 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestGet case 1 status != 200 failed ~")
	}

	// case 2
	header := make(map[string]string)
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	header["Accept-Language"] = "ja,zh-CN;q=0.8,zh;q=0.6"
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0"

	resposne, err = Get(client, url, nil, header, retry)
	if err != nil {
		t.Fatal("TestGet case 2 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestGet case 2 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestGet case 2 status != 200 failed ~")
	}
}

func TestPost(t *testing.T) {
	client := NewHTTPClient(time.Duration(10 * time.Second))
	url := "http://www.baidu.com"
	retry := 3

	// case 1
	resposne, err := Post(client, url, nil, nil, retry)
	if err != nil {
		t.Fatal("TestGet case 1 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestGet case 1 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestGet case 1 status != 200 failed ~")
	}

	// case 2
	header := make(map[string]string)
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	header["Accept-Language"] = "ja,zh-CN;q=0.8,zh;q=0.6"
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0"

	resposne, err = Post(client, url, nil, header, retry)
	if err != nil {
		t.Fatal("TestGet case 2 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestGet case 2 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestGet case 2 status != 200 failed ~")
	}
}
