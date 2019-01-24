#!/bin/bash

# 第一步：先设置Go环境变量
# jenkeins Go插件自动设置的GOROOT不正确，这里重新设置一次
export GOROOT=$GOROOT/go

# 将GOPATH设置为当前目录
export GOPATH=${WORKSPACE}
export GOBIN=$GOPATH/bin
export GOPKG=$GOPATH/pkg
export PATH=$PATH:$GOBIN:$GOROOT/bin

# 第二步: git clone项目的代码
cd ${WORKSPACE}
rm -rf src

# git clone项目
export PRJ=`git config --get remote.origin.url | sed 's/\.git$//' | sed 's/^https:\/\///' | sed 's/^ssh:\/\///' | sed 's/^git@//' | sed 's/:/\//'`

# Clean directory
git clean -dfx

# Move project into GOPATH
mkdir -p $GOPATH/src/$PRJ
ls -1 | grep -v ^src | xargs -I{} mv {} $GOPATH/src/$PRJ/

# 第三步: rpm构建
cd $GOPATH/src/$PRJ/build/rpm
module=$1

pkgName=""
function packageRPM() {
    # 以下几个值是jenkins构建设置的变量
    echo "moduleName=$module"
    echo "version=${version}"
    echo "iteration=${BUILD_NUMBER}"
    echo "env=${env}"
    pkgName=`sh build_rpm.sh $module ${version} ${BUILD_NUMBER} ${env}`
    echo $pkgName
}

packageRPM

# 第四步: 同步rpm到yum源
function syncToYum() {
	echo "sync to yum:$pkgName"
	cp $pkgName /www/jenkins_home/.jenkins/formalpackages/
}

# 如果勾选同步到yum源会把rpm包同步到yum源上
if $yum; then
	syncToYum
fi

