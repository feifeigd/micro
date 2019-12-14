package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	proto "d7kj.com/m/v2/examples/pubsub/srv/proto"
)

func main() {
	service := micro.NewService(

		)
	// parse command line
	service.Init()

	// register subscriber
	micro.RegisterSubscriber("example.topic.pubsub.1", service.Server(), new(Sub))
	// register subscriber width queue, each message is delivered to a unique subscriber
	micro.RegisterSubscriber("example.topic.pubsub.2", service.Server(), subEv, server.SubscriberQueue("queue.pubsub"))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type Sub struct {

}

/// Method can be of any name
func (s *Sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

// Alternatively a function can be used
func subEv(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}