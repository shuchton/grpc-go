package main

import (
	pb "github.com/shuchton/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"author_id,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Content  string             `bson:"content,omitempty"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}

func (b *BlogItem) ToPb() *pb.Blog {
	return &pb.Blog{
		Id:       b.ID.Hex(),
		AuthorId: b.AuthorId,
		Title:    b.Title,
		Content:  b.Content,
	}
}
