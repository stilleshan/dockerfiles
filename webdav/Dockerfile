FROM alpine:3.8
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 4.1.0

WORKDIR /

RUN set -xe && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

RUN set -x && \
	if [ "$(uname -m)" = "x86_64" ]; then export ARCH=amd64 ; else if [ "$(uname -m)" = "aarch64" ]; then export ARCH=arm64 ; fi fi && \
	wget --no-check-certificate https://github.com/hacdias/webdav/releases/download/v${VERSION}/linux-${ARCH}-webdav.tar.gz && \ 
	tar xvf linux-${ARCH}-webdav.tar.gz && \
	rm -rf *.tar.gz

COPY config.yml /

CMD /webdav -c /config.yml
