FROM node:lts-alpine
LABEL org.opencontainers.image.authors="stille@ioiox.com"

WORKDIR /

RUN npm i docsify-cli -g
RUN docsify init ./docs

EXPOSE 3000

CMD docsify serve docs
