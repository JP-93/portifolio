package producer

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic     = "message-log"
	bronker_1 = "localhost:9091"
	bronker_2 = "localhost:9092"
	bronker_3 = "localhost:9093"
)

func StartProducer(ctx context.Context) {
	//incializando um contador
	i := 0
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{bronker_1, bronker_2, bronker_3},
		Topic:   topic,
	})
	// cada mensagem kafka tem uma chave e um valor. A chave é usada
	// para decidir qual partição (e consequentemente, qual corretor)
	// a mensagem é publicada
	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			fmt.Errorf("Erro ao escrever a mensagem", err)
		}
		fmt.Println("WriteMessages", i)
		i++
		time.Sleep(time.Second)
	}

}
