# vnstat-dashboard

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/vnstat-dashboard](https://hub.docker.com/r/stilleshan/vnstat-dashboard)
> *docker image support for X86 and ARM*

## 简介
基于 [vnStat](https://humdi.net/vnstat/) 网络流量监控软件前端面板项目的 docker 镜像.

## 更新
**2021-06-11** 更新 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
需要服务器已安装`vnStat`软件,详情访问 [vnStat](https://humdi.net/vnstat/) 或 [vergoh/vnstat](https://github.com/vergoh/vnstat) .
```shell
docker run -d \
    --name=vnstat-dashboard \
    --restart=always \
    -p 12345:80 \
    -v /usr/bin/vnstat:/usr/bin/vnstat \
    -v /var/lib/vnstat:/var/lib/vnstat \
    -e TZ=Asia/Shanghai \
    stilleshan/vnstat-dashboard
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/vnstat-dashboard/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
参考以下项目修复 bug 重新构建:
- GitHub [tomangert/vnstat-dashboard](https://github.com/tomangert/vnstat-dashboard)
- GitHub [alexandermarston/vnstat-dashboard](https://github.com/alexandermarston/vnstat-dashboard)

