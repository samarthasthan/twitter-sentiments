package consumer

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/samarthasthan/twitter-sentiments/producer"
	"github.com/samarthasthan/twitter-sentiments/sentiments"
	"github.com/samarthasthan/twitter-sentiments/types"
)

type Consumer interface{}

type KafkaConsumer struct {
	Consumer *kafka.Consumer
	Producer producer.Producer
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

func (c *KafkaConsumer) Consume(analyser *sentiments.Analyser) {
	c.Consumer.SubscribeTopics([]string{"fetcher"}, nil)
	for {
		msg, err := c.Consumer.ReadMessage(time.Second)
		if err == nil {

			// Consume Tweet fromm Kafka
			var tweet types.Tweet
			json.Unmarshal(msg.Value, &tweet) // Marshal Kafka Message to Tweet Struct

			// Analyse Sentiment from tweet content and returns score
			score := analyser.Analyse(tweet.Content)

			result := types.SentimentResult{Username: tweet.Username, Content: tweet.Content, Score: score}
			c.Producer.Produce(&result)

		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
