FROM ubuntu:21.10
MAINTAINER Stille <stille@ioiox.com>

ENV VERSION 5.5.2

RUN apt -y update && apt -y install curl && \
    curl -O https://raw.githubusercontent.com/sprov065/v2-ui/master/install.sh && chmod +x install.sh&& ./install.sh

CMD [ "sh", "-c", "/usr/local/v2-ui/v2-ui"]

