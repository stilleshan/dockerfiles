# serverstatus

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/serverstatus](https://hub.docker.com/r/stilleshan/serverstatus)
> *docker image support for X86 and ARM*

## 简介
ServerStatus 在 GitHub 上已经有各种版本,本项目是基于 [ToyoDAdoubi/ServerStatus-Toyo](https://github.com/ToyoDAdoubi/ServerStatus-Toyo) 一键脚本项目,并整合网友制作的[美化主题](https://www.hostloc.com/thread-494384-1-1.html),打包为 docker 镜像,方便一键安装和迁移.
> *docker image support for X86 and ARM*

![screenshot](./screenshot.jpg)

## 更新
**2021-06-06** 更新 docker 镜像同时支持 X86 和 ARM

## 部署
### 创建配置文件
推荐直接 **git clone** 本仓库到服务器 **/root** 目录内.或者手动下载仓库中的 **config.json** 至 **/root/ServerStatus/** 目录内.
```shell
cd /root
git clone https://github.com/stilleshan/ServerStatus.git
```

### 启动容器
示例 **8888** 端口为 Web 访问端口, **35601** 为客户端通信端口,根据需求自行修改映射端口.注意防火墙需放行此端口.
```shell
docker run -d --name=serverstatus --restart=always -p 8888:80 -p 35601:35601 -v ~/ServerStatus/config.json:/ServerStatus/server/config.json stilleshan/serverstatus
```

### 访问地址
```
http://服务器IP:8888
```
> 使用域名和 HTTPS 协议可配置 Nginx 反向代理


### 配置
**config.json** 为服务器端配置文件,默认已经添加示例配置,可以根据示例格式修改,删除或者增加服务器.修改完毕后重启容器.
```shell
docker restart serverstatus
```

### 自定义前端页面
上述默认启动命令没有挂载 web 目录,如需自定义修改前端页面,需 **git clone** 本仓库到服务器 **/root** 目录内,执行以下命令挂载 web 目录.
```shell
docker run -d --name=serverstatus --restart=always -p 8888:80 -p 35601:35601 -v ~/ServerStatus/config.json:/ServerStatus/server/config.json -v ~/ServerStatus/web:/usr/share/nginx/html stilleshan/serverstatus
```

### 客户端
客户端需获取客户端服务器的运行情况,建议使用一键脚本安装.可参考图文教程: [服务器安装ServerStatus监控探针教程](https://www.ioiox.com/archives/27.html)
```shell
wget -N --no-check-certificate https://raw.githubusercontent.com/stilleshan/ServerStatus/master/status.sh && chmod +x status.sh && bash status.sh c
```

## 相关链接
- GitHub [stilleshan/ServerStatus](https://github.com/stilleshan/ServerStatus)
- Docker [stilleshan/serverstatus](https://hub.docker.com/r/stilleshan/serverstatus)
- 原版项目Github [ToyoDAdoubi/ServerStatus-Toyo](https://github.com/ToyoDAdoubi/ServerStatus-Toyo)
- Dockerfile参考 [cppla/ServerStatus](https://github.com/cppla/ServerStatus)

