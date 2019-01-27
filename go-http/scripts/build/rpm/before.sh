#!/bin/bash

# RPM_INSTALL_PREFIX是内置变量 在构建rpm的时候通过--prefix指定了
if [ -z "${RPM_INSTALL_PREFIX}" ];then
    echo '${RPM_INSTALL_PREFIX} can not be blank'
    exit 1
elif [ "${RPM_INSTALL_PREFIX}" = "/" ];then
    echo '${RPM_INSTALL_PREFIX} can not be /'
    exit 1
fi

# 如果已经有启动实例，需要先安全关闭
if [ -f "${RPM_INSTALL_PREFIX}/bin/stop.sh" ];then
    sh ${RPM_INSTALL_PREFIX}/bin/stop.sh
fi
