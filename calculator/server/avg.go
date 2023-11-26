package main

import (
	"errors"
	"io"
	"log"

	pb "github.com/shuchton/grpc-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalcService_AvgServer) error {
	log.Println("avg function invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			var result float64 = 0
			if count > 0 {
				result = float64(sum) / float64(count)
			}
			return stream.SendAndClose(&pb.AvgResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		count += 1
		sum += req.Input
	}
}
