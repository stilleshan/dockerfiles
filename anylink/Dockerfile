# web
FROM node:16.17.1-alpine3.15 as builder_node
ENV VERSION 0.9.4
WORKDIR /web
COPY ./web /web
RUN yarn install \
    && yarn run build \
    && ls /web/ui

# server
FROM golang:1.19-alpine as builder_golang
#TODO 本地打包时使用镜像
ENV GOPROXY=https://goproxy.io
ENV GOOS=linux
WORKDIR /anylink
COPY . /anylink
COPY --from=builder_node /web/ui  /anylink/server/ui

#TODO 本地打包时使用镜像
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache git gcc musl-dev curl
RUN cd /anylink/server;go mod tidy;go build -o anylink -ldflags "-X main.CommitId=$(curl -s "https://api.github.com/repos/bjdgyc/anylink/git/refs/tags/v${VERSION}" | grep sha | head -n 1 | cut -d '"' -f 4)" \
    && /anylink/server/anylink tool -v

# anylink
FROM alpine
LABEL maintainer="github.com/bjdgyc"

#ENV IPV4_CIDR="192.168.10.0/24"

WORKDIR /app
COPY --from=builder_golang /anylink/server/anylink  /app/
COPY docker_entrypoint.sh  /app/

#COPY ./server/bridge-init.sh /app/
COPY ./server/conf  /app/conf
COPY ./LICENSE  /app/LICENSE


#TODO 本地打包时使用镜像
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache bash iptables \
    && chmod +x /app/docker_entrypoint.sh \
    && ls /app

EXPOSE 443 8800

#CMD ["/app/anylink"]
ENTRYPOINT ["/app/docker_entrypoint.sh"]