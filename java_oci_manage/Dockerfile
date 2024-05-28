FROM openjdk:20-ea-11-jdk-slim-bullseye
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 5.0.1

RUN apt-get update && apt-get install -y wget procps

WORKDIR /app
RUN wget -O gz_client_bot.tar.gz  https://github.com/semicons/java_oci_manage/releases/download/v${VERSION}/gz_client_bot.tar.gz \
  && tar -zxvf gz_client_bot.tar.gz --exclude=client_config \
  && tar -zxvf gz_client_bot.tar.gz --skip-old-files client_config \
  && rm -rf gz_client_bot.tar.gz \
  && chmod +x sh_client_bot.sh
