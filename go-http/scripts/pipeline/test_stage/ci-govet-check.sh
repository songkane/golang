# Go语言源码中静态错误的简单工具
#!/bin/bash

echo '********** go vet check start ... **********'

# check has error
# mac下sed命令和linux有所区别
# mac sed: sed -i 需要带一个字符串，用来备份源文件，这个字符串加在源文件名后面组成备份文件名
# sed -i "bs" 's/Atl/Dog/g' example.txt 则会生成example.txtbs 的备份文件
# sed -i "" 's/Atl/Dog/g' example.txt 如果这个字符串长度为0，就是说是个空串，那么不备份
errors=$(go list ./... | grep -v vendor | sed -e s=gitlab.local.com/golang/go-http/=./= | grep -v gitlab | xargs -n 1 go vet 2>&1)

# grep errors has go file
echo "${errors}" | grep -E "\.go|package"

# $? equal 0 has error
if [ $? -eq 0 ]; then
    echo '********** go vet check failed ~ **********'
    echo ${errors}
    exit 1
else
    echo '********** go vet check ok ~ **********'
    exit 0
fi
