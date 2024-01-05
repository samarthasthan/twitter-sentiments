package consumer

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/samarthasthan/twitter-sentiments/database"
	"github.com/samarthasthan/twitter-sentiments/types"
)

type Consumer interface{}

type KafkaConsumer struct {
	Consumer *kafka.Consumer
	DB       *database.MySqlDB
}

func NewKafkaConsumer(db *database.MySqlDB) *KafkaConsumer {
	kafkaPort := os.Getenv("KAFKA_INTERNAL_PORT")
	kafkaUrl := fmt.Sprintf("kafka:%s", kafkaPort)
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaUrl,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Println(err)
	}
	return &KafkaConsumer{
		Consumer: c,
		DB:       db,
	}
}

func (c *KafkaConsumer) Consume() {
	c.Consumer.SubscribeTopics([]string{"analyser"}, nil)
	for {
		var tweets []*types.SentimentResult
		for i := 0; i < 10; {
			msg, err := c.Consumer.ReadMessage(time.Second * 1)
			if err == nil {

				// Consume Sentiments result fromm Kafka
				var result types.SentimentResult
				json.Unmarshal(msg.Value, &result) // Marshal Kafka Message to Sentiments result struct

				tweets = append(tweets, &result)
				i++

			} else if !err.(kafka.Error).IsTimeout() {
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
		c.DB.CreateTweet(tweets)
	}
}
