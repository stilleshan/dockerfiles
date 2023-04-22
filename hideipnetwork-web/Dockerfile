FROM ioiox/node:16-alpine
LABEL maintainer="Stille <stille@ioiox.com>"

WORKDIR /app

RUN git clone -b v2 https://github.com/Hideipnetwork/hideipnetwork-web.git . && npm install

EXPOSE 56559

CMD ["npm", "run", "start"]
