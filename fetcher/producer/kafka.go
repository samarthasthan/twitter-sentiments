package producer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/samarthasthan/twitter-sentiments/types"
)

type DataProducer interface {
	Produce(tweet *types.Tweet)
}

type KafkaProducer struct {
	Producer *kafka.Producer
}

func NewKafkaProducer() *KafkaProducer {
	kafkaPort := os.Getenv("KAFKA_INTERNAL_PORT")
	kafkaUrl := fmt.Sprintf("kafka:%s", kafkaPort)
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaUrl})
	if err != nil {
		panic(err)
	}
	return &KafkaProducer{
		Producer: p,
	}
}

func (k *KafkaProducer) Produce(tweet *types.Tweet) {
	topic := "fetcher"
	data, err := json.Marshal(tweet)
	if err != nil {
		log.Fatalln(err)
	}
	k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}, nil)
}
