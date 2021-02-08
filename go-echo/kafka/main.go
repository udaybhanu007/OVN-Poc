package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go-echo-poc/app/dao"
	"go-echo-poc/app/datasources/cassandra/users_db"
	"go-echo-poc/config"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "test-cassandra"
	brokerAddress = "localhost:9092"
)

var messages []kafka.Message

var profile = ""
var configServer = "http://localhost:8090"

func init() {
	name := flag.String("name", "echo-cassandra-demo", "Application name")
	profileTerminal := flag.String("profile", "", "Configuration profile URL")
	configServerTerminal := flag.String("config", "http://localhost:8090/", "Config server base url")
	flag.Parse()
	if profileTerminal != nil {
		profile = *profileTerminal
	}
	if configServerTerminal != nil {
		configServer = *configServerTerminal
	}
	if len(profile) == 0 {
		fmt.Println("profile flag is empty")
		os.Exit(1)
	}
	config.ApplicationName = *name
	config.ConfigProfile = profile
	config.ConfigServer = configServer
	config, configErr := config.LoadConfiguration()
	if configErr != nil {
		fmt.Println(configErr)
		os.Exit(1)
	}
	users_db.ConnectDB(config)
}
func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	//go produce(ctx)
	consume(ctx)
}

func produce(ctx context.Context) {
	// initialize a counter
	i := 0

	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	for {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)

		if i > 10 {
			break
		}
	}
}

func consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	// l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
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

		var logMsg dao.Log
		err2 := json.Unmarshal(msg.Value, &logMsg)
		if err2 != nil {
			fmt.Println(err2)
		}
		dbSaveErr := dao.LogDaoService.Save(&logMsg)

		if dbSaveErr != nil {
			fmt.Println("error saving log:", dbSaveErr)
		}
	}
}
