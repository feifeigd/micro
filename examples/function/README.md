
生成协议代码
```shell script
protoc --proto_path=. --go_out=. --micro_out=. proto/greeter.proto
```

运行服务器
```shell script
go run main.go
```
Call function
```shell script
micro call com.d7kj.micro.fnc.greeter Greeter.Hello '{"name": "john"}'
```

调用一次之后，服务就会退出
