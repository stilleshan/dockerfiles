FROM ubuntu:21.10
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 0.3.1

COPY Shanghai /usr/share/zoneinfo/Asia/

RUN apt -y update && apt -y install curl && \
    curl -O https://raw.githubusercontent.com/sprov065/x-ui/master/install.sh && chmod +x install.sh&& ./install.sh

CMD [ "sh", "-c", "/usr/local/x-ui/x-ui"]
