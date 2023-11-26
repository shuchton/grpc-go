package main

import (
	"errors"
	"fmt"
	"io"
	"log"

	pb "github.com/shuchton/grpc-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function invoked")
	res := ""

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
