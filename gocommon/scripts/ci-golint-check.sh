# Go语言源码编码规范检查
#!/bin/bash

echo '#### golint checking ####'
illegal_golint_files=`go list ./... | grep -v vendor | sed -e s=gitlab.local.com/golang/gocommon/=./= | xargs -n 1 golint`
echo "${illegal_golint_files}" |grep "\.go" >/dev/null 2>&1
if [ $? -eq 0 ]; then
  # found something
  echo 'golint checking failed'
  echo "###### Illegal Golint Files ######"
  echo "##############################"
  go list ./... | grep -v vendor | sed -e s=gitlab.local.com/golang/golog/=./= | xargs -n 1 golint
  exit 1
else
  echo 'golint checking ok'
fi

exit 0
