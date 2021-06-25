# coscmd

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/coscmd](https://hub.docker.com/r/stilleshan/coscmd)
> *docker image support for X86 and ARM*

## 简介
基于 [腾讯云 COSCMD 工具](https://cloud.tencent.com/document/product/436/10976) 的 docker 镜像.

## 更新
**2021-06-25** 更新 docker 镜像,已测试同时支持 X86 和 ARM 架构.

## 使用
腾讯云 COSCMD 工具 docker 镜像用于启动一次性容器来完成单次的上传或下载需求.

### 配置文件
参考 [官方文档](https://cloud.tencent.com/document/product/436/10976) 创建`.cos.conf`配置文件.
```conf
[common]
secret_id = AKIDA6wUmImTMzvXZNbGLCgtusZ2E8mG****
secret_key = TghWBCyf5LIyTcXCoBdw1oRpytWk****
bucket = configure-bucket-1250000000
region = ap-chengdu
max_thread = 5
part_size = 1
retry = 5
timeout = 60
schema = https
verify = md5
anonymous = False
```

### docker
挂载`.cos.conf`配置文件至容器内`/root/.cos.conf`以及本地需要上传或下载的目录.
```shell
docker run --rm \
    -v /path/.cos.conf:/root/.cos.conf \
    -v /data:/data \
    stilleshan/coscmd \
    coscmd upload -rfs --delete /data /dst_data
```

## 参考
- [腾讯云 COSCMD 工具](https://cloud.tencent.com/document/product/436/10976)
- [COSCMD 工具类常见问题](https://cloud.tencent.com/document/product/436/30744)

