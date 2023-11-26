package main

import (
	"errors"
	"io"
	"log"

	pb "github.com/shuchton/grpc-go/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("error while sending data to the client: %v\n", err)
		}
	}
}
