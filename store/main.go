package main

import (
	"fmt"
	"os"

	"github.com/samarthasthan/twitter-sentiments/consumer"
	"github.com/samarthasthan/twitter-sentiments/database"
	"github.com/samarthasthan/twitter-sentiments/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBUsername string
	DBPassword string
	DBName     string
	DBPort     string
)

func init() {
	DBUsername = os.Getenv("DB_USERNAME")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(storedb:%s)/%s", DBUsername, DBPassword, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&types.SentimentResult{})

	mysql := database.NewMySqlDB(db)
	consumer := consumer.NewKafkaConsumer(mysql)
	consumer.Consume()
}
