// Package config struct define
// Created by chenguolin 2018-11-16
package config

// Config struct define
type Config struct {
	Deploy   *DeployConf   `toml:"deploy"`   //部署配置
	Mysql    *MysqlConf    `toml:"mysql"`    //mysql配置
	Redis    *RedisConf    `toml:"redis"`    //redis配置
	Memcache *MemcacheConf `toml:"memcache"` //mc配置
	Kafka    *KafkaConf    `toml:"kafka"`    //kafka配置
}

// DeployConf deploy config
type DeployConf struct {
	APIAddr      string `toml:"api_addr"`      //api HTTP server listen address
	InternalAddr string `toml:"internal_addr"` //internal HTTP server listen address
	DevopsAddr   string `toml:"devops_addr"`   //dev ops HTTP server listen address
}

// MysqlConf mysql config
type MysqlConf struct {
	Master           string   `toml:"master"`
	Slaves           []string `toml:"slaves"`
	Port             int      `toml:"port"`
	Username         string   `toml:"username"`
	Password         string   `toml:"password"`
	Dbname           string   `toml:"dbname"`
	MaxOpenConnCount int      `toml:"max_open_conn_count"`
	MaxIdleConnCount int      `toml:"max_idle_conn_count"`
	ConnWaitTimeMs   int      `toml:"conn_wait_time_ms"`
	ConnIdleTimeMs   int      `toml:"conn_idle_time_ms"`
	ConnTimeoutMs    int      `toml:"conn_timeout_ms"`
	WriteTimeoutMs   int      `toml:"write_timeout_ms"`
	ReadTimeoutMs    int      `toml:"read_timeout_ms"`
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

// MemcacheConf memcached conifg
type MemcacheConf struct {
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
