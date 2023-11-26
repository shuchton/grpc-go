package main

import (
	"context"
	"errors"
	"io"
	"log"

	pb "github.com/shuchton/grpc-go/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("ListBlogs was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error while calling list blogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("something happened while reading the stream: %v\n", err)
		}

		log.Println(res)
	}

}
