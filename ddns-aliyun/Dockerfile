FROM golang:alpine as builder
ENV CGO_ENABLED=0 \
    GO111MODULE=on

RUN apk add --update git curl
ADD . $GOPATH/src/github.com/honwen/aliyun-ddns-cli
RUN set -ex \
    && cd $GOPATH/src/github.com/honwen/aliyun-ddns-cli \
    && go build -ldflags "-X main.VersionString=$(curl -sSL https://api.github.com/repos/honwen/aliyun-ddns-cli/commits/master | \
            sed -n '{/sha/p; /date/p;}' | sed 's/.* \"//g' | cut -c1-10 | tr '[:lower:]' '[:upper:]' | sed 'N;s/\n/@/g' | head -1)" . \
    && mv aliyun-ddns-cli $GOPATH/bin/


FROM alpine
LABEL MAINTAINER honwen <https://github.com/honwen>

# /usr/bin/aliyun-ddns-cli
COPY --from=builder /go/bin /usr/bin

ENV AKID=1234567890 \
    AKSCT=abcdefghijklmn \
    DOMAIN=ddns.example.win \
    IPAPI=[IPAPI-GROUP] \
    REDO=0

CMD aliyun-ddns-cli \
    --ipapi ${IPAPI} \
    ${IPV6:+-6} \
    auto-update \
    --domain ${DOMAIN} \
    --redo ${REDO}
