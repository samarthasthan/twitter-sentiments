package main

import "github.com/samarthasthan/twitter-sentiments/consumer"

func main() {
	consumer := consumer.NewKafkaConsumer()
	consumer.Consume()
}
