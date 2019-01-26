// Package config unit test
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
	conf := loadFrom(os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/config/conf/config-pre.toml")
	if conf == nil {
		t.Fatal("TestLoadFrom loadFrom api conf is nil")
	}

	bytes, err := json.Marshal(conf)
	if err != nil {
		t.Fatal("TestLoadFrom err != nil")
	}
	fmt.Println(string(bytes))
}

func TestGetConfig(t *testing.T) {
	// case 1
	filePath := ""
	conf := GetConfig(filePath)
	if conf != nil {
		t.Fatal("TestGetConfig case 1 get apiConf != nil")
	}

	// case 2
	// TODO 用户需要自行修改路径
	filePath = os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/config/conf/config-pre.toml"
	conf = GetConfig(filePath)
	if conf == nil {
		t.Fatal("TestGetConfig case 2 get apiConf == nil")
	}

	// deploy
	if conf.Deploy == nil {
		t.Fatal("apiConf.Deploy is nil")
	}
	// mysql
	if conf.Mysql == nil {
		t.Fatal("apiConf.Mysql is nil")
	}
	// redis
	if conf.Redis == nil {
		t.Fatal("apiConf.Redis is nil")
	}
	// memcache
	if conf.Memcache == nil {
		t.Fatal("apiConf.Memcache is nil")
	}
	// kafka
	if conf.Kafka == nil {
		t.Fatal("apiConf.Kafka is nil")
	}

	bytes, err := json.Marshal(conf)
	if err != nil {
		t.Fatal("json Marshal conf error")
	}
	fmt.Println(string(bytes))
}

func TestGetDeployConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/config/conf/config-pre.toml"
	conf := GetConfig(filePath)
	if conf == nil {
		t.Fatal("TestGetDeployConf GetConfig == nil")
	}

	deploy := GetDeployConf()
	if deploy == nil {
		t.Fatal("TestGetDeployConf GetDeployConf == nil")
	}
	fmt.Println(deploy)
}

func TestGetMysqlConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/config/conf/config-pre.toml"
	conf = GetConfig(filePath)
	if conf == nil {
		t.Fatal("TestGetMysqlConf GetConfig == nil")
	}

	mysql := GetMysqlConf()
	if mysql == nil {
		t.Fatal("TestGetMysqlConf GetMysqlConf == nil")
	}
}

func TestGetRedisConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/config/conf/config-pre.toml"
	conf = GetConfig(filePath)
	if conf == nil {
		t.Fatal("TestGetRedisConf GetConfig == nil")
	}

	redis := GetRedisConf()
	if redis == nil {
		t.Fatal("TestGetRedisConf GetRedisConf == nil")
	}
}

func TestGetMemcacheConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/config/conf/config-pre.toml"
	conf = GetConfig(filePath)
	if conf == nil {
		t.Fatal("TestGetMemcacheConf GetConfig == nil")
	}

	mc := GetMemcacheConf()
	if mc == nil {
		t.Fatal("TestGetMemcacheConf GetMemcacheConf == nil")
	}
}

func TestGetKafkaConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/config/conf/config-pre.toml"
	conf = GetConfig(filePath)
	if conf == nil {
		t.Fatal("TestGetKafkaConf GetConfig == nil")
	}

	kafka := GetKafkaConf()
	if kafka == nil {
		t.Fatal("TestGetKafkaConf GetKafkaConf == nil")
	}
}
