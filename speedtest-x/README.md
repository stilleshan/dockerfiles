# speedtest-x

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/speedtest-x](https://hub.docker.com/r/stilleshan/speedtest-x)
> *docker image support for X86 and ARM*

## 简介
基于 [BadApple9/speedtest-x](https://github.com/BadApple9/speedtest-x) 项目的 docker 镜像.

## 更新
**2021-06-14** 更新 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d --name=speedtest-x --restart=always -p 12345:80 stilleshan/speedtest-x
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/speedtest-x/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
GitHub [BadApple9/speedtest-x](https://github.com/BadApple9/speedtest-x) 

