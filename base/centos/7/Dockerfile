FROM centos:7
LABEL maintainer="Stille <stille@ioiox.com>"

RUN yum -y install wget git zip

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone
