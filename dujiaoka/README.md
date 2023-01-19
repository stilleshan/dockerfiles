# dujiaoka
GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/dujiaoka](https://hub.docker.com/r/stilleshan/dujiaoka)
> *docker image support for X86 and ARM*


## 简介
基于 🦄独角数卡 [assimon/dujiaoka](https://github.com/assimon/dujiaoka) 的 docker 镜像，支持 docker-compose 数据持久化部署。

> *本项目基于现有的 [Apocalypsor/dujiaoka-docker](https://github.com/Apocalypsor/dujiaoka-docker) 项目重新制作支持 ARM64 架构镜像，你也可以参考下文作者博客链接部署。*


## 更新
- **2023-01-19** 更新`2.0.6`版 docker 镜像,同时支持 X86 和 ARM 架构.
- **2022-08-24** 更新`2.0.5`版 docker 镜像,同时支持 X86 和 ARM 架构.
- **2022-07-19** 更新`2.0.4`版 docker 镜像,同时支持 X86 和 ARM 架构.
- **2022-07-19** 更新`2.0.3`版 docker 镜像,同时支持 X86 和 ARM 架构.


## 示例
https://shop.ioiox.xyz


## 部署
更详细的图文教程请访问我的博客：  
[docker 部署 dujiaoka 独角数卡自动售货系统 支持 X86 和 ARM 架构](https://www.ioiox.com/archives/159.html)

### 准备配置文件
```shell
mkdir dujiaoka && cd dujiaoka
# 创建主目录
mkdir storage uploads
chmod -R 777 storage uploads
# 创建数据目录用于挂载
```

下载本项目中的`env.conf`和`docker-compose.yml`到上述的`dujiaoka`目录中。

```shell
chmod -R 777 env.conf
# 重要步骤
```

### 配置 nginx
参考`domian.conf`修改`域名`和`证书路径`，最重要的是`location ^~ /`中的相关配置。
```nginx
    location ^~ / {
        proxy_pass http://127.0.0.1:56789;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header REMOTE-HOST $remote_addr;
        proxy_set_header X-Forwarded-Proto  $scheme;

        add_header X-Cache $upstream_cache_status;

        proxy_set_header Accept-Encoding "";
        sub_filter "http://" "https://";
        sub_filter_once off;
    }
```

### 启动容器
确保目录中有`docker-compose.yml`,`env.conf`,`storage`和`uploads`后执行以下命令`首次`启动：
> *注意修改 docker-compose.yml 中的 mysql 密码*
```shell
docker-compose up -d
```
首次启动后，会自动生成`mysql`和`redis`的数据目录。

### 网页安装
访问域名开始安装：
- 将`mysql`的地址改为`db`
- 将`mysql`的用户名改为`dujiaoka`
- 将`redis`的地址改为`redis`
- 填写网站名称
- 网站 url 填写完整域名地址，例如`https://shop.ioiox.com`

点击安装，并成功安装，先修改其他参数后在进行登录使用。
```shell
docker-compose down
# 停止服务
```
将`docker-compose.yml`中的`INSTALL`改为`false`  
将`env.conf`中的`APP_DEBUG`改为`false`
此时可以看到`env.conf`中的相关配置已经自动变为安装时填写的参数。
```shell
docker-compose up -d
# 再次启动开始使用
```

### 完成部署
自此已完成安装部署，访问域名开始使用： 
- 所有配置，数据，缓存，数据库都存储在整个`dujiaoka`目录中，定期备份即可。
- 如需迁移服务器，仅需打包本目录到新服务器，执行上述所有`chmod 777`的权限命令，再次`docker-compose up -d`即可恢复上线。


## 链接
- [Apocalypsor/dujiaoka-docker](https://github.com/Apocalypsor/dujiaoka-docker)
- [如何优雅地搭建自己的发卡站](https://blog.dov.moe/posts/49102/)
- [docker 部署 dujiaoka 独角数卡自动售货系统 支持 X86 和 ARM 架构](https://www.ioiox.com/archives/159.html)
