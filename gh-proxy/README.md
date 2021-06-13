# gh-proxy

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/gh-proxy](https://hub.docker.com/r/stilleshan/gh-proxy)
> *docker image support for X86 and ARM*

## 简介
基于 [hunshcn/gh-proxy](https://github.com/hunshcn/gh-proxy) 项目的 docker 镜像.

## 更新
**2021-06-09** 更新`2.1`版 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d --name=gh-proxy --restart=always -p 12345:80 stilleshan/gh-proxy
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/gh-proxy/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
[ghproxy.com](https://ghproxy.com)

