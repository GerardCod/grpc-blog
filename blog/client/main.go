package main

import (
	"log"

	pb "github.com/GerardCod/grpc-blog/blog/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	doReadBlog(c, id)
	//doReadBlog(c, "aNonExistingId")
	updateBlog(c, id)
	listBlogs(c)
	deleteBlog(c, id)
}
