# java_oci_manage
GitHub [stilleshan/java_oci_manage](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/java_oci_manage](https://hub.docker.com/r/stilleshan/java_oci_manage)
> *docker image support for X86 and ARM*

## 简介
基于 [semicons/java_oci_manage](https://github.com/semicons/java_oci_manage) 项目的 docker 镜像.  
本项目仅为方便 docker 容器化部署,相关使用教程及问题请参考官方项目.


## docker 部署
### 准备配置文件
- 参考官方项目,准备 oracle api key 私钥文件`key.pem`.
- 参考官方项目,配置 client_config 文件,注意 key 路径为`key_file=/app/key.pem`.

### 启动 docker 容器
```bash
docker run -itd --name rrr --restart always \
  -v /root/config/client_config:/app/client_config \
  -v /root/config/key.pem:/app/key.pem \
  -p 9527:9527 \
  stilleshan/java_oci_manage
```

### 进入 docker 容器
```bash
docker exec -it rrr bash
```

### 启动服务
```bash
bash sh_client_bot.sh
```
启动服务,日志提示成功后即可 ctrl + c 退出终端,服务将在容器内后台执行.

```bash
tail -f log_r_client.log 
# 查看日志
ps -ef | grep r_client.jar | grep -v grep | awk '{print $2}' | xargs kill -9
# 停止服务进程
```

## docker compose 部署
参考完成上述配置后,下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/java_oci_manage/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 链接
- [semicons/java_oci_manage](https://github.com/semicons/java_oci_manage)
