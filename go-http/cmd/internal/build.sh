# 编译构建脚本
#!/bin/bash

# 删除老的dist目录 同时创建几个子目录
rm -rf dist
mkdir -p ./dist/bin
mkdir -p ./dist/conf
mkdir -p ./dist/logs

# TODO(@cgl) 自行更改app_name
app_name="go-http-internal"
env_type="$1"

# 编译构建
# GOOS=linux GOARCH=amd64 go build -o ${app_name}
go build -o ${app_name}

mv ${app_name} ./dist/bin/
cp ./bin/* ./dist/bin
cp ./conf/config-${env_type}.toml ./dist/conf/config.toml

# 输出构建成功
echo "${appName} build successful ~"
