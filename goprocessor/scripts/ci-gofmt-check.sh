# Go语言源码格式化 如果有Go文件不符合gofmt格式化直接报错
#!/bin/bash

echo '#### gofmt checking ####'
illegal_format_files=`go list ./... | grep -v vendor | sed -e s=gitlab.local.com/golang/goprocessor/=./= | xargs -n 1 gofmt -l`
echo "${illegal_format_files}" |grep "\.go" 1>&2 >/dev/null
if [ $? -eq 0 ]; then
  # found something
  echo 'gofmt checking failed'
  echo "###### Illegal Format Files ######"
  echo "${illegal_format_files}"
  echo "##############################"
  go list ./... | grep -v vendor | sed -e s=gitlab.local.com/golang/goprocessor/=./= | xargs -n 1 gofmt -d
  exit 1
else
  echo 'gofmt checking ok'
fi

exit 0
