# forsaken-mail

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/forsaken-mail](https://hub.docker.com/r/stilleshan/forsaken-mail)
> *docker image support for X86 and ARM*

## 简介
基于 [forsaken-mail](https://github.com/denghongcai/forsaken-mail) 临时邮箱项目的 docker 镜像.

## 更新
**2021-06-21** 更新 docker 镜像,已简单测试能够正常收取`126`和`hotmail`邮箱发送的邮件.

## 部署
### docker
```shell
docker run -d --name=forsaken-mail --restart=always -p 25:25 -p 3000:3000 stilleshan/forsaken-mail
```

### docker compose
下载 [docker-compose.yml](https://raw.githubusercontent.com/stilleshan/dockerfiles/main/forsaken-mail/docker-compose.yml) 执行以下命令启动:
```shell
docker-compose up -d
```

### 配置域名
示例,使用`temp.yourdomain.com`作为临时邮箱后缀.将会生成类似`xxxxx@temp.yourdomain.com`的临时邮箱.
#### 设置邮箱域名 A 记录
设置`temp`的`A`记录至服务器 IP .

#### 设置邮箱域名 MX 记录
设置`temp`的`MX`记录至上述`temp.yourdomain.com`,优先级设置为`10`.

## 参考
GitHub [denghongcai/forsaken-mail](https://github.com/denghongcai/forsaken-mail)

