# Go单元测试检查
#!/bin/bash

echo '#### go test checking ####'
go list ./... | grep -v vendor | sed -e s=gitlab.local.com/golang/golog/=./= | xargs -n 1 go test
if [ $? -ne 0 ]; then
    echo 'go test checking failed'
    go list ./... | grep -v vendor | sed -e s=gitlab.local.com/golang/golog/=./= | xargs -n 1 go test
    exit 1
else
    echo 'go test checking ok'
fi

exit 0
