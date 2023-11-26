package main

import (
	"context"
	"log"
	"net"

	pb "github.com/shuchton/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("error connecting to MongoDb: %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Printf("listening on: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v\n", err)
	}

}
