FROM tindy2013/subconverter
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 0.7.2

WORKDIR /base
RUN apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

COPY groups.txt rulesets.txt /base/snippets/
COPY pref.yml /base

EXPOSE 25500

WORKDIR /base

CMD subconverter
