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

## Run实例

1. 运行一个容器实例
    ```shell
    docker run --name 容器名 -i -t -p 主机端口:容器端口 -d -v 主机目录:容器目录:ro 镜像TD或镜像名:TAG

    # --name 指定容器名，可自定义，不指定自动命名
    # -i 以交互模式运行容器
    # -t 分配一个伪终端，即命令行，通常组合来使用
    # -p 指定映射端口，将主机端口映射到容器内的端口
    # -d 后台运行容器
    # -v 指定挂载主机目录到容器目录，默认为rw读写模式，ro表示只读
    ```
2. 进入一个容器
    ```shell
    docker exec -it 容器ID或者容器名 /bin/bash

    # 进入正在运行的容器并且开启交互模式终端
    # /bin/bash是固有写法，作用是因为docker后台必须运行一个进程，否则容器就会退出，在这里表示启动容器后启动bash。
    # 也可以用docker exec在运行中的容器执行命令
    ```
3. 拷贝文件
    ```shell
    docker cp 主机文件路径 容器ID或容器名:容器路径 # 主机中文件拷贝到容器中
    docker cp 容器ID或容器名:容器路径 主机文件路径 # 容器中文件拷贝到主机中
    ```
4. 编译镜像
    ```shell
    docker build -f Dockerfile -t zong/tomcat:v2.0 .

    # -f Dockerfile路径，默认是当前目录
    # -t 指定新镜像的名字以及TAG
    ```
5. 镜像导入导出,备份与恢复
    ```shell
    docker export 容器ID >centos_v1.tar
    docker import - centos_v1 <centos_v1.tar
    docker save -o centos_backup.tar centos:update
    docker load <centos_backup.tar
    ```
6. 使用镜像
    ```shell
    vi /etc/docker/daemon.json
    # 填写自己的加速器地址
    # {
    #    "registry-mirrors": ["https://xxxxxx.mirror.aliyuncs.com"]
    # }
    systemctl daemon-reload
    systemctl restart docker
    service docker restart
    ```