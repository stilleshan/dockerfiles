FROM debian:stable-slim
LABEL maintainer="Stille <stille@ioiox.com>"

RUN apt-get update
RUN apt-get install -y wget curl zip git

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && apt-get install tzdata \
    && apt-get clean \
    && apt-get autoclean
