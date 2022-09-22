package topic

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func CreateTopic() {
	// to create topics when auto.create.topics.enable='true'
	_, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "my-topic", 0)
	if err != nil {
		panic(err.Error())
	}

}
