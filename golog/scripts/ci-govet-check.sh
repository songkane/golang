# Go语言源码中静态错误的简单工具
#!/bin/bash

echo '#### go vet checking ####'
go vet ./...
if [ $? -ne 0 ]; then
  echo 'go vet checking failed'
  exit 1
else
  echo 'go vet checking ok'
fi

exit 0
