FROM alpine
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 1.11.0

RUN set -xe && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

RUN set -xe \
    && apk add --no-cache tinyproxy \
    && sed -i -e '/^Allow /s/^/#/' \
              -e '/^ConnectPort /s/^/#/' \
              -e '/^#DisableViaHeader /s/^#//' \
              /etc/tinyproxy/tinyproxy.conf

VOLUME /etc/tinyproxy
EXPOSE 8888

CMD ["tinyproxy", "-d"]
