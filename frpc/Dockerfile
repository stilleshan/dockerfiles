FROM alpine:3.8
MAINTAINER Stille <stille@ioiox.com>

ENV FRP_VERSION 0.37.1
WORKDIR /

RUN set -xe && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

RUN set -x && \
	if [ "$(uname -m)" = "x86_64" ]; then export PLATFORM=amd64 ; else if [ "$(uname -m)" = "aarch64" ]; then export PLATFORM=arm64 ; fi fi && \
	wget --no-check-certificate https://github.com/fatedier/frp/releases/download/v${FRP_VERSION}/frp_${FRP_VERSION}_linux_${PLATFORM}.tar.gz && \ 
	tar xzf frp_${FRP_VERSION}_linux_${PLATFORM}.tar.gz && \
	cd frp_${FRP_VERSION}_linux_${PLATFORM} && \
	mkdir /frp && \
	mv frpc frpc.ini /frp && \
	cd .. && \
	rm -rf *.tar.gz frp_${FRP_VERSION}_linux_${PLATFORM}

VOLUME /frp

CMD /frp/frpc -c /frp/frpc.ini
