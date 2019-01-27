# golog
golog 是一个golang 日志框架，内部使用uber开源的zap框架, 性能好灵活性高。

1. 支持不同log级别输出
2. 支持不同io.Writer输出
3. 支持HTTP Service Access Log定制输出
4. 支持输出到文件并按固定格式切割文件

# sample
```
// use 1
// 默认level为Info
// 默认使用JSONEncoder
// 默认输出到Stdout
golog.Info("use 1", golog.String("name", "chenguolin"))

// use 2
// level设置为Info
// 使用Console Encodr
// 输出到Stderr
log, err := golog.NewLogger(golog.WithInfoLevel(), golog.WithConsoleEncoder(), golog.WithOutput(os.Stderr))
if err != nil {
    panic(err)
}
golog.SetLogger(log)
golog.Info("use 2", golog.String("name", "chenguolin"))
```

# AccessLog
```
// set logger
fileName := os.Getenv("GOPATH") + "/src/gitlab.local.com/golog/sample/access.log"
f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0660)
if err != nil {
    fmt.Printf("open file error", err)
}
defer f.Close()

log, err := golog.NewLogger(golog.WithOutput(io.Writer(f)), golog.WithJSONEncoder(), golog.WithInfoLevel())
if err != nil {
    fmt.Printf("golog NewLogger error", err)
}

// start gin HTTP server
r := gin.New()
r.GET("/ping", golog.AccessLogFunc(log), HelloServer)
// listen and serve on 0.0.0.0:8899
r.Run(":8899")
```

# RotateWriter
```
// new rotate writer
fileName := os.Getenv("GOPATH") + "/src/gitlab.local.com/golog/sample/access.log"
// 按小时切割
writer, err := golog.NewRotateWriter(fileName, "20060102-15")
if err != nil {
    fmt.Println("golog NewRotateWriter error", err)
}

// new logger
log, err := golog.NewLogger(golog.WithOutput(writer), golog.WithJSONEncoder(), golog.WithInfoLevel())
if err != nil {
    fmt.Printf("golog NewLogger error", err)
}

// start gin HTTP server
r := gin.New()
r.GET("/ping", golog.AccessLogFunc(log), HelloServer)
// listen and serve on 0.0.0.0:8899
r.Run(":8899")
```