package main

import (
	"log"
	"net"

	pb "github.com/shuchton/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Printf("listening on %s\n", addr)

	opts := []grpc.ServerOption{}

	tls := true
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}

}
