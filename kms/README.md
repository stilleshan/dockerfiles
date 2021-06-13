## kms

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/kms](https://hub.docker.com/r/stilleshan/kms)
> *docker image support for X86 and ARM*

## 简介
基于 KMS 项目的 docker 镜像.

## 更新
**2021-06-13** 更新 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d --name=kms --restart=always -p 1688:1688 stilleshan/kms
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/kms/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

> **注意:** 需开放防火墙`1688`端口

## 参考
[Windows 与 Office 的下载安装及使用 KMS 密钥管理服务激活](https://www.ioiox.com/archives/107.html)

