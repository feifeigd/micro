package main

import (
	"context"
	proto "d7kj.com/m/v2/examples/service/proto"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter.client"),
		// wrap the client
		micro.WrapClient(logWrap),	// 相当于hook吧
		)
	service.Init()
	greeter := proto.NewGreeterService("greeter", service.Client())
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name:"John"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Msg)
}

type logWrapper struct {
	client.Client
}
func (l * logWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[wrapper] client request service: %s method: %s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func logWrap(c client.Client) client.Client  {
	return &logWrapper{c}
}
