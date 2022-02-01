FROM python:3.7-alpine  
LABEL maintainer="Stille <stille@ioiox.com>"

WORKDIR /

ADD . /

RUN pip install --no-cache-dir tinify

ENTRYPOINT ["/bin/sh", "-c", "/tinypng.sh"]
