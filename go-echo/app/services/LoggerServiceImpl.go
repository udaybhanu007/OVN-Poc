package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

type Hostname struct {
	name string
	err  string
}

// Activity Struct for Log
type Activity struct {
	UUID       string
	Action     string
	StatusCode string
	Message    string
}

const (
	topic         = "test-cassandra"
	brokerAddress = "localhost:9092"
)

var ctx = context.Background()
var messages []kafka.Message

func LogActivity(uuid string, action string, statusCode string, message string) {
	l := log.New(os.Stdout, message, 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	fmt.Println("UUID:", uuid)
	fmt.Println("Action:", action)
	fmt.Println("StatusCode", statusCode)
	fmt.Println("Msg:", message)

	value := &Activity{uuid, action, statusCode, message}

	jsonValue, _ := json.Marshal(value)
	fmt.Println(string(jsonValue))

	w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(statusCode),
		Value: []byte(jsonValue),
		Time:  time.Now(),
	})

}

func AnalyzeActivities() {
	fmt.Println("AnalyzeActivities started...")
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	// l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
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
		// fmt.Println("messages:", messages)
	}
}
