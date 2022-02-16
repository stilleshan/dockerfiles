FROM python:3.7-slim
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 1.8.6.24
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN pip install --upgrade --no-cache-dir coscmd
