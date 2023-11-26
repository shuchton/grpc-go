package main

import (
	"log"
	"net"

	pb "github.com/shuchton/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalcServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Printf("listening on: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v\n", err)
	}

}
