FROM golang:alpine as builder

ENV VERSION 0.8.1

RUN apk add --no-cache make git
RUN git clone --depth 1 https://github.com/zu1k/nali /nali-src
WORKDIR /nali-src
RUN go mod download && \
    make docker && \
    mv ./bin/nali-docker /nali

FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /nali /
ENTRYPOINT ["/nali"]
