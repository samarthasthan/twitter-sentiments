package producer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/samarthasthan/twitter-sentiments/types"
)

type Producer interface {
	Produce(*types.SentimentResult)
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

func (k *KafkaProducer) Produce(result *types.SentimentResult) {
	topic := "analyser"
	data, err := json.Marshal(result)
	if err != nil {
		log.Fatalln(err)
	}
	k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}, nil)
}
