// Package config common unit test
// Created by chenguolin 2018-11-16
package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin/json"
)

func TestLoadFrom(t *testing.T) {
	// load api conf
	// TODO 路径需要用户自行修改
	conf := loadFrom(os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/cmd/api/conf/config-pre.toml")
	if conf == nil {
		t.Fatal("TestLoadFrom loadFrom api conf is nil")
	}

	bytes, err := json.Marshal(conf)
	if err != nil {
		t.Fatal("TestLoadFrom err != nil")
	}
	fmt.Println(string(bytes))
}
