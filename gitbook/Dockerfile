FROM node:10.22.0-alpine3.9
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 3.2.3

RUN npm install gitbook-cli -g
RUN gitbook fetch $VERSION

EXPOSE 4000
VOLUME /gitbook

WORKDIR /gitbook

CMD ["gitbook", "--help"]
