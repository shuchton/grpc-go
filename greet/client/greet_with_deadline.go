package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shuchton/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Scott",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("deadline exceeded")
			} else {
				log.Printf("unexpected grpc error: %v\n", err)
			}
		} else {
			log.Printf("non grpc error: %v\n", err)
		}
		return
	}

	log.Printf("greetWithDeadline:\n%s\n", res.Result)

}
