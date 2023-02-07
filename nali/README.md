# nali
GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/nali](https://hub.docker.com/r/stilleshan/nali)
> *docker image support for X86 and ARM*

## 简介
基于 [zu1k/nali](https://github.com/zu1k/nali) 项目的 docker 镜像.


## docker
由于镜像不含 IP 数据库,所以建议后台启动容器服务,挂载目录,每次使用将会自动下载和保留 IP 数据库到本地,方便后续多次使用.
```shell
docker run -itd --name nali --restart always \
  -v /root/nali:/root/.local\
  stilleshan/nali
```

查询命令
```shell
docker exec -it nali /nali 1.1.1.1
# 返回 1.1.1.1 [澳大利亚 APNIC/CloudFlare公共DNS服务器] 
```

## 参考
- GitHub [zu1k/nali](https://github.com/zu1k/nali) 
