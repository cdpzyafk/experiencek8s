# 使用minikube

## 创建一个简单的hello world

1. simpleweb
2. 创建一个Dockerfile
3. ```cd simpleweb; docker build -t simpleweb:0.0.1 .```

## minikube创建集群

```shell
minikube start
minikube start --vm-driver=hyperkit --alsologtostderr --v 10
# 从https://github.com/AliyunContainerService/minikube/releases下载国内版本
minikube start --registry-mirror=https://registry.docker-cn.com
# 如果出错,需要
minikube delete # 如果还不行
rm -rf ~/.minikube

minikube status
#minikubeVM: Running
#localkube: Running
minikube dashboard # 在浏览器中打开dashboard
```