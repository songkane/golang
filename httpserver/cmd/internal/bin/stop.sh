# 停止脚本
#!/bin/bash

# 相关变量
app_dir="/www/httpserver-internal"
log_dir="${app_dir}/logs"
app_name="httpserver-internal"
pid_file="${app_dir}/${app_name}.pid"

# 如果没有pid文件认为未启动直接退出
if [ -z "${pidFile}" ];then
    echo "not found pid file exit now ~"
    exit 1
fi

# 获取pid
pid=`cat ${pid_file}`
echo "start to kill ${app_name} ..."

# kill进程
kill ${pid}

# 循环检测进程是否已退出
while [ 1 ];do
    ps -p ${pid}
    alive=`ps -p ${pid} | grep ${pid} | wc -l`
    if [ ${alive} -le 0 ];then
        break;
    fi
    sleep 2s
done

# 删除pid文件
rm -rf ${pid_file}
# 输出停止成功
echo "${app_name} stop successful ~"
