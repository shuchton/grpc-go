package main

import (
	"context"
	"log"

	pb "github.com/shuchton/grpc-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog invoked")
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("unable to delete blog: %v\n", err)
	}

	log.Println("blog was deleted")
}
