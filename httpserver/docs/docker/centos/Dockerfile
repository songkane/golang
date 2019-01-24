# 基于centos7定制的镜像
# 基础镜像
FROM centos:7

# MAINTAINER 维护者信息
MAINTAINER chenguolin

# RUN 执行以下命令
# 1). 更新yum
# 2). 安装fpm依赖
# 3). 安装fpm
# 4). 删除无效数据
RUN yum update -y \
    && yum install -y fpm wget gcc ruby-devel make rpm-build rubygems \
    && gem install --no-ri --no-rdoc fpm \
    && yum clean all \
    && rm -rf /usr/local/src/*

# 安装golang
RUN wget https://dl.google.com/go/go1.10.linux-amd64.tar.gz \
    && tar zxvf go1.10.linux-amd64.tar.gz -C /usr/local/ \
    && mkdir -p /go \
    && echo "export GOROOT=/usr/local/go" >> /etc/profile \
    && echo "export GOBIN=/usr/local/go/bin" >> /etc/profile \
    && echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile \
    && echo "export GOPATH=/go" >> /etc/profile \
    && echo "source /etc/profile" >> /root/.bashrc

# 设置GOPATH
ENV GOPATH /go

