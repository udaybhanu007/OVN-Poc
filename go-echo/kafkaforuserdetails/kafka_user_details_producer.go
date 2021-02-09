package kafkaforuserdetails

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go-echo-poc/app/domain"

	"github.com/segmentio/kafka-go"
)

const (
	topicForUserCreate         = "user_create"
	brokerAddressForUserCreate = "localhost:9092"
)

var ctx = context.Background()

func ProduceKafkaForUserCreate(user *domain.User) {
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
	// sleep for a second
	time.Sleep(time.Second)
}
