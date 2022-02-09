FROM python:3.8-alpine3.15
LABEL maintainer="Stille <stille@ioiox.com>"

RUN apk add --no-cache bash git curl zip

ENV TZ=Asia/Shanghai
RUN apk add --no-cache tzdata \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone
