# goproxy

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/goproxy](https://hub.docker.com/r/stilleshan/goproxy)
> *docker image support for X86 and ARM*

## 简介
基于 [snail007/goproxy](https://github.com/snail007/goproxy) 项目的 docker 镜像.

## 部署
### docker
#### 单个代理服务
启动一个端口为`11111`的 **http** 代理
```shell
docker run -d --name=goproxy --restart=always --network=host stilleshan/goproxy /proxy http -p :11111
```

启动一个端口为`22222`的 **socks5** 代理
```shell
docker run -d --name=goproxy --restart=always --network=host stilleshan/goproxy /proxy socks -p :22222
```

#### 多个代理服务
单个容器同时启动 **http** 和 **socks5** 代理
```shell
docker run -d --name=goproxy --restart=always --network=host stilleshan/goproxy sh -c '/proxy http -p :11111 & /proxy socks -p :22222'
```
> *自行添加命令,每项服务命令之间使用 **&** 符号分隔,最后一项服务命令末尾无需 **&** 符号.*

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/goproxy/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```
> *自行添加命令,每项服务命令之间使用 **&** 符号分隔,最后一项服务命令末尾无需 **&** 符号.*

## 参考
- GitHub [snail007/goproxy](https://github.com/snail007/goproxy)
- [轻量级高性能 HTTP/HTTPS SOCKS5 代理软件 goproxy docker 部署教程](https://www.ioiox.com/archives/131.html)

