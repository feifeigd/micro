
安装micro
```shell
go get -u github.com/micro/micro
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/micro/protoc-gen-micro
```

指定服务发现注册中心
设置环境变量
```shell script
MICRO_REGISTRY=etcd
```
或者在命令行加上参数
```shell script
--registry etcd
```
