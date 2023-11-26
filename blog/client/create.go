package main

import (
	"context"
	"log"

	pb "github.com/shuchton/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog invoked")
	blog := &pb.Blog{
		AuthorId: "Scott",
		Title:    "First Blog Post",
		Content:  "Lorem ipsum dolor sit amet consectetur adipisicing elit. Expedita ullam numquam delectus quod rerum quam nihil, reprehenderit facilis pariatur amet ea, ab minima modi corrupti ipsam. Deserunt non neque est.",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}
