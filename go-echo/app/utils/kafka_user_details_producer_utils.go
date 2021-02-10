package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"go-echo-poc/app/model"

	"github.com/segmentio/kafka-go"
)

const (
	topicForUserCreate         = "user_create"
	brokerAddressForUserCreate = "localhost:9092"
)

var ctx = context.Background()

func ProduceKafkaForUserCreate(user *model.User) {
	fmt.Print("entered producer")
	msgBytes, marshalingErr := json.Marshal(user)
	if marshalingErr != nil {
		panic("could not write message " + marshalingErr.Error())
	}
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddressForUserCreate},
		Topic:   topicForUserCreate,
	})
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(user.FirstName),
		// create an arbitrary message payload for the value
		Value: []byte(msgBytes),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
	fmt.Print("produced successfully!")
}
