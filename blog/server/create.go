package main

import (
	"context"

	pb "github.com/GerardCod/grpc-blog/blog/proto"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {

}
