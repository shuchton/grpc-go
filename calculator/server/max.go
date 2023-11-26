package main

import (
	"errors"
	"io"
	"log"

	pb "github.com/shuchton/grpc-go/calculator/proto"
)

var currentMax int32 = 0

func (s *Server) Max(stream pb.CalcService_MaxServer) error {
	log.Println("Max was invoked")

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		if req.Input > currentMax {
			currentMax = req.Input
			err = stream.Send(&pb.MaxResponse{
				Result: currentMax,
			})
			if err != nil {
				log.Fatalf("error while sending data to the client: %v\n", err)
			}
		}
	}
}
