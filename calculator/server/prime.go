package main

import (
	"log"

	pb "github.com/shuchton/grpc-go/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalcService_PrimeServer) error {
	log.Printf("prime invoked with %v\n", in)
	var k int32 = 2
	n := in.Input

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n = n / k
		} else {
			k += 1
		}
	}

	return nil
}
