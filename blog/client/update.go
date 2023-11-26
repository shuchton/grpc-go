package main

import (
	"context"
	"log"

	pb "github.com/shuchton/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Starla",
		Title:    "A New Title",
		Content:  "That old content was way too long. This has been updated to be more readable",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("error updating blog with id: %s - %v\n", id, err)
	}

	log.Println("blog was updated")
}
