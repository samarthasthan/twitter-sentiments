package main

import (
	"github.com/samarthasthan/twitter-sentiments/consumer"
	"github.com/samarthasthan/twitter-sentiments/producer"
	"github.com/samarthasthan/twitter-sentiments/sentiments"
)

func main() {
	consumer := consumer.NewKafkaConsumer()
	producer := producer.NewKafkaProducer()
	consumer.Producer = producer
	analyser := sentiments.NewAnalyser()
	consumer.Consume(analyser)
}
