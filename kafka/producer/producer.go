package producer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

const TCP = "localhost:9092"

func ProducerFunc() {
	writer := &kafka.Writer{
		Addr:  kafka.TCP(TCP),
		Topic: "quickstart",
	}
	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("message"),
		Headers: []protocol.Header{{
			Key:   "session",
			Value: []byte{123},
		}},
	})
	if err != nil {
		log.Fatal("Cannot write messages", err)
	}

}
