# php-nginx
build arm image based [webdevops/Dockerfile](https://github.com/webdevops/Dockerfile)

## 镜像链
所需镜像`php-nginx:7.4-alpine`需要上游基础镜像`php:7.4-alpine`.  
而`php:7.4-alpine`又需要上游基础镜像`toolbox`.

上述镜像由 [webdevops/Dockerfile](https://github.com/webdevops/Dockerfile) 维护,均为 X86 架构.  
根据每个基础镜像的 Dockerfile 文件发现了部分 X86 二进制执行文件,很难直接使用 buildx 直接构建多架构.  
需手动修改相关命令支持 ARM, 尝试使用 GitHub Action 来构建多次因部分测试命令报错, 最终只能使用 ARM 服务器来一层层手动构建.

## 构建 ARM 镜像
### toolbox
```shell
cd toolbox/latest
docker build -t tookbox .
```

### php:7.4-alpine
```shell
cd php/7.4-alpine
docker build -t php:7.4-alpine .
```

### php-nginx:7.4-alpine
```shell
cd 7.4-alpine
docker build -t php-nginx:7.4-alpine .
```

## 合并镜像
上述已构建出 ARM 架构`php-nginx:7.4-alpine`, 加上官方的 X86 架构`webdevops/php-nginx:7.4-alpine`, 使用起来并不方便,推荐使用`manifest`来合并镜像.

### 推送现有镜像
制作单镜像多架构,需要先把 2 个不同架构的镜像分别推送至 docker hub 上.  
所以需要 X86 和 ARM 两台服务器分别执行.  
在推送之前可以使用`tag`命令规划好镜像名方便区分.

- X86 ioiox/php-nginx-x86:latest
- ARM ioiox/php-nginx-arm:latest

#### X86
> *注意请使用 X86 服务器,先下载 webdevops 的 X86 版本,重新`tag`并推送至 docker hub.*
```shell
docker pull webdevops/php-nginx:7.4-alpine
docker tag webdevops/php-nginx:7.4-alpine ioiox/php-nginx-x86:latest
docker push ioiox/php-nginx-x86:latest
```

#### ARM
> *注意请使用 ARM 服务器将上述构建的 ARM 版`php-nginx:7.4-alpine`重新`tag`并推送至 docker hub.*
```shell
docker tag php-nginx:7.4-alpine ioiox/php-nginx-arm:latest
docker push ioiox/php-nginx-arm:latest
```

### 创建 manifest 列表
创建单镜像的 manifest 列表, 建议加上`tag`,列表中包含了上文推送的 2 个架构的镜像.
> *以下操作可以在任意架构服务器执行*
```shell
docker manifest create ioiox/php-nginx:7.4-alpine \
    ioiox/php-nginx-x86:latest \
    ioiox/php-nginx-arm:latest
```

### 设置 manifest 列表
为对应的架构设置对应的镜像
```shell
docker manifest annotate ioiox/php-nginx:7.4-alpine \
    ioiox/php-nginx-x86:latest \
    --os linux --arch x86_64

docker manifest annotate ioiox/php-nginx:7.4-alpine \
    ioiox/php-nginx-arm:latest \
    --os linux --arch arm64 --variant v8
```

### 查看 manifest 列表
查看并检查 manifest 列表
```shell
docker manifest inspect ioiox/php-nginx:7.4-alpine
```

### 推送 manifest 列表
推送 manifest 列表至 docker hub
```shell
docker manifest push ioiox/php-nginx:7.4-alpine
```

此时已完成镜像合并,在 X86 或 ARM 服务器上使用`ioiox/php-nginx:7.4-alpine`镜像就会使用各自的架构镜像,同时可以在 docker hub 上删除刚才推送上去的`ioiox/php-nginx-x86:latest`和`ioiox/php-nginx-arm:latest`也不会影响新镜像的使用.
