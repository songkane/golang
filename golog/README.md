# golog
golog 是一个golang logger框架，内部使用uber开源的zap框架, 性能好灵活性高。

1. 支持不同log级别输出
2. 支持不同io.Writer输出
3. 支持HTTP Service Access Log输出定制

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
