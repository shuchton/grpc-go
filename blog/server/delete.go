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
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("DeleteBlog invoked")
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid id",
		)
	}

	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"unable to delete blog post",
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("blog with id (%s) not found\n", in.Id),
		)
	}

	log.Println("blog deleted: ", in.Id)

	return &emptypb.Empty{}, nil

}
