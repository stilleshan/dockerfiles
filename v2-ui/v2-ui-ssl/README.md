# v2-ui-ssl

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/v2-ui](https://hub.docker.com/r/stilleshan/v2-ui)
> *docker image support for X86 and ARM*

## 简介
整合 v2-ui 和 acme.sh 的 docker compose 一键部署方案
- 支持全自动证书申请,定时每月续签证书.
- 备份和迁移方便,仅需保存 v2-ui-ssl 一个目录,一条命令即可在其他服务器启动使用.
- 支持 X86 和 ARM 架构

## 更新
**2021-06-17** 更新`一键安装脚本`.

## 一键安装脚本部署
- 使用一键安装脚本后无需继续参考以下`手动部署`段落.
- 一键安装脚本依旧需要基于本服务器安装 docker 和 docker compose 环境.
- 一键安装脚本仅限于首次使用,确保服务器之前没有使用过本仓库的`v2-ui`或`v2-ui-ssl`项目.
- 脚本安装完毕后,会自行删除,避免误操作.
- 脚本安装完毕后,依旧需要手动执行以下`自动更新证书`段落的操作,以便`v2-ui`定时重启使得证书生效.
```shell
wget https://raw.githubusercontent.com/stilleshan/dockerfiles/main/v2-ui/v2-ui-ssl/install.sh && chmod +x install.sh && ./install.sh
```


## 手动部署
### 下载文件
部署所需文件已打包,方便直接下载解压使用.
```shell
wget https://raw.githubusercontent.com/stilleshan/dockerfiles/main/v2-ui/v2-ui-ssl/v2-ui-ssl.tar
tar -xvf v2-ui-ssl.tar
cd v2-ui-ssl
```

### 配置
修改`conf/account.conf`文件,在`DOMAIN=`之后填写域名,并参考一下官方文档,填写其他对应的 API 信息.  
https://github.com/acmesh-official/acme.sh/wiki/dnsapi  
https://www.ioiox.com/archives/87.html


### 启动
执行以下命令启动,请确保`account.conf`填写正确,容器每次启动都会检测`account.conf`中的`DOMAIN`变量,以及`ssl`目录下是否存在该域名证书,如果没有将会自动申请证书,如已存在则不会申请.
```shell
docker-compose up -d
```
> 启动后稍等一分钟, **ssl** 目录下将会生成证书文件.

### 自动更新证书
由于系统将在 GMT 时间每月`1`日`0`点,也就是北京时间`1`日早`8`点重新强制更新证书,并覆盖至`ssl`目录下.则需要同时为`v2-ui`服务设置一个定时重启计划任务：
```shell
crontab -e
# 添加以下计划任务
0 0 2 * * docker restart v2-ui
# 为避免时区问题,将在每月 2 号 0 点执行
```

### 手动更新证书
也可以手动执行更新证书命令：
```shell
docker exec acme sh -c /conf/acme.sh
```

## 使用
访问`http://服务器IP:65432`使用账号`admin`密码`admin`登录.注意需开放相关端口防火墙,并及时修改账号密码.

### 证书使用
**v2-ui** 中创建账户开启`tls`时所需用到的证书路径为
> /ssl/yourdomain.cer  
/ssl/yourdomain.key


## 参考
GitHub [sprov065/v2-ui](https://github.com/sprov065/v2-ui)

