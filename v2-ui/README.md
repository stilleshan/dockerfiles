# v2-ui

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/v2-ui](https://hub.docker.com/r/stilleshan/v2-ui)
> *docker image support for X86 and ARM*

## 简介
基于 [sprov065/v2-ui](https://github.com/sprov065/v2-ui) 项目的 docker 镜像.

## 更新
- **2021-06-17** 更新`一键安装脚本`.
- **2021-06-13** 更新`5.5.2`版 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 一键安装脚本部署
- 使用一键安装脚本后无需继续参考以下`手动部署`段落.
- 脚本依旧需要基于本服务器安装 docker 和 docker compose 环境.
- 脚本仅限于首次使用,确保服务器之前没有使用过本仓库的`v2-ui`或`v2-ui-ssl`项目.
- 脚本安装完毕后,会自行删除,避免误操作.
- 脚本安装完毕后,依旧需要参考 [自动更新证书](https://github.com/stilleshan/dockerfiles/tree/main/v2-ui-ssl#%E8%87%AA%E5%8A%A8%E6%9B%B4%E6%96%B0%E8%AF%81%E4%B9%A6) 段落手动执行定时任务操作,以便`v2-ui`定时重启使得证书生效.
```shell
wget https://raw.githubusercontent.com/stilleshan/dockerfiles/main/v2-ui-ssl/install.sh && chmod +x install.sh && ./install.sh
```

## 手动部署
### docker
```shell
docker run -d --name=v2-ui --restart=always --network=host stilleshan/v2-ui
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/v2-ui/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

### 挂载
为方便数据备份和迁移,需将容器内`/etc/v2-ui/v2-ui.db`文件拷贝至宿主机存储,该文件存储所有面板数据及账号信息.并再挂载至容器内.同时创建`ssl`目录,上传证书文件挂载至容器内使用.
> **注意:** 挂载 ssl 证书后,在 v2-ui 面板中创建账号使用证书的路径为 /ssl/yourdomain.crt .
#### docker
通过上文首次执行`docker run`启动容器后:
```shell
docker cp v2-ui:/etc/v2-ui/v2-ui.db .
# 拷贝 v2-ui.db 至宿主机
docker stop v2-ui
# 停止容器
docker rm v2-ui
# 删除容器
```
创建`ssl`目录并上传证书.  
修改以下命令中挂载路径,启动使用:
```shell
docker run -d --name=v2-ui --restart=always --network=host -v /your_path/v2-ui.db:/etc/v2-ui/v2-ui.db -v /your_path/ssl:/ssl stilleshan/v2-ui
```

#### docker-compose
通过上文首次执行`docker-compose up -d`启动容器后:
```shell
docker cp v2-ui:/etc/v2-ui/v2-ui.db .
# 拷贝 v2-ui.db 至宿主机
docker-compose down
# 停止并删除容器
```
创建`ssl`目录并上传证书.  
修改`docker-compose.yml`,取消注释,参考如下:
```yml
version: '3.7'
services:
  v2-ui:
    image: stilleshan/v2-ui
    container_name: v2-ui
    volumes:
      - ./v2-ui.db:/etc/v2-ui/v2-ui.db
      - ./ssl:/ssl
    restart: always
    network_mode: host
```

确认挂载路径无误后启动使用:
```shell
docker-compose up -d
```

## 使用
访问`http://服务器IP:65432`使用账号`admin`密码`admin`登录.注意需开放相关端口防火墙,并及时修改账号密码.

## 参考
GitHub [sprov065/v2-ui](https://github.com/sprov065/v2-ui)

