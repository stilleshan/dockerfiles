# sub

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/sub](https://hub.docker.com/r/stilleshan/sub)
> *docker image support for X86 and ARM*

## 简介
基于 subweb 和 subconverter 前后端整合容器,支持自定义配置.

## 部署
**注意,本项目必须搭配域名反代使用.**

### docker
```shell
docker run -d --name subweb --restart always \
  -p 18080:80 \
  -v /PATH/subweb/conf:/usr/share/nginx/html/conf \
  stilleshan/sub
```
修改挂载路径,根据需求自行修改`config.js`中的相关配置.

推荐使用`nginx`配置域名反向代理至`18080`端口.

`subconverter`同样支持挂载外部配置文件,参考容器内部路径:`/base/snippets/rulesets.txt`.

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/sub/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
[stilleshan/subweb](https://github.com/stilleshan/subweb)
[stilleshan/subconverter](https://github.com/stilleshan/subconverter)
