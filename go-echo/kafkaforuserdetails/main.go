package kafkaforuserdetails

import "context"

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	//go produce(ctx)
	consumeKafkaForUserCreate(ctx)
}
