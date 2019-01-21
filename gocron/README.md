# 介绍
定时任务模块，如果有新的业务逻辑只需要在handler目录下新增一个go文件, 同时在main函数里面AddFunc新增一行配置即可

# sample
```
c := cron.NewCron()

// TODO add handle
c.AddHandle(cron.NewScheduler(cron.WithSecond(5), time.Now()), handle.GetCrawlAddressTxsHandle())

// run cron
c.Start()
fmt.Println("Start cron handler ...")

// shutdown
stopSignalChan := make(chan os.Signal, 1)
signal.Notify(stopSignalChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
sig := <-stopSignalChan
if sig != nil {
    fmt.Println("got system signal:" + sig.String() + ", going to shutdown crontab")
	// Stop the scheduler (does not stop any jobs already running).
	c.Stop()
}
fmt.Println("Stop cron handler ~")
```
