package main

import (
	"context"
	"errors"
	"io"
	"log"

	pb "github.com/shuchton/grpc-go/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient, firstName string) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: firstName,
	}

	stream, err := c.GreetManytimes(context.Background(), req)

	if err != nil {
		log.Fatalf("could not greet many times: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v\n", err)
		}

		log.Printf("greeting many times: %s", msg)
	}

}
