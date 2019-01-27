/*
Package config cron模块配置文件读取
Created by chenguolin 2018-11-16
*/
package config

// CronConfig 配置文件结构定义
type CronConfig struct {
	Deploy *DeployConf `toml:"deploy"`
	Mysql  *MysqlConf  `toml:"mysql"`
}

var cronConf *CronConfig

// GetCronConfig 获取Cron模块配置
func GetCronConfig(filePath string) *CronConfig {
	if cronConf != nil {
		return cronConf
	}
	if filePath == "" {
		return nil
	}

	conf := &CronConfig{}
	LoadFrom(filePath, conf)
	cronConf = conf
	return conf
}

// GetCronDeployConf 获取cron模块部署配置
func GetCronDeployConf() *DeployConf {
	if cronConf == nil {
		return nil
	}

	return cronConf.Deploy
}

// GetCronMysqlConf 获取cron模块mysql配置
func GetCronMysqlConf() *MysqlConf {
	if cronConf == nil {
		return nil
	}

	return cronConf.Mysql
}
