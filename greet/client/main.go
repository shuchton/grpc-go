package main

import (
	"log"

	pb "github.com/shuchton/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("error while loading CA trust certificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)

	// doGreetManyTimes(c, "Scott")

	// doLongGreet(c)

	//doGreetEveryone(c, "Scott", "Starla")

	// doGreetWithDeadline(c, 1*time.Second)
}
