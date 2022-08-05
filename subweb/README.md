# subweb
## 简介
subweb 是基于 subconverter 订阅转换的前端项目,方便用户快速生成各平台的订阅链接.

> *subweb 是我个人入门 vuejs 学习时简单做的一个案例,使用还算方便,开源出来,欢迎各路大佬贡献维护.*

*GitHub [stilleshan/subweb](https://github.com/stilleshan/subweb)  
Docker [stilleshan/subweb](https://hub.docker.com/r/stilleshan/subweb)*
> *docker image support for X86 and ARM*

## 示例
[https://sub.ops.ci](https://sub.ops.ci)  
*`前后端示例,可以直接使用.`*

## 部署
### docker 本地版
*适用于本机快速部署使用*
```shell
docker run -d --name subweb --restart always \
  -p 18080:80 \
  stilleshan/subweb
```

访问 `http://127.0.0.1:18080`

### docker 自定义版 + 短链接版
自定义版可以挂载配置文件来修改`API 地址`,`短链接地址`,`站点名称`,`导航链接`.  
参考以下命令,修改本地挂载路径,启动容器后会生成`config.js`配置文件,更改后刷新生效.

```shell
docker run -d --name subweb --restart always \
  -p 18080:80 \
  -v /PATH/subweb/conf:/usr/share/nginx/html/conf \
  stilleshan/subweb
```

同时也可以不挂载目录,直接通过`-e`环境变量来修改`API 地址`,`短链接地址`和`站点名称`,但是无法修改`导航链接`.
`注意:以下域名请严格填写 http/https 协议,结尾不要 /`
```shell
docker run -d --name subweb --restart always \
  -p 18080:80 \
  -e SITE_NAME=subweb \
  -e API_URL=https://sub.ops.ci \
  -e SHORT_URL=https://s.ops.ci \
  stilleshan/subweb
```

访问 `http://127.0.0.1:18080`  
> *推荐使用 nginx 反向代理部署*

### subweb + subconverter + myurls 合并进阶版
详情查看 [stilleshan/sub](https://github.com/stilleshan/dockerfiles/tree/main/sub)

## 链接
- [stilleshan/sub](https://github.com/stilleshan/dockerfiles/tree/main/sub)
- [stilleshan/subweb](https://github.com/stilleshan/subweb)
- [stilleshan/subconverter](https://github.com/stilleshan/subconverter)
