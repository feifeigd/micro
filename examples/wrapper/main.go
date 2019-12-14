package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "d7kj.com/m/v2/examples/service/proto"
	"github.com/micro/go-micro/server"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.WrapHandler(logWrapper),	// 相当于hook吧
		)
	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

type Greeter struct {

}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{})error{
		log.Printf("[wrapper] server request: %v\n", req.Endpoint())
		return fn(ctx, req, rsp)
	}
}