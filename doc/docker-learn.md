# docker使用过程

## Dockerfile基础命令

1. FROM -> 母镜像
2. MAINTAINER -> 维护者信息
3. WORKDIR -> 工作目录
4. ADD -> 将文件复制到镜像中
5. RUN -> 执行操作（就跟在终端执行语句一样
6. EXPOSE -> 暴露端口
7. CMD -> 程序入口

## 编译镜像

```shell
docker build -t simpleweb:0.0.1 .  # 编译,注意最后的点
docker images                      # 查看创建的镜像
docker run -p 8080:8080 -d simpleweb:0.0.1 # demon方式运行,暴露端口8080
docker container ls                # 查看上步命令是否成功
docker container logs 42b65e1a0648 # 查看container的日志
docker ps                          # 查看运行起来的container
docker rm                          # 删除container
docker rmi                         # 删除image
```