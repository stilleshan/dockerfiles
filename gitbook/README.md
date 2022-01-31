# gitbook

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/gitbook](https://hub.docker.com/r/stilleshan/gitbook)
> *docker image support for X86 and ARM*

## 简介
gitbook for docker, 方便本地使用 docker 来启动服务编写文档.

## 使用
### 初始化 init
```shell
docker run --rm -v "$PWD:/gitbook" -p 4000:4000 stilleshan/gitbook gitbook init
# 初始化当前目录
```

### 服务 serve
```shell
docker run --rm -v "$PWD:/gitbook" -p 4000:4000 stilleshan/gitbook gitbook serve
# 启动 serve 服务
```
访问`http://localhost:4000`

### 构建 build
```shell
docker run --rm -v "$PWD:/gitbook" -p 4000:4000 stilleshan/gitbook gitbook build
# 打包构建
```

## 参考
- [Gitbook](https://github.com/GitbookIO/gitbook)
