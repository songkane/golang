# Go语言源码编码规范检查
#!/bin/bash

echo '********** golint check start ... **********'

# check has errors
# mac下sed命令和linux有所区别
# mac sed: sed -i 需要带一个字符串，用来备份源文件，这个字符串加在源文件名后面组成备份文件名
# sed -i "bs" 's/Atl/Dog/g' example.txt 则会生成example.txtbs 的备份文件
# sed -i "" 's/Atl/Dog/g' example.txt 如果这个字符串长度为0，就是说是个空串，那么不备份
errors=$(go list ./... | grep -v vendor | sed -e s=gitlab.local.com/golang/go-healthcheck/=./= | grep -v gitlab | xargs -n 1 golint 2>&1)

# grep errors has go file
echo "${errors}" | grep "\.go"

# $? equal 0 found check error
if [ $? -eq 0 ]; then
    echo '********** golint check failed ~ **********'
    echo "${errors}"
    exit 1
else
    echo '********** golint check ok ~ **********'
    exit 0
fi
