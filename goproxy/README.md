# goproxy for docker

GitHub [stilleshan/dockerfile](https://github.com/stilleshan/dockerfile)  
Docker [stilleshan/goproxy](https://hub.docker.com/r/stilleshan/goproxy)
> *docker image support for X86 and ARM*

## docker 启动
### 单个代理服务
启动一个端口为`11111`的 http 代理
```shell
docker run -d --name goproxy --restart always --network host stilleshan/goproxy /proxy http -p :11111
```

启动一个端口为`22222`的 socks5 代理
```shell
docker run -d --name goproxy --restart always --network host stilleshan/goproxy /proxy socks -p :22222
```

### 多个代理服务
单个容器同时启动`http`和`socks5`代理
```shell
docker run -d --name goproxy --restart always --network host stilleshan/goproxy sh -c '/proxy http -p :11111 & /proxy socks -p :22222'
```
> *更多服务参考 [snail007/goproxy](https://github.com/snail007/goproxy) 自行添加命令,每项服务命令之间使用 **&** 符号分隔,最后一项服务命令末尾无需 **&** 符号.*

## docker compose 启动
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfile/main/goproxy/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```
> *更多服务参考 [snail007/goproxy](https://github.com/snail007/goproxy) 自行添加命令,每项服务命令之间使用 **&** 符号分隔,最后一项服务命令末尾无需 **&** 符号.*
