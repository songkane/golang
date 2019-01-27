#!/bin/bash

# RPM_INSTALL_PREFIX是内置变量 在构建rpm的时候通过--prefix指定了
chmod +x ${RPM_INSTALL_PREFIX}/bin/*
sh ${RPM_INSTALL_PREFIX}/bin/start.sh