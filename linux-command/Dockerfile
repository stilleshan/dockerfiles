FROM node:alpine3.18 as builder
LABEL maintainer="Stille <stille@ioiox.com>"

WORKDIR /

RUN apk add --no-cache git
RUN git clone https://github.com/jaywcjlove/linux-command --depth=1 \
    && cd linux-command \
    && npm install \
    && npm run build \
    && npm run dash

FROM nginx:stable-alpine

RUN apk add --no-cache ca-certificates
COPY --from=builder /linux-command/.deploy /usr/share/nginx/html/

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
