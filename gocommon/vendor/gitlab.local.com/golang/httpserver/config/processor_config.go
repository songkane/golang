/*
Package config processor模块配置文件读取
Created by chenguolin 2018-11-16
*/
package config

// ProcessorConfig 配置文件结构定义
type ProcessorConfig struct {
	Deploy    *DeployConf    `toml:"deploy"`
	Mysql     *MysqlConf     `toml:"mysql"`
	Redis     *RedisConf     `toml:"redis"`
	Kafka     *KafkaConf     `toml:"kafka"`
	Ethereum  *EthereumConf  `toml:"ethereum"`
	Processor *ProcessorConf `toml:"processor"`
}

// ProcessorConf 处理机配置定义
type ProcessorConf struct {
	HandleConcurrency int    `toml:"handle_concurrency"`
	Topic             string `toml:"topic"`
	ConsumerGroupID   string `toml:"consumer_group_id"`
}

var processorConf *ProcessorConfig

// GetProcessorConfig 获取processor模块配置
func GetProcessorConfig(filePath string) *ProcessorConfig {
	if processorConf != nil {
		return processorConf
	}
	if filePath == "" {
		return nil
	}

	conf := &ProcessorConfig{}
	LoadFrom(filePath, conf)
	processorConf = conf
	return conf
}

// GetProcessorDeployConf 获取processor模块部署配置
func GetProcessorDeployConf() *DeployConf {
	if processorConf == nil {
		return nil
	}

	return processorConf.Deploy
}

// GetProcessorMysqlConf 获取processor模块Mysql配置
func GetProcessorMysqlConf() *MysqlConf {
	if processorConf == nil {
		return nil
	}

	return processorConf.Mysql
}

// GetProcessorRedisConf 获取processor模块Redis配置
func GetProcessorRedisConf() *RedisConf {
	if processorConf == nil {
		return nil
	}

	return processorConf.Redis
}

// GetProcessorKafkaConf 获取processor模块Kafka配置
func GetProcessorKafkaConf() *KafkaConf {
	if processorConf == nil {
		return nil
	}

	return processorConf.Kafka
}

// GetProcessorEthereumConf 获取processor模块Etherum配置
func GetProcessorEthereumConf() *EthereumConf {
	if processorConf == nil {
		return nil
	}

	return processorConf.Ethereum
}

// GetProcessorProcessorConf 获取processor模块处理机配置
func GetProcessorProcessorConf() *ProcessorConf {
	if processorConf == nil {
		return nil
	}

	return processorConf.Processor
}
