package main

import (
	"context"
	proto "d7kj.com/m/v2/examples/pubsub/srv/proto"
	"fmt"
	"github.com/google/uuid"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"time"
)

func main()  {
	// create a service
	service := micro.NewService(
			micro.Name("go.micro.cli.pubsub"),
		)
	service.Init()

	// create publisher
	pub1 := micro.NewPublisher("example.topic.pubsub.1", service.Client())
	pub2 := micro.NewPublisher("example.topic.pubsub.2", service.Client())

	// pub to topic 1
	go sendEv("example.topic.pubsub.1", pub1)
	// pub to topic 2
	go sendEv("example.topic.pubsub.2", pub2)
	// block forever
	select {}
}

// send events using tht publisher
func sendEv(topic string, p micro.Publisher)  {
	t := time.NewTicker(time.Second)

	for _ = range t.C {
		// create new event
		_uuid, _ := uuid.NewUUID()
		ev := &proto.Event{
			Id: _uuid.String(),
			Timestamp: time.Now().Unix(),
			Message: fmt.Sprintf("Message you all day on %s", topic),
		}
		log.Logf("publishing %+v\n", ev)
		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil{
			log.Logf("error publishing %v\n", ev)
		}
	}
}
