# sub

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/sub](https://hub.docker.com/r/stilleshan/sub)
> *docker image support for X86 and ARM*

## 简介
基于 subweb 和 subconverter 前后端加上 myurls 短链接整合容器,支持自定义配置.

## 部署
**注意,本项目必须搭配域名反代使用.**

### docker
> 已更新支持短链接,如需要更换短链接或自己部署,请使用以下 docker compose 部署.
```shell
docker run -d --name subweb --restart always \
  -p 18080:80 \
  -v /PATH/subweb/conf:/usr/share/nginx/html/conf \
  stilleshan/sub
```
修改挂载路径,根据需求自行修改`conf/config.js`中的相关配置.

推荐使用`nginx`配置域名反向代理至`18080`端口.

`subconverter`同样支持挂载外部配置文件,参考容器内部路径:`/base/snippets/rulesets.txt`.

### docker compose
docker compose 已包含 myurl 短链接:
- 如无需部署 myurls 服务,可删除`12-32`行,将默认使用本站短链接.也可以修改`conf/config.js`来使用其他`myurls`短链接服务.
- 如需自行部署 myurls 服务,需修改`docker-compose.yml`中的`MYURLS_DOMAIN`,以及`conf/config.js`中的`shortUrl`,注意请严格按照示例格式填写.
- myurls 服务需要单独配置 nginx 反代以及证书,可以参考`myurls.conf`配置.注意需要修改`域名`,`证书路径`,`日志路径`.
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/sub/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

## 参考
- [stilleshan/subweb](https://github.com/stilleshan/subweb)
- [stilleshan/subconverter](https://github.com/stilleshan/subconverter)
