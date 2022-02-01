FROM alpine:latest
LABEL maintainer="Stille <stille@ioiox.com>"

RUN echo "http://dl-4.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
    && apk add --update openjdk8-jre-base \
    && rm -rf /var/cache/apk/*

