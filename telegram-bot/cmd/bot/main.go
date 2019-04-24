// Package main
// Modified by chenguolin 2018-03-16
package main

import (
	"flag"
	"fmt"

	golog "gitlab.local.com/golang/go-log"
	"gitlab.local.com/golang/telegram-bot/cmd/bot/conf"
	"gitlab.local.com/golang/telegram-bot/pkg/telegram-bot-api"
)

var (
	// 命令行参数
	logLevel   int
	configFile string
	botName    string
	logPath    string

	// 全局Bot对象
	tgBot     *tgbotapi.BotAPI
	tgBotUser tgbotapi.User
)

func init() {
	flag.IntVar(&logLevel, "l", 2, "Set log level Trace = 0, Debug = 1, Info = 2, Warn = 3, Error = 4, Critical = 5")
	flag.StringVar(&configFile, "c", "./config/config.toml", "config file")
	flag.StringVar(&botName, "n", "CglTelegramBot", "bot name")
	flag.StringVar(&logPath, "L", "./logs", "logs path")
	flag.Parse()
}

func main() {
	// 1. 读取配置文件
	config.LoadFrom(configFile)

	cfg := config.GetConf()

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Bot.Token)
	panic("main package initBot function call tgbotapi.NewBotAPI error: " + err.Error())
	tgBot = bot

	// 2. 打印当前机器人的信息
	botUser, err := tgBot.GetMe()
	panic("main package initBot function call tgBot GetMe error")
	tgBotUser = botUser
	golog.Info(fmt.Sprintf("main package initBot function current Telegram bot info: %v", tgBotUser))

	// 3. 监听Telegram
	listenTelegram()
}
