# mc

GitHub [stilleshan/dockerfiles](https://github.com/stilleshan/dockerfiles)  
Docker [stilleshan/mc](https://hub.docker.com/r/stilleshan/mc)
> *docker image support for X86 and ARM*

## 简介
基于 [minio/mc](https://github.com/minio/mc) 项目的**执行** docker 镜像.

## 使用
参照示例,挂载数据目录.
```shell
docker run --rm \
    -v /data/path:/data \
    stilleshan/mc sh -c 'mc config host add minio https://minio.domain.com user password --api s3v4 && mc cp /data/foo.tar minio/buckets/path/'
```
