package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/samarthasthan/twitter-sentiments/types"
)

type Consumer interface{}

type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

func NewKafkaConsumer() *KafkaConsumer {
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
	}
}

func (c *KafkaConsumer) Consume() {
	c.Consumer.SubscribeTopics([]string{"analyser"}, nil)
	for {
		msg, err := c.Consumer.ReadMessage(time.Second)
		if err == nil {

			// Consume Sentiments result fromm Kafka
			var result types.SentimentResult
			json.Unmarshal(msg.Value, &result) // Marshal Kafka Message to Sentiments result struct

			log.Printf("%v\n", result)

		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
