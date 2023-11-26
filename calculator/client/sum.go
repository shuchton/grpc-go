package main

import (
	"context"
	"log"

	pb "github.com/shuchton/grpc-go/calculator/proto"
)

func doSum(c pb.CalcServiceClient, val1, val2 int32) {
	log.Println("doSum was invoked")

	res, err := c.Add(context.Background(), &pb.SumRequest{
		Val1: val1,
		Val2: val2,
	})

	if err != nil {
		log.Fatalf("could not sum: %v\n", err)
	}

	log.Printf("%d + %d = %d", val1, val2, res.Result)
}
