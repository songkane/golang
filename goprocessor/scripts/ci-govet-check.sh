# Go语言源码中静态错误检查，如果有Go文件存在静态错误直接报错
#!/bin/bash

echo '#### go vet checking ####'
go list ./... | grep -v vendor | sed -e s=gitlab.local.com/golang/golog/=./= | xargs -n 1 go vet
if [ $? -ne 0 ]; then
  echo 'go vet checking failed'
  exit 1
else
  echo 'go vet checking ok'
fi

exit 0
