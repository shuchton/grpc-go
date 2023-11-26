package main

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	pb "github.com/shuchton/grpc-go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient, names ...string) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{}

	for _, name := range names {
		reqs = append(reqs, &pb.GreetRequest{FirstName: name})
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
