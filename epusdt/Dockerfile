FROM golang:alpine AS build

WORKDIR /
ENV CGO_ENABLED=0
RUN apk add git
RUN git clone https://github.com/assimon/epusdt
RUN mv epusdt/src /

WORKDIR /src
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/epusdt .

FROM alpine
LABEL maintainer="stille@ioiox.com"

WORKDIR /app
COPY --from=build /out/epusdt /app/epusdt
COPY --from=build /src/static /app/static
RUN set -xe \
    && chmod +x /app/epusdt

CMD /app/epusdt http start
