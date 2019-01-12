# 介绍
定时任务模块，如果有新的业务逻辑只需要在handler目录下新增一个go文件, 同时在main函数里面AddFunc新增一行配置即可

# sample
```
func Print1() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "hello world 1 ~")
}

func Print2() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "hello world 2 ~")
}

func Print3() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "hello world 3 ~")
}

func Print4() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "hello world 4 ~")
}

func Print5() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "hello world 5 ~")
}

func main() {
	// 1. start cron
	c := cron.New()
	// all handlers
	// etc one by one
	now := time.Now()
	c.AddJob(cron.NewScheduler(cron.WithSecond(5), now), Print1)
	c.AddJob(cron.NewScheduler(cron.WithSecond(5), now), Print2)
	c.AddJob(cron.NewScheduler(cron.WithMinute(1), now), Print3)
	c.AddJob(cron.NewScheduler(cron.WithHour(1), time.Date(2019, 01, 05, 15, 10, 00, 00, time.Local)), Print4)
	c.AddJob(cron.NewScheduler(cron.WithDay(1), time.Date(2019, 01, 05, 00, 00, 00, 00, time.Local)), Print5)
	c.Start()

	// 2. block shutdown
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stopSignalChan
	if sig != nil {
		fmt.Println("got system signal:" + sig.String() + ", going to shutdown crontab")
		// Stop the scheduler (does not stop any jobs already running).
		c.Stop()
	}
}
```
