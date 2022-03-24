FROM alpine:3.13.5
LABEL maintainer="Amin Vakil <info@aminvakil.com>"

ENV VERSION 1.1.6

RUN apk add --no-cache bash

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN buildDeps=( \
		curl \
		g++ \
		gnutls-dev \
		gpgme \
		libev-dev \
		libnl3-dev \
		libseccomp-dev \
		linux-headers \
		linux-pam-dev \
		lz4-dev \
		make \
		readline-dev \
		tar \
		xz \
	); \
	set -x \
	&& apk add --update --virtual .build-deps "${buildDeps[@]}" \
	&& curl -SL --connect-timeout 8 --max-time 120 --retry 128 --retry-delay 5 "ftp://ftp.infradead.org/pub/ocserv/ocserv-$VERSION.tar.xz" -o ocserv.tar.xz \
	&& curl -SL --connect-timeout 8 --max-time 120 --retry 128 --retry-delay 5 "ftp://ftp.infradead.org/pub/ocserv/ocserv-$VERSION.tar.xz.sig" -o ocserv.tar.xz.sig \
	&& gpg --keyserver keys.gnupg.net --recv-key 96865171 \
	&& gpg --verify ocserv.tar.xz.sig \
	&& mkdir -p /usr/src/ocserv \
	&& tar -xf ocserv.tar.xz -C /usr/src/ocserv --strip-components=1 \
	&& rm ocserv.tar.xz* \
	&& cd /usr/src/ocserv \
	&& ./configure \
	&& make \
	&& make install \
	&& mkdir -p /etc/ocserv \
	&& cp /usr/src/ocserv/doc/sample.config /etc/ocserv/ocserv.conf \
	&& cd / \
	&& rm -fr /usr/src/ocserv \
	&& runDeps="$( \
		scanelf --needed --nobanner /usr/local/sbin/ocserv \
			| awk '{ gsub(/,/, "\nso:", $2); print "so:" $2 }' \
			| xargs -r apk info --installed \
			| sort -u \
		)" \
	&& readarray runDepsArr <<< "$runDeps" \
	&& apk add --virtual .run-deps "${runDepsArr[@]}" gnutls-utils iptables libnl3 readline libseccomp-dev lz4-dev \
	&& apk del .build-deps \
	&& rm -rf /var/cache/apk/*

# Setup config
COPY routes.txt /tmp/
RUN set -x \
	&& sed -i 's/\.\/sample\.passwd/\/etc\/ocserv\/ocpasswd/' /etc/ocserv/ocserv.conf \
	&& sed -i 's/\(max-same-clients = \)2/\110/' /etc/ocserv/ocserv.conf \
	&& sed -i 's/\.\.\/tests/\/etc\/ocserv/' /etc/ocserv/ocserv.conf \
	&& sed -i 's/#\(compression.*\)/\1/' /etc/ocserv/ocserv.conf \
	&& sed -i '/^ipv4-network = /{s/192.168.1.0/192.168.99.0/}' /etc/ocserv/ocserv.conf \
	&& sed -i 's/192.168.1.2/8.8.8.8/' /etc/ocserv/ocserv.conf \
	&& sed -i 's/^route/#route/' /etc/ocserv/ocserv.conf \
	&& sed -i 's/^no-route/#no-route/' /etc/ocserv/ocserv.conf \
	&& sed -i '/\[vhost:www.example.com\]/,$d' /etc/ocserv/ocserv.conf \
	&& sed -i '/^cookie-timeout = /{s/300/3600/}' /etc/ocserv/ocserv.conf \
	&& sed -i 's/^isolate-workers/#isolate-workers/' /etc/ocserv/ocserv.conf \
	&& cat /tmp/routes.txt >> /etc/ocserv/ocserv.conf \
	&& rm -rf /tmp/routes.txt

WORKDIR /etc/ocserv

COPY docker-entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 443
CMD ["ocserv", "-c", "/etc/ocserv/ocserv.conf", "-f"]
