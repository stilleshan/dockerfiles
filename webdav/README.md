# webdav

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/webdav](https://hub.docker.com/r/stilleshan/webdav)
> *docker image support for X86 and ARM*

## 简介
基于 [hacdias/webdav](https://github.com/hacdias/webdav) 项目的 docker 镜像.
> 一款 go 语言,轻量化的 WebDAV 服务端程序,支持 docker 一键快速部署.

## 更新
- **2021-09-11** 创建 docker 镜像,同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d --name=webdav --restart=always -p 8080:8080 stilleshan/webdav
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/webdav/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 使用
### 快速使用
为避免端口冲突,镜像默认端口为`8080`, 默认挂载路径为`/data`, 使用 WebDAV 客户端连接 http://ip:8080 , 账号`admin`密码`admin`登陆.

### 自定义使用
可参考原始项目修改`config.yml`来配置用户,权限,路径,证书等参数,并挂载至`/config.yml`使用,示例:
```shell
docker run -d --name=webdav --restart=always -v /path/config.yml:/config.yml -p 8080:8080 stilleshan/webdav
```

### 反向代理
可以配置`nginx`反向代理来实现`HTTPS`加密及无端口访问,以下为相关反代配置,已经过`macOS`,`Windows`,`infuse`,`nPlayer`等客户端测试.
```nginx
    location / {
        proxy_redirect off;
        proxy_pass http://127.0.0.1:8080;
        # 反代地址及端口以当前服务器环境而定.

        proxy_set_header  Host                $http_host;
        proxy_set_header  X-Real-IP           $remote_addr;
        proxy_set_header  X-Forwarded-Ssl     on;
        proxy_set_header  X-Forwarded-For     $proxy_add_x_forwarded_for;
        proxy_set_header  X-Forwarded-Proto   $scheme;
        proxy_set_header  X-Frame-Options     SAMEORIGIN;

        client_max_body_size        100m;
        client_body_buffer_size     128k;

        proxy_buffer_size           4k;
        proxy_buffers               4 32k;
        proxy_busy_buffers_size     64k;
        proxy_temp_file_write_size  64k;
    }
```

## 参考
GitHub [hacdias/webdav](https://github.com/hacdias/webdav)
