# 编译构建脚本
#!/bin/bash

# 删除老的dist目录 同时创建几个子目录
rm -rf dist
mkdir -p ./dist/bin
mkdir -p ./dist/conf
mkdir -p ./dist/logs

app_name="go-http-internal"
env_type="$1"

# 编译构建
project_root="../.."
target_os="linux"
target_arch="amd64"
go_project="gitlab.local.com/golang/go-http"
version=`grep "^version" ${project_root}/Changelog | tail -1 | cut -d " " -f2`
build_date=`date -u +'%Y-%m-%dT%H:%M:%SZ'`
git_revision=`git rev-parse --short HEAD`

# build
GOOS="${target_os}" GOARCH="${target_arch}" go build -v -ldflags \
    "-X ${go_project}/version.Version=${version} -X ${go_project}/version.BuildDate=${build_date} -X ${go_project}/version.BuildCommit=${git_revision}" \
    -o ${app_name}
if [ $? -ne 0 ]; then
    echo "Failed to build ${app_name}"
    exit 1
fi

# 拷贝文件到dist目录
mv ${app_name} ./dist/bin/
cp ./bin/* ./dist/bin
cp ${project_root}/config/conf/config-${env_type}.toml ./dist/conf/config.toml
chmod 775 ./dist/bin/*
if [ $? -ne 0 ]; then
    echo "Failed to cp config toml file"
    exit 1
fi

# 输出构建成功
echo "${appName} build successful ~"

