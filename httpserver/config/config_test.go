/*
Package config 单元测试
Created by chenguolin 2018-11-16
*/
package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin/json"
)

func TestGetApiConfig(t *testing.T) {
	// case 1
	filePath := ""
	apiConf := GetConfig(filePath)
	if apiConf != nil {
		t.Fatal("TestGetApiConfig case 1 get apiConf != nil")
	}

	// case 2
	// TODO 用户需要自行修改路径
	filePath = os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/cmd/api/conf/config-pre.toml"
	apiConf = GetConfig(filePath)
	if apiConf == nil {
		t.Fatal("TestGetApiConfig case 2 get apiConf == nil")
	}

	// deploy
	if apiConf.Deploy == nil {
		t.Fatal("apiConf.Deploy is nil")
	}
	// mysql
	if apiConf.Mysql == nil {
		t.Fatal("apiConf.Mysql is nil")
	}
	// redis
	if apiConf.Redis == nil {
		t.Fatal("apiConf.Redis is nil")
	}
	// memcached
	if apiConf.Memcached == nil {
		t.Fatal("apiConf.Memcached is nil")
	}
	// kafka
	if apiConf.Kafka == nil {
		t.Fatal("apiConf.Kafka is nil")
	}
	// dependency
	if apiConf.Dependency == nil {
		t.Fatal("apiConf.Dependency is nil")
	}

	bytes, err := json.Marshal(apiConf)
	if err != nil {
		t.Fatal("json Marshal conf error")
	}
	fmt.Println(string(bytes))
}

func TestGetApiDeployConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/cmd/api/conf/config-pre.toml"
	apiConf = GetConfig(filePath)
	if apiConf == nil {
		t.Fatal("TestGetApiDeployConf get apiConf == nil")
	}

	deployConf := GetDeployConf()
	if deployConf == nil {
		t.Fatal("TestGetApiDeployConf get deployConf == nil")
	}
}

func TestGetApiMysqlConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/cmd/api/conf/config-pre.toml"
	apiConf = GetConfig(filePath)
	if apiConf == nil {
		t.Fatal("TestGetApiMysqlConf get apiConf == nil")
	}

	mysqlConf := GetMysqlConf()
	if mysqlConf == nil {
		t.Fatal("TestGetApiMysqlConf get mysqlConf == nil")
	}
}

func TestGetApiRedisConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/cmd/api/conf/config-pre.toml"
	apiConf = GetConfig(filePath)
	if apiConf == nil {
		t.Fatal("TestGetApiRedisConf get apiConf == nil")
	}

	redisConf := GetRedisConf()
	if redisConf == nil {
		t.Fatal("TestGetApiRedisConf get redisConf == nil")
	}
}

func TestGetApiMemcachedConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/cmd/api/conf/config-pre.toml"
	apiConf = GetConfig(filePath)
	if apiConf == nil {
		t.Fatal("TestGetApiMemcachedConf get apiConf == nil")
	}

	mcConf := GetMemcachedConf()
	if mcConf == nil {
		t.Fatal("TestGetApiMemcachedConf get mcConf == nil")
	}
}

func TestGetApiKafkaConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/cmd/api/conf/config-pre.toml"
	apiConf = GetConfig(filePath)
	if apiConf == nil {
		t.Fatal("TestGetApiKafkaConf get apiConf == nil")
	}

	kafkaConf := GetKafkaConf()
	if kafkaConf == nil {
		t.Fatal("TestGetApiKafkaConf get kafkaConf == nil")
	}
}

func TestGetApiDependencyConf(t *testing.T) {
	// TODO 用户需要自行修改路径
	filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/cmd/api/conf/config-pre.toml"
	apiConf = GetConfig(filePath)
	if apiConf == nil {
		t.Fatal("TestGetApiDependencyConf get apiConf == nil")
	}

	dependencyConf := GetDependencyConf()
	if dependencyConf == nil {
		t.Fatal("TestGetApiDependencyConf get dependencyConf == nil")
	}
}
