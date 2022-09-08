package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func ConsumerFunc() {
	fmt.Println(" topico: ", "quickstart")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer",
		Topic:    "quickstart",
		MinBytes: 0,
		MaxBytes: 10e6,
	})
	for i := 0; i <= 1; i++ {
		message, err := reader.ReadMessage(context.Background())
		for _, val := range message.Headers {
			if val.Key == "session" && string(val.Value) == "123" {
				fmt.Println("sessão correta")
			}
		}
		if err != nil {
			log.Fatal("fail consumer", err)
			reader.Close()
		}
		fmt.Print("receiver a message: ", string(message.Value))
	}
	reader.Close()
}
