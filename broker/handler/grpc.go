package handler

import (
	"log"

	pb "github.com/samarthasthan/twitter-sentiments/proto"
	"google.golang.org/grpc"
)

const (
	address = "store:50051"
)

type GrpcHandler struct {
	Conn   *grpc.ClientConn
	Client pb.TweetServiceClient
}

func NewGrpcHandler() *GrpcHandler {
	return &GrpcHandler{}
}

func (g *GrpcHandler) Initialise() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("GRPC connected!")

	g.Conn = conn

	c := pb.NewTweetServiceClient(g.Conn)

	g.Client = c
}

func (g *GrpcHandler) Close() {
	if g.Conn != nil {
		g.Conn.Close()
		log.Println("GRPC connection closed")
	}
}
