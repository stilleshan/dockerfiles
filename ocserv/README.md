# ocserv

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/ocserv](https://hub.docker.com/r/stilleshan/ocserv)
> *docker image support for X86 and ARM*

## 简介
基于 [openconnect/ocserv](https://github.com/openconnect/ocserv) 项目的 docker 镜像.

## 更新
- **2022-03-24** 更新`1.1.6`版 docker 镜像.
- **2021-06-09** 更新`1.1.2`版 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d \
    --name=ocserv \
    --restart=always \
    --sysctl=net.ipv4.ip_forward=1 \
    --cap-add=NET_ADMIN \
    --security-opt=no-new-privileges \
    -p 443:443 \
    -p 443:443/udp \
    stilleshan/ocserv
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/ocserv/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
GitHub [aminvakil/docker-ocserv](https://github.com/aminvakil/docker-ocserv)

