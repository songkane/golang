// Package http signature unit test
// Created by chenguolin 2018-11-17
package http

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGenSignature(t *testing.T) {
	c := &gin.Context{}
	c.Request = &http.Request{}
	c.Request.URL = &url.URL{}
	c.Request.URL.Path = "http://localhost:8080/user/info.json"

	form := make(map[string][]string)
	form["id"] = []string{"123456"}
	form["name"] = []string{"chenguolin"}
	form["sigTime"] = []string{"1234567890"}
	c.Request.Form = form

	sig := GenSignature(c)
	if sig != "3d1f23aa2ca39142544ef66d64a231bb" {
		t.Fatal("TestGetSignature get sig != \"3d1f23aa2ca39142544ef66d64a231bb\"")
	}
}
