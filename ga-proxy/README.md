# ga-proxy

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/ga-proxy](https://hub.docker.com/r/stilleshan/ga-proxy)
> *docker image support for X86 and ARM*

## 简介
基于 [giuem/ga-proxy](https://github.com/giuem/ga-proxy) 项目的 docker 镜像.

## 更新
**2021-06-09** 更新`1.2.0`版 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d --name=ga-proxy --restart=always -p 12345:80 stilleshan/ga-proxy
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/ga-proxy/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
GitHub [giuem/ga-proxy](https://github.com/giuem/ga-proxy)

