/*
Package config api模块配置文件读取
Created by chenguolin 2018-11-16
*/
package config

// APIConfig api模块配置文件结构定义
type APIConfig struct {
	Deploy     *DeployConf     `toml:"deploy"`
	Mysql      *MysqlConf      `toml:"mysql"`
	Redis      *RedisConf      `toml:"redis"`
	Memcached  *MemcachedConf  `toml:"memcached"`
	Kafka      *KafkaConf      `toml:"kafka"`
	Dependency *DependencyConf `toml:"dependency"`
}

// DeployConf 部署配置定义
type DeployConf struct {
	Environment string   `toml:"environment"`
	Host        []string `toml:"host"`
	DevopsHost  string   `toml:"devops_host"`
}

// DependencyConf 依赖配置定义
type DependencyConf struct {
	TokenServiceURL string `toml:"token_service_url"`
	IPServiceURL    string `toml:"ip_service_url"`
}

var apiConf *APIConfig

// GetAPIConfig 获取api模块配置
func GetAPIConfig(filePath string) *APIConfig {
	if apiConf != nil {
		return apiConf
	}
	if filePath == "" {
		return nil
	}

	conf := &APIConfig{}
	LoadFrom(filePath, conf)
	apiConf = conf
	return conf
}

// GetAPIDeployConf 获取部署相关配置
func GetAPIDeployConf() *DeployConf {
	if apiConf == nil {
		return nil
	}

	return apiConf.Deploy
}

// GetAPIMysqlConf 获取Mysql相关配置
func GetAPIMysqlConf() *MysqlConf {
	if apiConf == nil {
		return nil
	}

	return apiConf.Mysql
}

// GetAPIRedisConf 获取Redis相关配置
func GetAPIRedisConf() *RedisConf {
	if apiConf == nil {
		return nil
	}

	return apiConf.Redis
}

// GetAPIMemcachedConf 获取MC相关配置
func GetAPIMemcachedConf() *MemcachedConf {
	if apiConf == nil {
		return nil
	}

	return apiConf.Memcached
}

// GetAPIKafkaConf 获取MC相关配置
func GetAPIKafkaConf() *KafkaConf {
	if apiConf == nil {
		return nil
	}

	return apiConf.Kafka
}

// GetAPIDependencyConf 获取api依赖服务配置
func GetAPIDependencyConf() *DependencyConf {
	if apiConf == nil {
		return nil
	}

	return apiConf.Dependency
}
