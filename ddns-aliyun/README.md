# ddns-aliyun

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/ddns-aliyun](https://hub.docker.com/r/stilleshan/ddns-aliyun)
> *docker image support for X86 and ARM*

## 简介
基于阿里云 DDNS 的 docker 镜像.

## 更新
**2021-06-11** 更新 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d \
    --name=ddns-aliyun \
    --restart=always \
    -e "AKID=AccessKey-ID" \
    -e "AKSCT=AccessKey-Secret" \
    -e "DOMAIN=ddns.yourdomain.com" \
    -e "REDO=600" \
    stilleshan/ddns-aliyun
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/ddns-aliyun/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

### 配置
- AKID: 填写`AccessKeyID`
- AKSCT: 填写`AccessKeySecret`
- DOMAIN: 填写`ddns 域名`

### 示例
> 使用二级域名来作为 ddns 域名,例如以下示例使用 ddns.ioiox.com ,需要使用的正式域名添加 CNAME 指向 ddns.ioiox.com 即可.
```shell
docker run -d \
    --name=ddns-aliyun \
    --restart=always \
    -e "AKID=kzazFUTW0uRIWWtk" \
    -e "AKSCT=kcbGxOVmc9PUpJWTBWNFNWWnNSbFJW" \
    -e "DOMAIN=ddns.ioiox.com" \
    -e "REDO=600" \
    stilleshan/ddns-aliyun
```

## 参考
- GitHub [honwen/aliyun-ddns-cli](https://github.com/honwen/aliyun-ddns-cli)
- [群晖NAS网络服务 - docker 部署配置阿里云 DDNS 动态域名解析](https://www.ioiox.com/archives/29.html)

