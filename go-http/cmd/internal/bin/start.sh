# 启动脚本
#!/bin/bash

# 相关变量
app_dir="/www/go-http-internal"
log_dir="${app_dir}/logs"
# TODO(@cgl) 自行更改app_name
app_name="go-http-internal"
pid_file="${app_dir}/${app_name}.pid"

# 先判断是否已经有进程在跑 如果有则不启动
count=`ps axu | grep ${app_name} | grep "${app_dir}/logs/"`
if [ "$count" -gt 0 ];then
    echo "${app_name} has started ~"
fi

# 启动进程
mkdir -p ${app_dir}/logs
nohup ./${app_name} -log_dir ${log_dir} >>${log_dir}/stdout.log 2>>${log_dir}/stderr.log &

# 输出启动成功信息
echo $!>${pid_file}
echo "${app_name} start successful ~"
