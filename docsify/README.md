# docsify

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/docsify](https://hub.docker.com/r/stilleshan/docsify)
> *docker image support for X86 and ARM*

## 简介
docsify serve docker, 方便在 macOS 或 Windows 本地使用 docker 来启动 docsify 服务进行编写文档.

## 使用
### docker
```shell
docker run --name docsify -p 3000:3000 -v /your/docs/path:/docs stilleshan/docsify
# 修改本地文档目录
```
访问`127.0.0.1:3000`
> 注意: 如本地文档目录内含有 index.html 并是初次初始化 docsify, 需重命名 index.html 进行备份,再次启动容器.

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/docsify/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
- docsify [https://docsify.js.org](https://docsify.js.org)
