FROM golang:1.20-alpine AS build
ARG TARGETARCH
RUN apk add --update git
RUN git clone https://github.com/CareyWang/MyUrls /app
WORKDIR /app
RUN go env -w GO111MODULE="on" && go env -w GOPROXY="https://goproxy.cn,direct"
RUN go mod tidy 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o myurls main.go 

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/myurls ./
COPY public/* ./public/
COPY start.sh ./

EXPOSE 8002
CMD [ "sh", "-c", "/app/start.sh" ]
