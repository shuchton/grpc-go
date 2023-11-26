package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/shuchton/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs was invoked")

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknown internal error: %v", err),
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("error decoding data from mongodb: %v", err),
			)
		}

		stream.Send(data.ToPb())
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknwon internal error: %v", err),
		)
	}

	return nil
}
