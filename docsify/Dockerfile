FROM node:lts-alpine
LABEL org.opencontainers.image.authors="stille@ioiox.com"

WORKDIR /

ADD entrypoint.sh /

RUN npm i docsify-cli -g

EXPOSE 3000

ENTRYPOINT ["/bin/sh", "-c", "/entrypoint.sh"]
