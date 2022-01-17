FROM golang as builder
WORKDIR /builder
ADD . .
RUN set -ex \
    && curl -sSL https://api.github.com/repos/honwen/aliyun-ddns-cli/commits/master | sed -n '{/sha/p; /date/p;}' | sed 's/.* \"//g' | cut -c1-10 | tr '[:lower:]' '[:upper:]' | sed 'N;s/\n/@/g' | head -n1 | tee .version \
    && go mod vendor \
    && GOOS=linux GOARCH=amd64 VERSION=$(cat .version) make release


FROM alpine
LABEL MAINTAINER honwen <https://github.com/honwen>

# /usr/bin/aliyun-ddns-cli
COPY --from=builder /builder/build/linux-amd64/aliddns /usr/bin/aliyun-ddns-cli

ENV AKID=1234567890 \
    AKSCT=abcdefghijklmn \
    DOMAIN=ddns.example.win \
    IPAPI=[IPAPI-GROUP] \
    REDO=555r \
    TTL=600

CMD aliyun-ddns-cli \
    --ipapi ${IPAPI} \
    ${IPV6:+-6} \
    auto-update \
    --domain ${DOMAIN} \
    --redo ${REDO} \
    --ttl ${TTL}
