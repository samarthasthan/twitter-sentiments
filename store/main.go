package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/samarthasthan/twitter-sentiments/consumer"
	"github.com/samarthasthan/twitter-sentiments/database"
	"github.com/samarthasthan/twitter-sentiments/handler"
	pb "github.com/samarthasthan/twitter-sentiments/proto"
	"github.com/samarthasthan/twitter-sentiments/types"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBUsername string
	DBPassword string
	DBName     string
	DBPort     string
)

const (
	GrpcPort = ":50051"
)

func init() {
	DBUsername = os.Getenv("DB_USERNAME")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")
}

func main() {
	// Use a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Initialize the database and start the Kafka consumer
	dsn := fmt.Sprintf("%s:%s@tcp(storedb:%s)/%s", DBUsername, DBPassword, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&types.SentimentResult{})

	mysqlDB := database.NewMySqlDB(db)

	// Start gRPC server in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		startGRPCServer(mysqlDB)
	}()

	consumer := consumer.NewKafkaConsumer(mysqlDB)
	consumer.Consume()

	// Wait for all goroutines to finish
	wg.Wait()
}

func startGRPCServer(db *database.MySqlDB) {
	lis, err := net.Listen("tcp", GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	ts := handler.NewTweetGrpcServer(db)

	pb.RegisterTweetServiceServer(s, ts)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
