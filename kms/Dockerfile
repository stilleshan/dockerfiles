FROM alpine:latest AS builder

WORKDIR /root
ENV VERSION 2021-06-13

RUN apk add --no-cache git make build-base && \
    git clone --branch master --single-branch https://github.com/Wind4/vlmcsd.git && \
    cd vlmcsd/ && \
    make

FROM alpine:latest
MAINTAINER Stille <stille@ioiox.com>

COPY --from=builder /root/vlmcsd/bin/vlmcsd /usr/bin/vlmcsd

RUN set -xe && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

EXPOSE 1688

CMD [ "/usr/bin/vlmcsd", "-D", "-e" ]
