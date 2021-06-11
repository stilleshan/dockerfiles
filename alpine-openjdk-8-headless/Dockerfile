FROM alpine:latest
MAINTAINER Stille <stille@ioiox.com>

ENV VERSION 2021-06-11
RUN echo "http://dl-4.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
    && apk add --update openjdk8-jre-base \
    && rm -rf /var/cache/apk/*

