
生成协议代码
```shell script
protoc --proto_path=. --go_out=. --micro_out=. proto/greeter.proto
```

运行服务器
```shell script
go run main.go
```
运行客户端
```shell script
go run main.go --run_client
```
