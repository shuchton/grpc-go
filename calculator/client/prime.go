package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	pb "github.com/shuchton/grpc-go/calculator/proto"
)

func doPrime(c pb.CalcServiceClient, input int32) {
	log.Printf("doPrime was involed with %d\n", input)

	req := &pb.PrimeRequest{
		Input: input,
	}

	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("could not do prime sieve: %v\n", err)
	}

	var res []int32

	for {
		msg, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v\n", err)
		}

		res = append(res, msg.Result)
	}

	log.Printf("The prime number decomposition of %d is %s\n", input, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), ","), "[]"))

}
