FROM henningn/selenium-standalone-firefox
USER root

# Install curl
RUN apt-get update -qqy \
  && apt-get -qqy install \
    curl

# install uuid
RUN apt-get update -qqy \
  && apt-get -qqy install \
    uuid

RUN apt-get update -qqy \
  && apt-get -qqy install \
    cmake

# install openssl
RUN apt-get update -qqy \
  && apt-get -qqy install \
    libssl-dev
    
# install libgit2
RUN apt-get update && wget https://github.com/libgit2/libgit2/archive/v0.27.0.tar.gz \
  && tar xzf v0.27.0.tar.gz \
  && cd libgit2-0.27.0/ \
  && cmake . \
  && make \
  && sudo make install

# Install Mysql
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -qqy \
  && apt-get -qqy install \
    mysql-server \
    libmysqlclient-dev

# Install git
RUN apt-get -qqy \
  install git

# Install nodejs v8
RUN curl -sL https://deb.nodesource.com/setup_8.x | bash -
RUN apt-get update -qqy \
  && apt-get -qqy install \
    nodejs

# Install yarn
RUN curl -o- -L https://yarnpkg.com/install.sh | bash -

# Update to firefox nightly
ARG FIREFOX_DOWNLOAD_URL=https://download.mozilla.org/?product=firefox-nightly-latest-ssl&lang=en-US&os=linux64
RUN wget --no-verbose -O /tmp/firefox.tar.bz2 $FIREFOX_DOWNLOAD_URL \
  && apt-get -y purge firefox \
  && rm -rf /opt/firefox \
  && tar -C /opt -xjf /tmp/firefox.tar.bz2 \
  && rm /tmp/firefox.tar.bz2 \
  && mv /opt/firefox /opt/firefox-nightly \
  && ln -fs /opt/firefox-nightly/firefox /usr/bin/firefox

# Install python
RUN apt-get update -qqy \
  && apt-get -qqy install \
    python-pip \
    python-dev \
    build-essential \
  && pip install --upgrade pip

# Install Tox
RUN pip install tox

ENV USER=seluser

WORKDIR /code

USER seluser

EXPOSE 5900
EXPOSE 4444
