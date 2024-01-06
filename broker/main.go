package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/samarthasthan/twitter-sentiments/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTweetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	r, err := c.TweetsHandler(ctx, &pb.Pagination{Limit: 10, Offset: 2})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%+v", r.GetTweets())

}
