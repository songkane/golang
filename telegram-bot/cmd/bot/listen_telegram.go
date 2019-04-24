// Package main listen telegram
// Modified by chenguolin 2018-03-16
package main

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"time"

	"github.com/go-playground/pool"

	golog "gitlab.local.com/golang/go-log"
	"gitlab.local.com/golang/telegram-bot/cmd/bot/conf"
	"gitlab.local.com/golang/telegram-bot/pkg/telegram-bot-api"
)

var (
	defaultPattern *regexp.Regexp
	processWorkers int
)

func initConfig() {
	cfg := config.GetConf()
	defaultPattern = regexp.MustCompile(cfg.Telegram.MessagePattern["default"])
}

func loopProcess(workerPool pool.Pool, config tgbotapi.UpdateConfig, systemCfg *config.Config) int {
	golog.Info(fmt.Sprintf("main package loopProcess function config.Offset: %d", config.Offset))

	// 1. get latest update message
	updates, err := tgBot.GetUpdates(config)
	if err != nil {
		golog.Warn(fmt.Sprintf("main package loopProcess function call tbBot.GetUpdates failed: %v, retrying in 3 seconds...", err))
		time.Sleep(time.Second * 3)
		return config.Offset
	}

	// 2. batch process
	b := workerPool.Batch()

	startTime := time.Now().UnixNano() / 1000000
	golog.Info(fmt.Sprintf("main package looProcess function start to process telegram message.  total message: %d", len(updates)))

	// 3. 批量处理完一批再处理下一批
	for _, update := range updates {
		if update.UpdateID >= config.Offset {
			config.Offset = update.UpdateID + 1

			var message *tgbotapi.Message
			if update.Message != nil {
				message = update.Message
			} else if update.EditedMessage != nil {
				message = update.EditedMessage
			}

			if message == nil || message.From == nil {
				continue
			}

			// 消息来自组信息
			if message.Chat.IsGroup() || message.Chat.IsSuperGroup() {
				// 如果消息来自机器人过滤这个消息
				if message.From.IsBot {
					golog.Warn(fmt.Sprintf("current message from bot. Message: %v", message))
					continue
				}

				msg := message
				b.Queue(
					func(wu pool.WorkUnit) (interface{}, error) {
						if wu.IsCancelled() {
							return nil, nil
						}
						// 处理消息请求发币
						processMsg(msg)
						return true, nil
					},
				)
			}
		}
	}
	b.QueueComplete()
	b.WaitAll()

	endTime := time.Now().UnixNano() / 1000000
	golog.Info(fmt.Sprintf("listen_telegram listenTelegram process updates done. process time: %d ms",
		endTime-startTime))

	return config.Offset
}

func processMsg(msg *tgbotapi.Message) {
	// 1. request process
	startTime := time.Now().UnixNano() / 1000000
	golog.Info("main package processMsg function start to handler request php api ~")

	matchStringList := defaultPattern.FindStringSubmatch(msg.Text)
	if matchStringList != nil || len(matchStringList) > 0 {
		msg.Text = matchStringList[1]
	}

	endTime := time.Now().UnixNano() / 1000000
	golog.Info(fmt.Sprintf("listen_telegram processMsg request php api done! process time: %d ms", endTime-startTime))

	// 2. send message
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, "OK ~")
	newMsg.ReplyToMessageID = msg.MessageID
	tgBot.Send(newMsg)
}

func broadcastMsg2Group(killChan chan os.Signal) {
loop:
	for {
		select {
		case _, ok := <-killChan:
			if !ok {
				log.Info("===================== Cloing broadcast message 2 group now ~ =====================")
				break loop
			}
		default:
			cfg := config.GetConf()
			//如果开关关闭的,则不进行广播
			if !cfg.Telegram.BroadcastSwitch {
				continue
			}
			for _, groupID := range cfg.Telegram.WhiteGroupID {
				newMsg := tgbotapi.NewMessage(groupID, "Hello World ~")
				tgBot.Send(newMsg)
			}
			time.Sleep(time.Duration(cfg.Telegram.BroadcastInterval) * time.Second)
		}
	}
}

func listenTelegram() {
	golog.Info("===================== Start Telegram bot =====================")

	// 1. 配置初始化
	initConfig()

	systemCfg := config.GetConf()

	// 2. 监听系统kill信号
	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan, os.Interrupt, os.Kill)

	// 3. new一个线程池
	workerPool := pool.NewLimited(uint(processWorkers))
	defer workerPool.Close()

	// 4. telegram update updateConfig
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Limit = systemCfg.Telegram.UpdatesLimit
	updateConfig.Timeout = 30

	// 5. 定期往Telegram群广播消息
	go broadcastMsg2Group(killChan)

	// 6. 监听Telegram
loop:
	for {
		select {
		case <-killChan:
			// 等待关闭广播消息的协程
			close(killChan)
			time.Sleep(time.Duration(systemCfg.Telegram.BroadcastInterval) * time.Second)
			break loop
		default:
			newOffset := loopProcess(workerPool, updateConfig, systemCfg)
			updateConfig.Offset = newOffset
		}
	}

	golog.Info("===================== Closing Telegram bot now ~ =====================")
}
