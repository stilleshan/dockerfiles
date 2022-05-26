# registry-proxy

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/registry-proxy](https://hub.docker.com/r/stilleshan/registry-proxy)
> *docker image support for X86 and ARM*

## 简介
代理 ghcr.io / gcr.io / k8s.gcr.io 等镜像仓库的 registry proxy 镜像.

## 使用
### docker
环境变量`PROXY_REMOTE_URL`:
- https://ghcr.io
- https://gcr.io
- https://k8s.gcr.io

```shell
docker run -d --restart always \
    -v /data/path:/var/lib/registry\
    -p 5000:5000 \
    -e PROXY_REMOTE_URL=https://ghcr.io \
    stilleshan/registry-proxy
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/registry-proxy/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
- GitHub [findsec-cn/registry-proxy](https://github.com/findsec-cn/registry-proxy)
