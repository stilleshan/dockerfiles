# ddns-aliyun

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/ddns-aliyun](https://hub.docker.com/r/stilleshan/ddns-aliyun)
> *docker image support for X86 and ARM*

## 使用
### docker 启动
```shell
docker run -d \
   -e "AKID=[ALIYUN's AccessKey-ID]" \
   -e "AKSCT=[ALIYUN's AccessKey-Secret]" \
   -e "DOMAIN=ddns.yourdomain.com" \
   -e "REDO=600" \
    stilleshan/ddns-aliyun
```

- AKID: 填写`AccessKeyID`
- AKSCT: 填写`AccessKeySecret`
- DOMAIN: 填写`ddns 域名`

### 示例
> 使用二级域名来作为 ddns 域名,例如以下示例使用 ddns.ioiox.com ,需要使用的正式域名添加 CNAME 指向 ddns.ioiox.com 即可.
```shell
docker run -d \
   -e "AKID=kzazFUTW0uRIWWtk" \
   -e "AKSCT=kcbGxOVmc9PUpJWTBWNFNWWnNSbFJW" \
   -e "DOMAIN=ddns.ioiox.com" \
   -e "REDO=600" \
    stilleshan/ddns-aliyun
```

## 参考
[群晖NAS网络服务 - docker 部署配置阿里云 DDNS 动态域名解析](https://www.ioiox.com/archives/29.html)  
更多请参考原始仓库 [honwen/aliyun-ddns-cli](https://github.com/honwen/aliyun-ddns-cli)

