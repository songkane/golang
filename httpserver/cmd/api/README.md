# 介绍
Golang实现的HTTP Service，提供HTTP接口调用

1. bin: 启动、停止相关脚本
2. conf: 配置文件
3. controller: api接口层逻辑
4. request: 请求和响应通用函数封装

# API
API分成2个部分
1. 业务API
2. 开发运维api常用于版本检查、健康检查等

## 业务API
1. 查询用户信息: curl "http://localhost:8080/user/select.json?uid=111"
2. 创建新用户: curl "http://localhost:8080/user/create.json" -d "uid=111&name=cgl&phone=123456"
3. 更新用户信息: curl "http://localhost:8080/user/update.json" -d "uid=111&name=cgl2&phone=123456789
4. 删除用户: curl "http://localhost:8080/user/delete.json" -d "uid=111"

## 开发运维API
1. 服务状态检查: curl "http://localhost:9010/devops/status"
