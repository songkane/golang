// Package init 全局实例对象调用初始化通用入口
// Created by chenguolin 2018-11-16
package init

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"gitlab.local.com/golang/gocommon/time"
	mysql "gitlab.local.com/golang/gomysql"
)

// 通用变量定义
var (
	// 部署环境
	environment string

	// 资源类
	mysqlClient *mysql.Mysql
	mcClient    memcache.Client
)

// PkgInitFunc pkg init function
type PkgInitFunc func()

var initFuncs = make([]PkgInitFunc, 0)

// AppInit 初始化进程Context
// module 模块名称 api、cron、processor
// filePath 表示配置文件路径
func AppInit(filePath string) {
	// 获取api模块配置
	// cfg := config.GetConfig(filePath)

	// TODO

	// pkg下相关的service执行Init函数
	for _, f := range initFuncs {
		f()
	}
	fmt.Println(time.GetCurrentTime(), "api context init successfully")
}

// AppInitTest 测试环境进程context初始化
// 方便写单元测试
func AppInitTest() {
	// TODO 用户需要自行修改本地配置文件路径
	// filePath := os.Getenv("GOPATH") + "/src/gitlab.local.com/golang/httpserver/cmd/api/conf/config-pre.toml"
	// TODO
}

// AddInitFunc 添加初始化函数
// name 表示pkg名称
// f 表示初始化函数
func AddInitFunc(name string, f PkgInitFunc) {
	initFuncs = append(initFuncs, f)
	fmt.Println(time.GetCurrentTime(), "register init func:"+name)
}

// GetEnvironment get environment
func GetEnvironment() string {
	return environment
}

// GetMysqlClient get mysql client
func GetMysqlClient() *mysql.Mysql {
	return mysqlClient
}

// GetMcClient get memcache client
func GetMcClient() memcache.Client {
	return mcClient
}
