package main

import (
	"context"
	"log"

	pb "github.com/shuchton/grpc-go/calculator/proto"
)

func (s *Server) Add(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("add function was invoked with %v\n", in)

	return &pb.SumResponse{
		Result: in.Val1 + in.Val2,
	}, nil

}
