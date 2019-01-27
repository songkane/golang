# 模块编译构建脚本
# 如果使用/bin/sh 会报sh: Syntax error: "}" unexpected相关错误
# 具体原因是是linux将sh指向了dash而不是bash

#!/bin/bash

# 函数总共4个参数
function build {
    # project根目录
    projectDir=`pwd`

    # 函数参数
    moduleName="$1"
    version=`cat $projectDir/VERSION`
    iteration="$2"
    envType="$3"

    # 进入模块目录
    moduleDir="$projectDir/cmd/$moduleName"
    cd $moduleDir

    # 编译
    appName=`sh build.sh "v${version}" ${envType}`
    echo "build ${appName} end"

    # rpm构建
    distDir="${moduleDir}/dist"
    deployDir="/www/${appName}/"
    rpmPrefixName="${appName}"
    rpmFileName="${appName}-${version}-${iteration}.${envType}.x86_64.rpm"

    fpm -s dir -t rpm --prefix "${deployDir}" -n "${rpmPrefixName}" -v "${version}" --iteration "${iteration}.${envType}"\
      --category 'Golang/httpserver' --url 'http://gitlab.local.com' --description "${appName}"\
      --license 'Commercial' -m 'Httpserver' --before-install ${projectDir}/scripts/build/rpm/before.sh\
      --after-install ${projectDir}/scripts/build/rpm/after.sh --before-upgrade ${projectDir}/scripts/build/rpm/before.sh\
      --after-upgrade ${projectDir}/scripts/build/rpm/after.sh -C ${distDir} bin/ conf/ logs/ >&2
    echo ${rpmFileName}

    cd $projectDir
}

# 编译迭代ID
iteration=${CI_PIPELINE_ID}
# 环境类型 默认pre
envType=pre

for moduleName in api cron processor
do
	build ${moduleName} ${iteration} ${envType}
done
