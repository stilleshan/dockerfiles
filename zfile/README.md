# zfile

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/zfile](https://hub.docker.com/r/stilleshan/zfile)
> *docker image support for X86 and ARM*

## 简介
基于 [zhaojun1998/zfile](https://github.com/zhaojun1998/zfile) 项目的 docker 镜像.

## 更新
- **2023-03-05** 更新`4.1.4`版 docker 镜像.
- **2022-11-27** 更新`4.1.3`版 docker 镜像.
- **2022-09-21** 更新`4.1.2`版 docker 镜像.
- **2022-08-30** 更新`4.1.1`版 docker 镜像.
- **2022-08-28** 更新`4.1.0`版 docker 镜像.
- **2022-08-17** 更新`4.0.10`版 docker 镜像.
- **2022-08-12** 更新`4.0.9`版 docker 镜像.
- **2022-08-07** 更新`4.0.8`版 docker 镜像.
- **2022-07-31** 更新`4.0.7`版 docker 镜像.
- **2022-07-28** 更新`4.0.6`版 docker 镜像.
- **2022-07-15** 更新`4.0.5`版 docker 镜像.
- **2022-07-14** 更新`4.0.4`版 docker 镜像.
- **2022-07-13** 更新`4.0.3`版 docker 镜像.
- **2022-07-12** 更新`4.0.2`版 docker 镜像.
- **2022-07-11** 更新`4.0.1`版 docker 镜像.
- **2022-07-10** 更新`4.0.0`版 docker 镜像.
- **2022-06-17** 更新`3.2.6`版 docker 镜像.
- **2022-05-16** 更新`3.2.5`版 docker 镜像.
- **2022-05-04** 更新`3.2.4`版 docker 镜像.
- **2022-04-13** 更新`3.2.3`版 docker 镜像.
- **2022-04-02** 更新`3.2.2`版 docker 镜像.
- **2022-02-02** 更新`3.2.1`版 docker 镜像.
- **2021-09-19** 更新`3.2`版 docker 镜像.
- **2021-06-11** 更新`3.1`版 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
首次运行会自动创建数据库目录和本地存储目录.后期迁移可直接将整个 zfile 目录备份恢复,并再次执行以下命令.
```shell
docker run -d \
    --name=zfile \
    --restart=always \
    -p 8080:8080 \
    -v /root/zfile/conf:/root/.zfile-v4 \
    -v /root/zfile/data:/root/zfile/data \
    stilleshan/zfile
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/zfile/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

### 群晖 docker
[群晖NAS高级服务 - docker 部署 zfile 在线文件目录](https://www.ioiox.com/archives/93.html)

### 配置
访问以下安装地址进行安装配置.如使用本地存储,路径请填写`/root/zfile/data`.同时可以配置反向代理使用域名访问.
> 可以自行修改挂载命令来配置本地存储路径.
```shell
http://IP:8080/install
# 安装地址
http://IP:8080/
# 前端页面
http://IP:8080/admin
# 管理页面
```

## 参考
- GitHub [zhaojun1998/zfile](https://github.com/zhaojun1998/zfile)
- [docker 部署 zfile 在线文件目录 - 支持本地,对象存储,OneDrive等各种环境.](https://www.ioiox.com/archives/92.html)
- [群晖NAS高级服务 - docker 部署 zfile 在线文件目录](https://www.ioiox.com/archives/93.html)

