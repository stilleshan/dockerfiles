FROM alpine
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 14.7

RUN set -xe && \
    UNAME=$(uname -m) && if [ "$UNAME" = "x86_64" ]; then export PLATFORM=amd64 ; else export PLATFORM=arm64-v8 ; fi && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
    wget https://github.com/snail007/goproxy/releases/download/v${VERSION}/proxy-linux-${PLATFORM}.tar.gz && \
    tar -xvf proxy-linux-${PLATFORM}.tar.gz

CMD ["/proxy"]
