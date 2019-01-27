// Package init common function
// Created by chenguolin 2018-11-16
package init

import (
	"errors"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis"

	"gitlab.local.com/golang/go-http/config"
	"gitlab.local.com/golang/go-mysql"
)

// newMysqlClient new mysql client
func newMysqlClient(conf *config.MysqlConf) (*mysql.Mysql, error) {
	if conf == nil {
		return nil, errors.New("config.MysqlConf is nil")
	}

	mysqlCfg := &mysql.Config{}
	mysqlCfg.SetMaster(conf.Master)
	mysqlCfg.SetSlaves(conf.Slaves)
	mysqlCfg.SetPort(conf.Port)
	mysqlCfg.SetUserName(conf.Username)
	mysqlCfg.SetPassword(conf.Password)
	mysqlCfg.SetDBName(conf.Dbname)
	if conf.MaxOpenConnCount > 0 {
		mysqlCfg.SetMaxOpenConnCount(conf.MaxOpenConnCount)
	}
	if conf.MaxIdleConnCount > 0 {
		mysqlCfg.SetMaxIdleConnCount(conf.MaxIdleConnCount)
	}
	if conf.ConnWaitTimeMs > 0 {
		mysqlCfg.SetConnWaitTimeMs(conf.ConnWaitTimeMs)
	}
	if conf.ConnIdleTimeMs > 0 {
		mysqlCfg.SetConnIdleTimeMs(conf.ConnIdleTimeMs)
	}
	if conf.ConnTimeoutMs > 0 {
		mysqlCfg.SetConnTimeoutMs(conf.ConnTimeoutMs)
	}
	if conf.WriteTimeoutMs > 0 {
		mysqlCfg.SetWriteTimeoutMs(conf.WriteTimeoutMs)
	}
	if conf.ReadTimeoutMs > 0 {
		mysqlCfg.SetReadTimeoutMs(conf.ReadTimeoutMs)
	}

	return mysql.NewMysql(mysqlCfg)
}

// newRedisClient new redis client
func newRedisClient(conf *config.RedisConf) (*redis.ClusterClient, error) {
	// TODO (@cgl)
	return nil, nil
}

// newMcClient new memcache client
func newMcClient(conf *config.MemcacheConf) (*memcache.Client, error) {
	// TODO (@cgl)
	return nil, nil
}
