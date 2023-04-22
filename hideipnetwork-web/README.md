# hideipnetwork-web

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/hideipnetwork-web](https://hub.docker.com/r/stilleshan/hideipnetwork-web)
> *docker image support for X86 and ARM*

## 简介
基于 node 的在线浏览器项目的 docker 镜像.

## 使用
### docker
```bash
docker run -d --name hideipnetwork -p 56559:56559 stilleshan/hideipnetwork-web
```
访问`127.0.0.1:56559`

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/hideipnetwork-web/docker-compose.yml) 执行以下命令启动:
```bash
docker-compose up -d
```

### 反向代理参考
```nginx
    location / {
        proxy_pass http://127.0.0.1:56559;
        proxy_set_header  Host                $http_host;
        proxy_set_header  X-Real-IP           $remote_addr;
        proxy_set_header  X-Forwarded-Ssl     on;
        proxy_set_header  X-Forwarded-For     $proxy_add_x_forwarded_for;
        proxy_set_header  X-Forwarded-Proto   $scheme;
        proxy_set_header  X-Frame-Options     SAMEORIGIN;

        proxy_set_header  Upgrade             $http_upgrade;
        proxy_set_header  Connection          upgrade;
    }
```

## 参考
- [自建在线网页代理](https://blog.tanglu.me/web-browser/)
