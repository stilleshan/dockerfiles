# docsify

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/docsify](https://hub.docker.com/r/stilleshan/docsify)
> *docker image support for X86 and ARM*

## 简介
docsify serve docker, 方便在 macOS 或 Windows 本地使用 docker 来启动 docsify 服务进行编写文档.

## 使用
### 现有项目
需挂载本地`docsify`文档目录到容器内`/docs`,启动容器使用`127.0.0.1:3000`访问.
```shell
docker run -d --name=docsify -p 3000:3000 -v /your/docs/path:/docs stilleshan/docsify
```

### 新建项目
```shell
docker run -d --name=docsify -p 3000:3000 -v stilleshan/docsify
# 不挂载启动容器,初始化新文档.
docker cp docsify:/docs .
# 将容器内初始化的文档拷贝出来.
```
后续可参考上述`现有项目`方式挂载来维护撰写文档.

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/docsify/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
- docsify [https://docsify.js.org](https://docsify.js.org)
