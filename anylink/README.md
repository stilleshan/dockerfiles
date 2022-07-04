# anylink

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/anylink](https://hub.docker.com/r/stilleshan/anylink)
> *docker image support for X86 and ARM*

## 简介
基于 [bjdgyc/anylink](https://github.com/bjdgyc/anylink) 项目的 docker 镜像.

## 更新
- **2022-07-04** 更新`0.8.1`版 docker 镜像.
- **2022-04-07** 更新`0.7.4`版 docker 镜像.
- **2022-02-16** 更新`0.7.3`版 docker 镜像.
- **2021-12-31** 更新`0.7.2`版 docker 镜像.
- **2021-12-29** 更新`0.7.1`版 docker 镜像.
- **2021-08-26** 更新`0.6.2`版 docker 镜像.
- **2021-08-02** 更新`0.5.1`版 docker 镜像.
- **2021-07-05** 更新`0.4.2`版 docker 镜像.
- **2021-06-09** 更新`0.3.3`版 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d \
    --name=anylink \
    --restart=always \
    --privileged=true \
    -p 443:443 \
    -p 8800:8800 \
    stilleshan/anylink
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/anylink/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

### 配置
```shell
https://ip:8800
```
关于初始配置,修改密码,网络等更多信息请参考: GitHub [bjdgyc/anylink](https://github.com/bjdgyc/anylink)


## 参考
- GitHub [bjdgyc/anylink](https://github.com/bjdgyc/anylink)
- [开源企业级远程办公 VPN 软件 AnyLink 的 docker 部署及使用心得](https://www.ioiox.com/archives/128.html)
- [群晖NAS高级服务 - docker 部署 AnyLink 企业级远程办公 VPN 服务](https://www.ioiox.com/archives/129.html)
