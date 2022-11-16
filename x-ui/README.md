# x-ui

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/x-ui](https://hub.docker.com/r/stilleshan/x-ui)
> *docker image support for X86 and ARM*

## 简介
基于 [vaxilu/x-ui](https://github.com/vaxilu/x-ui) 项目的 docker 镜像.

## 更新
- **2022-11-16** 更新`0.3.2`版 docker 镜像,同时支持 X86 和 ARM 架构.
- **2021-08-01** 更新`0.3.1`版 docker 镜像,同时支持 X86 和 ARM 架构.

## 注意
x-ui 还在开发测试中,如果遇到 xray 状态 为`error`,尝试切换版本即可.

## x-ui-ssl 一键安装脚本部署
一键安装脚本部署请访问子目录 [x-ui-ssl](./x-ui-ssl) 查看详情.

## x-ui 手动部署
### docker
```shell
docker run -d --name=x-ui --restart=always --network=host stilleshan/x-ui
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/x-ui/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

### 挂载
为方便数据备份和迁移,需将容器内`/etc/x-ui/x-ui.db`文件拷贝至宿主机存储,该文件存储所有面板数据及账号信息.并再挂载至容器内.同时创建`ssl`目录,上传证书文件挂载至容器内使用.
> **注意:** 挂载 ssl 证书后,在 x-ui 面板中创建账号使用证书的路径为 /ssl/yourdomain.crt .
#### docker
通过上文首次执行`docker run`启动容器后:
```shell
docker cp x-ui:/etc/x-ui/x-ui.db .
# 拷贝 x-ui.db 至宿主机
docker stop x-ui
# 停止容器
docker rm x-ui
# 删除容器
```
创建`ssl`目录并上传证书.  
修改以下命令中挂载路径,启动使用:
```shell
docker run -d --name=x-ui --restart=always --network=host -v /your_path/x-ui.db:/etc/x-ui/x-ui.db -v /your_path/ssl:/ssl stilleshan/x-ui
```

#### docker-compose
通过上文首次执行`docker-compose up -d`启动容器后:
```shell
docker cp x-ui:/etc/x-ui/x-ui.db .
# 拷贝 x-ui.db 至宿主机
docker-compose down
# 停止并删除容器
```
创建`ssl`目录并上传证书.  
修改`docker-compose.yml`,取消注释,参考如下:
```yml
version: '3.7'
services:
  x-ui:
    image: stilleshan/x-ui
    container_name: x-ui
    volumes:
      - ./x-ui.db:/etc/x-ui/x-ui.db
      - ./ssl:/ssl
    restart: always
    network_mode: host
```

确认挂载路径无误后启动使用:
```shell
docker-compose up -d
```

## 使用
访问`http://服务器IP:54321`使用账号`admin`密码`admin`登录.注意需开放相关端口防火墙,并及时修改账号密码.

## 参考
GitHub [vaxilu/x-ui](https://github.com/vaxilu/x-ui)

