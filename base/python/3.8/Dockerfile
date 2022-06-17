FROM python:3.8
LABEL maintainer="Stille <stille@ioiox.com>"

RUN apt-get update
RUN apt-get install -y zip vim
RUN pip install requests bs4 asyncio aiohttp lxml

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && apt-get install tzdata \
    && apt-get clean \
    && apt-get autoclean
