#!/bin/sh

cd `dirname $0`
currDir=`pwd`

cd ../../
projectRootDir=`pwd`

cd ${currDir}

# 构建需要5个变量
# 模块名称 例如api、cron、processor
moduleName=$1
# 版本号 例如1.0.0
version=$2
# 当前编译编号 例如88
iteration=$3
# 环境类型 例如pre、beta、release
envType=$4

# 模块名
moduleDir="${projectRootDir}/cmd/${moduleName}"

# 调用模块下的build.sh脚本
appName=`sh ${moduleDir}/build.sh "v${version}" "${envType}"`

# rpm包名称
rpmPrefixName="${appName}"
rpmFileName="${appName}-${version}-${iteration}.${envType}.x86_64.rpm"

# rpm包安装路径
deployDir="/www/${appName}/"

# 编译结果目录
distDir="${moduleDir}/dist"

# 构建rpm包
# TODO 用户需要变更 --category 参数对应的值
# TODO 用户需要变更 -m 参数对应的值
fpm -s dir -t rpm --prefix "${deployDir}" -n "${rpmPrefixName}" -v "${version}" --iteration "${iteration}.${envType}" \
 --category 'xxxx/Projects' --description "${appName}" \
 --license 'Commercial' -m 'xxxx' --before-install ${projectRootDir}/build/rpm/before.sh \
 --after-install ${projectRootDir}/build/rpm/after.sh --before-upgrade ${projectRootDir}/build/rpm/before.sh \
 --after-upgrade ${projectRootDir}/build/rpm/after.sh -C ${distDir} bin/ conf/ logs/ >&2

echo ${rpmFileName}

