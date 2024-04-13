FROM node:14-alpine AS build
LABEL maintainer="Stille <stille@ioiox.com>"

ENV SUBCONVERTER_VERSION=v0.9.0
WORKDIR /
RUN apk add --no-cache bash git curl zip
RUN if [ "$(uname -m)" = "x86_64" ]; then export PLATFORM=linux64 ; else if [ "$(uname -m)" = "aarch64" ]; then export PLATFORM=aarch64 ; fi fi \
  && wget https://github.com/tindy2013/subconverter/releases/download/${SUBCONVERTER_VERSION}/subconverter_${PLATFORM}.tar.gz \
  && tar xzf subconverter_${PLATFORM}.tar.gz
RUN git clone https://github.com/stilleshan/subweb app
RUN cd /app && npm install && npm run build

FROM nginx:1.16-alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY --from=build /subconverter /base
COPY . /app
EXPOSE 80
CMD [ "sh", "-c", "/app/start.sh" ]
