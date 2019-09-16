# 使用minikube

## 创建一个简单的hello world

1. simpleweb
2. 创建一个Dockerfile
3. ```cd simpleweb; docker build -t simpleweb:0.0.1 .```

## minikube创建集群

首先,使用阿里提供的国内版本minikube
```shell
curl -Lo minikube http://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/releases/v1.3.1/minikube-darwin-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/

curl -Lo minikube http://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/releases/v1.3.1/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
```

```shell
# minikube start
# minikube start --vm-driver=hyperkit --alsologtostderr --v 10
# 从https://github.com/AliyunContainerService/minikube/releases下载国内版本
# minikube start --registry-mirror=https://registry.docker-cn.com
minikube start --registry-mirror=https://docker.mirrors.ustc.edu.cn
# 如果出错,需要
minikube delete # 如果还不行
rm -rf ~/.minikube

minikube status
#minikubeVM: Running
#localkube: Running
minikube dashboard # 在浏览器中打开dashboard
```

## minikube架构说明

下面是从[云栖社区](https://yq.aliyun.com/articles/221687)找到的一张架构图
![minikube-architecture](minikube-architecture.jpeg)

1. minikube生成kubeconfig文件
2. minikube生成minikube虚拟机
3. minikube在虚拟机里面创建kubernetes
4. kubectl使用kubernetes来进行各种操作
