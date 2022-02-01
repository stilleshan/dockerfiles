ARG PHP_VERSION=7.4-apache
FROM alpine as downloader

ENV VERSION 1.3.0
ARG CHEVERETO_VERSION=1.3.0
RUN apk add --no-cache curl && \
    curl -sS -o /tmp/chevereto.zip -L "https://github.com/Chevereto/Chevereto-Free/archive/${CHEVERETO_VERSION}.zip" && \
    mkdir -p /extracted && \
    cd /extracted && \
    unzip /tmp/chevereto.zip  && \
    mv "chevereto-free-${CHEVERETO_VERSION}/" Chevereto/
COPY settings.php /extracted/Chevereto/app/settings.php

FROM php:$PHP_VERSION

# Install required packages and configure plugins + mods for Chevereto
RUN apt-get update && apt-get install -y \
        libgd-dev \
        libzip-dev && \
    bash -c 'if [[ $PHP_VERSION == 7.4.* ]]; then \
      docker-php-ext-configure gd --with-freetype=/usr/include/ --with-jpeg=/usr/include/; \
    else \
      docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/; \
    fi' && \
    docker-php-ext-install \
        exif \
        gd \
        mysqli \
        pdo \
        pdo_mysql \
        zip && \
    a2enmod rewrite

# Download installer script
COPY --from=downloader --chown=33:33 /extracted/Chevereto /var/www/html

# Expose the image directory as a volume
VOLUME /var/www/html/images

# DB connection environment variables
ENV CHEVERETO_DB_HOST=db CHEVERETO_DB_USERNAME=chevereto CHEVERETO_DB_PASSWORD=chevereto CHEVERETO_DB_NAME=chevereto CHEVERETO_DB_PREFIX=chv_ CHEVERETO_DB_PORT=3306
ARG BUILD_DATE
ARG CHEVERETO_VERSION=1.2.2

# Set all required labels, we set it here to make sure the file is as reusable as possible
LABEL org.label-schema.url="https://github.com/tanmng/docker-chevereto" \
      org.label-schema.name="Chevereto Free" \
      org.label-schema.license="Apache-2.0" \
      org.label-schema.version="${CHEVERETO_VERSION}" \
      org.label-schema.vcs-url="https://github.com/tanmng/docker-chevereto" \
      maintainer="Tan Nguyen <tan.mng90@gmail.com>" \
      build_signature="Chevereto free version ${CHEVERETO_VERSION}; built on ${BUILD_DATE}; Using PHP version ${PHP_VERSION}"
