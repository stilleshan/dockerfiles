# vnstat-dashboard-kshitiz-b 

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/vnstat-dashboard-kshitiz-b](https://hub.docker.com/r/stilleshan/vnstat-dashboard-kshitiz-b)
> *docker image support for X86 and ARM*

## 简介
基于 [Kshitiz-b/vnstat-dashboard](https://github.com/Kshitiz-b/vnstat-dashboard) 网络流量监控软件前端面板项目的 docker 镜像备份.

## 部署
### docker
需要服务器已安装`vnStat`软件,详情访问 [vnStat](https://humdi.net/vnstat/) 或 [vergoh/vnstat](https://github.com/vergoh/vnstat) .
```shell
docker run -d \
    --name=vnstat-dashboard \
    --privileged \
    --restart=always \
    -p 12345:8050 \
    -v /var/lib/vnstat:/var/lib/vnstat \
    -e TZ=Asia/Shanghai \
    stilleshan/vnstat-dashboard-kshitiz-b
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/vnstat-dashboard-kshitiz-b/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
参考以下原项目备份镜像:
- GitHub [Kshitiz-b/vnstat-dashboard](https://github.com/Kshitiz-b/vnstat-dashboard)
