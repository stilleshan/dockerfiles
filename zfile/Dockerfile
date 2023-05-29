FROM ibm-semeru-runtimes:open-8-jre
LABEL maintainer="Stille <stille@ioiox.com>"

ENV VERSION 4.1.5

WORKDIR /root

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
RUN curl -L -o app.jar https://github.com/zfile-dev/zfile/releases/download/${VERSION}/zfile-${VERSION}.jar

EXPOSE 8080

ENTRYPOINT java $JAVA_OPTS -Xshareclasses -Xquickstart -jar /root/app.jar
