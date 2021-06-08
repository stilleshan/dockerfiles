# tinyproxy for docker

GitHub [stilleshan/dockerfile](https://github.com/stilleshan/dockerfile)  
Docker [stilleshan/tinyproxy](https://hub.docker.com/r/stilleshan/tinyproxy)
> *docker image support for X86 and ARM*

## docker 启动
```shell
docker run -d --name tinyproxy --restart always -p 8888:8888 stilleshan/tinyproxy
```

## docker compose 启动
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfile/main/tinyproxy/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 使用示例
```shell
curl -x https://IP:8888 https://ifconfig.co
```

## BasicAuth
Tinyproxy 支持 BasicAuth 身份验证,需启动一次容器并执行以下命令将`tinyproxy.conf`拷贝至宿主机,并删除该容器.
```shell
docker cp tinyproxy:/etc/tinyproxy/tinyproxy.conf .
```
配置`BasicAuth user password`参数并将`tinyproxy.conf`挂载至`/etc/tinyproxy/tinyproxy.conf`后再次启动容器.
```shell
docker run -d --name tinyproxy --restart always -p 8888:8888 -v /path/tinyproxy.conf:/etc/tinyproxy/tinyproxy.conf stilleshan/tinyproxy
```

### 使用示例
```shell
curl -x https://user:password@IP:8888 https://ifconfig.co
```
