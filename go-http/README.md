# HTTP Service
用Gin实现的通用的HTTP Service框架, 项目结构规范约定如下

1. cmd: 可执行程序
    * api: 对外提供API接口服务
    * internal: 对内提供API接口服务
2. bizerror: 业务定义的错误类型
3. config: 配置文件读取
4. docs: 相关文档
5. instance: 应用初始化
6. pkg: 业务代码，每个业务是一个package
7. scripts: 相关脚本
8. vendor: golang 依赖库
9. version: 版本相关信息, 使用ldflags动态设置变量值

# 编译构建
## api
$ cd cmd/api
$ sh build.sh pre
$ cd dist/bin
$ ./go-http-api version  (可以查看相关信息)

## internal
$ cd cmd/internal
$ sh build.sh pre
$ cd dist/bin
$ ./go-http-internal version  (可以查看相关信息)
