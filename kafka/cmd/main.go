package main

import (
	"github.com/portifolio/kafka/consumer"
	"github.com/portifolio/kafka/producer"
)

func main() {
	producer.ProducerFunc()
	consumer.ConsumerFunc()

}
