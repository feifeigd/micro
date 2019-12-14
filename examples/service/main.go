package main

import (
	"context"
	proto "d7kj.com/m/v2/examples/service/proto"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"os"
)

func main(){
	service := micro.NewService(
		micro.Name("com.d7kj.micro.greeter"),
		micro.Metadata(map[string]string{"type": "hello,world"}),
		// 命令行选项
		micro.Flags(cli.BoolFlag{Name: "run_client", Usage: "launch the client",}),
		)

	service.Init(
		micro.Action(func(c *cli.Context){
			if c.Bool("run_client"){
				runClient(service)
				os.Exit(0)
			}
		}),
		)
	
	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

type Greeter struct{

}

func (g *Greeter)Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func runClient(service micro.Service) {
	// Create new greeter client
	greeter := proto.NewGreeterService("com.d7kj.micro.greeter", service.Client())
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Msg)
}
