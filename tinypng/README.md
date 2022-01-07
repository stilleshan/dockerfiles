# tinypng

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/tinypng](https://hub.docker.com/r/stilleshan/tinypng)
> *docker image support for X86 and ARM*

## 简介
使用 TinyPNG 免费压缩图片 docker 小程序，方便本地批量压缩，支持多个 **API key** 配置。

> [TinyPNG](https://tinypng.com) 是一个非常强大，高画质，高压缩比的在线压缩图片的网站，并提供免费的 API key 用于本地批量压缩。

## 使用
### 获取 API key
访问 [TinyPNG Developer API](https://tinypng.com/developers) 获取免费 API key，免费用户每个 key 每月可以压缩 500 张图片，可以创建多个免费账户来获取多个 key 配合本程序使用。

### 准备目录及 key 文件
1. 在需要压缩照片的目录中创建`api_key.txt`并将`key`粘贴进去，多个`key`时，每行一个`key`。
2. 确保该目录下没有`Output`文件夹，`Output`文件夹是由本 docker 压缩图片下载后生成。
3. 图片支持 jpg / jpeg / png 格式,并支持多层子目录。

### docker
```shell
docker run --rm -v /your/pic/path:/pic stilleshan/tinypng
# 修改 /your/pic/path 为需要压缩的图片目录
```

## 参考
- Python 来源于 [haoma2012/PythonProject](https://github.com/haoma2012/PythonProject)
