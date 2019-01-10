/*
Package config toml配置文件读取
Created by chenguolin 2018-11-16
*/
package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// LoadFrom 加载配置文件
func LoadFrom(filePath string, conf interface{}) {
	_, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}

	_, err = toml.DecodeFile(filePath, conf)
	if err != nil {
		panic(err)
	}
}

// MysqlConf mysql配置
type MysqlConf struct {
	Master              string   `toml:"master"`
	Slaves              []string `toml:"slaves"`
	Port                uint16   `toml:"port"`
	Username            string   `toml:"username"`
	Password            string   `toml:"password"`
	Dbname              string   `toml:"dbname"`
	ConnectionTimeoutMs int64    `toml:"connection_timeout_ms"`
	ReadTimeoutMs       int64    `toml:"read_timeout_ms"`
	WriteTimeoutMs      int64    `toml:"write_timeout_ms"`
	MaxWaitTimeMs       int64    `toml:"max_wait_time_ms"`
	MaxIdleTimeMs       int64    `toml:"max_idle_time_ms"`
	MaxOpen             int      `toml:"max_open"`
	MaxIdle             int      `toml:"max_idle"`
}

// RedisConf redis配置
type RedisConf struct {
	Master           string   `toml:"master"`
	Slaves           []string `toml:"slaves"`
	MaxConns         int      `toml:"max_conns"`
	MaxIdle          int      `toml:"max_idle"`
	ReadTimeoutMs    int64    `toml:"read_timeout_ms"`
	WriteTimeoutMs   int64    `toml:"write_timeout_ms"`
	ConnectTimeoutMs int64    `toml:"connect_timeout_ms"`
	IdleTimeoutMs    int64    `toml:"idle_timeout_ms"`
	DB               int      `toml:"DB"`
	Password         string   `toml:"Password"`
}

// MemcachedConf Memcached配置
type MemcachedConf struct {
	Servers              string `toml:"servers"`
	MaxActiveConnections int    `toml:"max_active_connections"`
	MaxIdleConnections   int    `toml:"max_idle_connections"`
	MaxWaitTimeMs        int    `toml:"max_wait_time_ms"`
	MaxIdleTimeMs        int    `toml:"max_idle_time_ms"`
	ConnectionTimeoutMs  int    `toml:"connection_timeout_ms"`
	ReadTimeoutMs        int    `toml:"read_timeout_ms"`
	WriteTimeoutMs       int    `toml:"write_timeout_ms"`
	Rehash               bool   `toml:"rehash"`
}

// KafkaConf kafka配置
type KafkaConf struct {
	ZK      []string `toml:"zk"`
	Brokers []string `toml:"brokers"`
}

// EthereumConf ethereum配置
type EthereumConf struct {
	Address       string `toml:"address"`
	TimeoutSecond int    `toml:"timeout_second"`
}
