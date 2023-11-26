package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shuchton/grpc-go/calculator/proto"
)

func doAvg(c pb.CalcServiceClient) {
	log.Println("doAvg was invoked")

	reqs := []*pb.AvgRequest{
		{Input: 1},
		{Input: 2},
		{Input: 3},
		{Input: 4},
		{Input: 5},
		{Input: 6},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("error while calling Avg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while recieving stream: %v\n", err)
	}

	log.Printf("Average: %f", res.Result)
}
