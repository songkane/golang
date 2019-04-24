// Package config define
// Created by chenguolin 2019-04-24
package config

import "time"

// 全局变量
var conf Config
var configFilePath string

// LoadFrom load from file
func LoadFrom(filePath string) {
	Load(filePath, &conf)
	configFilePath = filePath
	goUpdateConfig()
}

func goUpdateConfig() {
	go update()
}

func update() {
	defaultConf := conf
	defer func() {
		if err := recover(); err != nil {
			conf = defaultConf
			goUpdateConfig()
		}
	}()

	for {
		select {
		case <-time.Tick(time.Second * 10):
			Load(configFilePath, &conf)
		}
	}
}

// GetConf get config
func GetConf() *Config {
	return &conf
}

// Config type
type Config struct {
	Telegram     *Telegram         `toml:"telegram"`
}

// Telegram type
type Telegram struct {
	MessagePattern    map[string]string `toml:"message_pattern"`
	UpdatesLimit      int               `toml:"updates_limit"`
	BroadcastSwitch   bool              `toml:"broadcast_switch"`   //广播开关
	BroadcastInterval int               `toml:"broadcast_interval"` //广播间隔
	WhiteGroupID      []int64           `toml:"white_group_id"`
	Bot               *TelegramBot      `toml:"bot"`
}

// TelegramBot type
// type TelegramMessagePattern map[string]string
type TelegramBot struct {
	Name  string `toml:"name"`
	Token string `toml:"token"`
}
