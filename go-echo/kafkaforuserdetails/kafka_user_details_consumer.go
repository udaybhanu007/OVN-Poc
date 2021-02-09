package kafkaforuserdetails

import (
	"context"
	"encoding/json"
	"fmt"
	"go-echo-poc/app/domain"

	"github.com/segmentio/kafka-go"
)

var messages []kafka.Message

func consumeKafkaForUserCreate(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddressForUserCreate},
		Topic:   topicForUserCreate,
		// GroupID: "my-group",
		// assign the logger to the reader
		// Logger: l,
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

		var logMsg domain.Log
		err2 := json.Unmarshal(msg.Value, &logMsg)
		if err2 != nil {
			fmt.Println(err2)
		}
		dbSaveErr := domain.LogDaoService.Save(&logMsg)

		if dbSaveErr != nil {
			fmt.Println("error saving log:", dbSaveErr)
		}
	}
}
