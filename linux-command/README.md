# linux-command
GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/linux-command](https://hub.docker.com/r/stilleshan/linux-command)
> *docker image support for X86 and ARM*

## 简介
基于 [jaywcjlove/linux-command](https://github.com/jaywcjlove/linux-command) 项目的 docker 镜像. 方便快速,离线部署使用.

## 示例
[https://linux.ioiox.com](https://linux.ioiox.com)

## 更新
由于该项目比较活跃和热门,本镜像将定时自动更新最新版.

## 部署
### docker
```shell
docker run  -d --name=linux-command --restart=always -p 8888:80 stilleshan/linux-command
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/linux-command/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 链接
- [jaywcjlove/linux-command](https://github.com/jaywcjlove/linux-command)
- [https://linux.ioiox.com](https://linux.ioiox.com)
