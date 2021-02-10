package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-echo-poc/app/dao"

	"github.com/segmentio/kafka-go"
)

const (
	topicForUserCreate         = "user_create"
	brokerAddressForUserCreate = "localhost:9092"
)

func main() {
	ctx := context.Background()
	consumeKafkaForUserCreate(ctx)
}

var ctx = context.Background()
var messages []kafka.Message

func consumeKafkaForUserCreate(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddressForUserCreate},
		Topic:   topicForUserCreate,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
		messages = append(messages, msg)

		var logMsg dao.Log
		err2 := json.Unmarshal(msg.Value, &logMsg)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}
