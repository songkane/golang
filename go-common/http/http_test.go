// Package http unit test
// Created by chenguolin 2019-02-11
package http

import (
	"fmt"
	"io/ioutil"
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
	resposne, err := Get(client, url, nil, retry)
	if err != nil {
		t.Fatal("TestGet case 1 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestGet case 1 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestGet case 1 status != 200 failed ~")
	}
	defer resposne.Body.Close()
	body, err := ioutil.ReadAll(resposne.Body)
	if err != nil {
		t.Fatal("TestGet case 1 ioutil.ReadAll err != nil")
	}
	if len(body) <= 0 {
		t.Fatal("TestGet case 1 len(body) <= 0")
	}
	// fmt.Printf(string(body))

	// case 2
	header := make(map[string]string)
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	header["Accept-Language"] = "ja,zh-CN;q=0.8,zh;q=0.6"
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0"
	header["Content-Type"] = "application/x-www-form-urlencoded"

	resposne, err = Get(client, url, header, retry)
	if err != nil {
		t.Fatal("TestGet case 2 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestGet case 2 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestGet case 2 status != 200 failed ~")
	}
	defer resposne.Body.Close()
	body, err = ioutil.ReadAll(resposne.Body)
	if err != nil {
		t.Fatal("TestGet case 2 ioutil.ReadAll err != nil")
	}
	if len(body) <= 0 {
		t.Fatal("TestGet case 2 len(body) <= 0")
	}
	// fmt.Printf(string(body))

	// case 3
	// https://mbd.baidu.com/newspage/data/landingsuper?context=%7B%22nid%22%3A%22news_8517410035950326229%22%7D&n_type=0&p_from=1
	url = "https://mbd.baidu.com/newspage/data/landingsuper"
	context := "{\"nid\":\"news_8517410035950326229\"}"
	nType := 0
	pFrom := 1
	url = fmt.Sprintf("%s?context=%s&n_type=%d&p_from=%d", url, context, nType, pFrom)

	header = make(map[string]string)
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	header["Accept-Language"] = "ja,zh-CN;q=0.8,zh;q=0.6"
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0"
	header["Content-Type"] = "application/x-www-form-urlencoded"

	resposne, err = Get(client, url, header, retry)
	if err != nil {
		t.Fatal("TestGet case 3 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestGet case 3 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestGet case 3 status != 200 failed ~")
	}
	defer resposne.Body.Close()
	body, err = ioutil.ReadAll(resposne.Body)
	if err != nil {
		t.Fatal("TestGet case 3 ioutil.ReadAll err != nil")
	}
	if len(body) <= 0 {
		t.Fatal("TestGet case 3 len(body) <= 0")
	}
	// fmt.Println(string(body))
}

func TestPost(t *testing.T) {
	client := NewHTTPClient(time.Duration(10 * time.Second))
	url := "http://www.baidu.com"
	retry := 3

	// case 1
	resposne, err := Post(client, url, nil, nil, retry)
	if err != nil {
		t.Fatal("TestPost case 1 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestPost case 1 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestPost case 1 status != 200 failed ~")
	}
	defer resposne.Body.Close()
	body, err := ioutil.ReadAll(resposne.Body)
	if err != nil {
		t.Fatal("TestPost case 1 ioutil.ReadAll err != nil")
	}
	if len(body) <= 0 {
		t.Fatal("TestPost case 1 len(body) <= 0")
	}
	//fmt.Printf(string(body))

	// case 2
	header := make(map[string]string)
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	header["Accept-Language"] = "ja,zh-CN;q=0.8,zh;q=0.6"
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0"
	header["Content-Type"] = "application/x-www-form-urlencoded"

	resposne, err = Post(client, url, nil, header, retry)
	if err != nil {
		t.Fatal("TestPost case 2 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestPost case 2 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestPost case 2 status != 200 failed ~")
	}

	// read body
	defer resposne.Body.Close()
	body, err = ioutil.ReadAll(resposne.Body)
	if err != nil {
		t.Fatal("TestPost case 2 ioutil.ReadAll err != nil")
	}
	if len(body) <= 0 {
		t.Fatal("TestPost case 2 len(body) <= 0")
	}
	//fmt.Printf(string(body))

	// case 3
	url = "https://mbd.baidu.com/newspage/data/landingsuper"
	context := "{\"nid\":\"news_8517410035950326229\"}"
	nType := 0
	pFrom := 1
	url = fmt.Sprintf("%s?context=%s&n_type=%d&p_from=%d", url, context, nType, pFrom)

	header = make(map[string]string)
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	header["Accept-Language"] = "ja,zh-CN;q=0.8,zh;q=0.6"
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0"
	header["Content-Type"] = "application/x-www-form-urlencoded"

	resposne, err = Post(client, url, nil, header, retry)
	if err != nil {
		t.Fatal("TestPost case 3 err != nil failed ~")
	}
	if resposne == nil {
		t.Fatal("TestPost case 3 response == nil failed ~")
	}
	if resposne.StatusCode != 200 {
		t.Fatal("TestPost case 3 status != 200 failed ~")
	}

	// read body
	defer resposne.Body.Close()
	body, err = ioutil.ReadAll(resposne.Body)
	if err != nil {
		t.Fatal("TestPost case 3 ioutil.ReadAll err != nil")
	}
	if len(body) <= 0 {
		t.Fatal("TestPost case 3 len(body) <= 0")
	}
	//fmt.Printf(string(body))
}
