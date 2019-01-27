// Package init unit test
// Created by chenguolin 2018-11-16
package init

import (
	"fmt"
	"os"
	"testing"
)

func TestInitAppContext(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/go-http/cmd/config/conf/config-pre.toml"
	AppInit(filePath)
}

func TestInitTestContext(t *testing.T) {
	AppInitTest()
}

func TestAddInitFunc(t *testing.T) {
	f := func() { fmt.Println("hello word func1") }
	AddInitFunc("func", f)

	f2 := func() { fmt.Println("hello word func2") }
	AddInitFunc("func2", f2)

	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/go-http/cmd/config/conf/config-pre.toml"
	AppInit(filePath)
}

func TestGetMysqlClient(t *testing.T) {
	// TODO
}

func TestGetRedisClient(t *testing.T) {
	// TODO
}

func TestGetMcClient(t *testing.T) {
	// TODO
}

func TestGetKafkaConf(t *testing.T) {
	// TODO
}
