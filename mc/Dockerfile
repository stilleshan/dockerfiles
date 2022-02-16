FROM alpine
LABEL maintainer="Stille <stille@ioiox.com>"

ENV TZ=Asia/Shanghai
RUN apk add --no-cache tzdata \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone
RUN if [ "$(uname -m)" = "x86_64" ]; then export PLATFORM=amd64 ; else if [ "$(uname -m)" = "aarch64" ]; then export PLATFORM=arm64 ; fi fi \
    && wget https://dl.minio.io/client/mc/release/linux-${PLATFORM}/mc -O /usr/local/bin/mc \
    && chmod +x /usr/local/bin/mc

CMD [ "/bin/sh" "-c" "mc" ]
