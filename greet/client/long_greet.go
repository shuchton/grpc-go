package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shuchton/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Scott"},
		{FirstName: "Starla"},
		{FirstName: "Colleen"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error reading server stream: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet:\n%s\n", res.Result)
}
