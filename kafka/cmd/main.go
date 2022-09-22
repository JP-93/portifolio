package main

import (
	"context"

	"kafka/v2/producer"
)

func main() {
	ctx := context.Background()
	go producer.StartProducer(ctx)

}
