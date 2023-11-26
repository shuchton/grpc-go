package main

import (
	"fmt"
	"log"

	pb "github.com/shuchton/grpc-go/greet/proto"
)

func (s *Server) GreetManytimes(in *pb.GreetRequest, stream pb.GreetService_GreetManytimesServer) error {
	log.Printf("greet many times invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hellow %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
