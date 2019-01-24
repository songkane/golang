// Package config 配置文件定义
// Created by chenguolin 2018-11-16
package config

// Config api模块配置文件结构定义
type Config struct {
	Deploy    *DeployConf    `toml:"deploy"`
	Mysql     *MysqlConf     `toml:"mysql"`
	Redis     *RedisConf     `toml:"redis"`
	Memcached *MemcachedConf `toml:"memcache"`
	Kafka     *KafkaConf     `toml:"kafka"`
}

// DeployConf deploy config
type DeployConf struct {
	Environment string   `toml:"environment"`
	Host        []string `toml:"host"`
	DevopsHost  string   `toml:"devops_host"`
}

// MysqlConf mysql config
type MysqlConf struct {
	Master             string   `toml:"master"`
	Slaves             []string `toml:"slaves"`
	Port               uint16   `toml:"port"`
	Username           string   `toml:"username"`
	Password           string   `toml:"password"`
	Dbname             string   `toml:"dbname"`
	MaxConnCount       int      `toml:"max_conn_count"`
	MaxIdleConnCount   int      `toml:"max_idle_conn_count"`
	ConnWaitTimeMs     int      `toml:"conn_wait_time_ms"`
	ConnIdleWaitTimeMs int      `toml:"conn_idle_time_ms"`
	ConnTimeoutMs      int      `toml:"conn_timeout_ms"`
	WriteTimeoutMs     int      `toml:"write_timeout_ms"`
	ReadTimeoutMs      int      `toml:"read_timeout_ms"`
}

// RedisConf redis config
type RedisConf struct {
	Master           string   `toml:"master"`
	Slaves           []string `toml:"slaves"`
	MaxConns         int      `toml:"max_conns"`
	MaxIdle          int      `toml:"max_idle"`
	ReadTimeoutMs    int      `toml:"read_timeout_ms"`
	WriteTimeoutMs   int      `toml:"write_timeout_ms"`
	ConnectTimeoutMs int      `toml:"connect_timeout_ms"`
	IdleTimeoutMs    int      `toml:"idle_timeout_ms"`
	DB               int      `toml:"DB"`
	Password         string   `toml:"Password"`
}

// MemcachedConf memcached conifg
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

// KafkaConf kafka config
type KafkaConf struct {
	ZK      []string `toml:"zk"`
	Brokers []string `toml:"brokers"`
}
