package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/shuchton/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with: %v\n", in)
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		log.Printf("invalid id: %s\n", in.Id)
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("invalid id: %s\n", in.Id),
		)
	}

	var data BlogItem
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(&data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"cannot find the blog with id provided",
		)
	}

	return data.ToPb(), nil

}
