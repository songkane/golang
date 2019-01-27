# 介绍
通用的rpm构建脚本

1. build_rpm.sh 运行该脚本会构建rpm包
2. after.sh rpm包安装或更新后置脚本
3. before.sh rpm包安装或更新前置脚本

例如: sh build_rpm.sh api 1.0.0 10 pre
1. 模块名称为 api
2. 版本为 1.0.0
3. 构建编号为 10
4. 环境为 pre
