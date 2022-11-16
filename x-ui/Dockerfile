FROM golang:latest AS builder
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 0.3.3.18

WORKDIR /root
RUN git clone https://github.com/FranzKafkaYu/x-ui
RUN cd x-ui && go build main.go

FROM debian:11-slim
RUN apt-get update \
    && apt-get install -y --no-install-recommends -y ca-certificates \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
ENV TZ=Asia/Shanghai
WORKDIR /root
COPY --from=builder /root/x-ui/main /root/x-ui
COPY --from=builder /root/x-ui/bin/. /root/bin/.
VOLUME [ "/etc/x-ui" ]
CMD [ "./x-ui" ]
