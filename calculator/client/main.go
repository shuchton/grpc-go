package main

import (
	"log"

	pb "github.com/shuchton/grpc-go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalcServiceClient(conn)

	//doSum(c, 10, 3)

	//doPrime(c, 210)
	// doAvg(c)
	// doMax(c, 1, 5, 3, 6, 2, 20)
	doSqrt(c, -10)
}
