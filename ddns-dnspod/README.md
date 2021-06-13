# ddns-dnspod

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/ddns-dnspod](https://hub.docker.com/r/stilleshan/ddns-dnspod)
> *docker image support for X86 and ARM*

## 简介
基于腾讯云 dnspod.cn DDNS 的 docker 镜像.

## 更新
**2021-06-11** 更新 docker 镜像,新增同时支持 X86 和 ARM 架构.

## 部署
### docker
```shell
docker run -d \
    --name=ddns-dnspod \
    --restart=always \
    -e "login_token=token_id,token" \
    -e "domain=domain.com" \
    -e "sub_domain=www" \
    -e "interval=10" \
    -e "ip_count=1" \
    stilleshan/ddns-dnspod
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/ddns-dnspod/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

### 配置
- login_token: 填写`dnspod.cn`申请的`API TOKEN`
- domain: 填写`根域名`
- sub_domain: 填写`子域名`

### 示例
> 使用二级域名来作为 ddns 域名,例如以下示例使用 ddns.ioiox.com ,需要使用的正式域名添加 CNAME 指向 ddns.ioiox.com 即可.
```shell
docker run -d \
    --name=ddns-dnspod \
    --restart=always \
    -e "login_token=123456,5MTlmZRFZrWkdFMVTU2VFE9PX" \
    -e "domain=ioiox.com" \
    -e "sub_domain=ddns" \
    -e "interval=10" \
    -e "ip_count=1" \
    stilleshan/ddns-dnspod
```

## 参考
- GitHub [strahe/dnspod-ddns](https://github.com/strahe/dnspod-ddns)
- [群晖NAS网络服务 - docker 部署配置腾讯云 DNSPod DDNS 动态域名解析](https://www.ioiox.com/archives/112.html)

