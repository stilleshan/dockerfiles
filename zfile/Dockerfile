FROM stilleshan/alpine-openjdk-8-headless:2022-02-01
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 3.2.1

WORKDIR /root

RUN wget https://github.com/zhaojun1998/zfile/releases/download/${VERSION}/zfile-${VERSION}.war \
    && mkdir zfile && unzip zfile-${VERSION}.war -d zfile && rm -rf zfile-${VERSION}.war \
    && chmod +x ~/zfile/bin/*.sh

EXPOSE 8080

CMD sh ./zfile/bin/start.sh && tail -f /dev/null
