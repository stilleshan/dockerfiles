# tinyproxy

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/tinyproxy](https://hub.docker.com/r/stilleshan/tinyproxy)
> *docker image support for X86 and ARM*

## 简介
基于 [tinyproxy/tinyproxy](https://github.com/tinyproxy/tinyproxy) 项目的 docker 镜像.

## 更新
- **2022-04-08** 更新`1.11.0`版 docker 镜像.
- **2021-06-09** 更新`1.10.0`版 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d --name=tinyproxy --restart=always -p 8888:8888 stilleshan/tinyproxy
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/tinyproxy/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 使用
```shell
curl -x https://IP:8888 https://ifconfig.co
```

### BasicAuth
Tinyproxy 支持 BasicAuth 身份验证,需启动一次容器并执行以下命令将`tinyproxy.conf`拷贝至宿主机,并删除该容器.
```shell
docker cp tinyproxy:/etc/tinyproxy/tinyproxy.conf .
```
配置`BasicAuth user password`参数并将`tinyproxy.conf`挂载至`/etc/tinyproxy/tinyproxy.conf`后再次启动容器.
```shell
docker run -d --name=tinyproxy --restart=always -p 8888:8888 -v /path/tinyproxy.conf:/etc/tinyproxy/tinyproxy.conf stilleshan/tinyproxy
```

### BasicAuth 示例
```shell
curl -x https://user:password@IP:8888 https://ifconfig.co
```

## 参考
[轻量级 HTTP/HTTPS 代理软件 Tinyproxy docker 部署教程](https://www.ioiox.com/archives/130.html)

