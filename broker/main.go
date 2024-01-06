package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samarthasthan/twitter-sentiments/handler"
	pb "github.com/samarthasthan/twitter-sentiments/proto"
)

func main() {
	grpc := handler.NewGrpcHandler()
	app := fiber.New()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		grpc.Initialise()
	}()

	wg.Wait()

	app.Get("/tweets", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := grpc.Client.TweetsHandler(ctx, &pb.Pagination{Limit: 10, Offset: 1})
		if err != nil {
			fmt.Println(err)
			return c.SendString(err.Error())
		}

		return c.JSON(res)
	})

	app.Listen(":8000")

	defer grpc.Close()
}
