// Package signature 单元测试
// Created by chenguolin 2018-11-17
package signature

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetSignature(t *testing.T) {
	c := &gin.Context{}
	c.Request = &http.Request{}
	c.Request.URL = &url.URL{}
	c.Request.URL.Path = "http://localhost:8080/user/info.json"

	form := make(map[string][]string)
	form["id"] = []string{"123456"}
	form["name"] = []string{"chenguolin"}
	form["sigTime"] = []string{"1234567890"}
	c.Request.Form = form

	sig := GetSignature(c)
	if sig != "1e21082ecda4fae10214b17feda9195b" {
		t.Fatal("TestGetSignature get sig != \"1e21082ecda4fae10214b17feda9195b\"")
	}
}

func TestSortFormValues(t *testing.T) {
	form := make(map[string][]string)
	form["id"] = []string{"123456"}
	form["name"] = []string{"chenguolin"}
	form["sigTime"] = []string{"1234567890"}
	formValues := sortFormValues(form)
	if len(formValues) != 2 {
		t.Fatal("TestSortFormValues len(formValues) != 2")
	}
	if formValues[0] != "123456" {
		t.Fatal("TestSortFormValues formValues[0] != \"123456\"")
	}
	if formValues[1] != "chenguolin" {
		t.Fatal("TestSortFormValues formValues[0] != \"chenguolin\"")
	}
}

func TestGenerateSignature(t *testing.T) {
	path := "http://localhost:8080/user/info.json"
	paramsArr := []string{"123456", "chenguolin"}
	sigTime := string("1234567890")
	sig := generateSignature(path, paramsArr, sigTime)
	if sig != "1e21082ecda4fae10214b17feda9195b" {
		t.Fatal("TestGenerateSignature get sig != \"1e21082ecda4fae10214b17feda9195b\"")
	}
}
