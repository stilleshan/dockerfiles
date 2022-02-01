FROM ubuntu:bionic-20200112 as builder
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 2.0
WORKDIR /

COPY . /
RUN apt-get update && \
    apt-get -y install wget && \
    /bin/bash -c '/bin/echo -e "1\n\nn\n" | ./status.sh' && \
    cp -rf /web /usr/local/ServerStatus/

FROM nginx:1.17.8
LABEL maintainer="Stille <stille@ioiox.com>"

COPY --from=builder /usr/local/ServerStatus/server /ServerStatus/server/
COPY --from=builder /usr/local/ServerStatus/web /usr/share/nginx/html/

EXPOSE 80 35601

CMD nohup sh -c '/etc/init.d/nginx start && /ServerStatus/server/sergate --config=/ServerStatus/server/config.json --web-dir=/usr/share/nginx/html'
