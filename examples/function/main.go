package main

import (
	"context"
	proto "d7kj.com/m/v2/examples/function/proto"
	"github.com/micro/go-micro"
)

func main()  {
	fnc := micro.NewFunction(
			micro.Name("com.d7kj.micro.fnc.greeter"),
		)
	fnc.Init()
	fnc.Handle(new(Greeter))
	fnc.Run()
}

type Greeter struct {

}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error  {
	rsp.Greeting = "Hello " + req.Name
	return nil
}
