*Prerequisite: Need RabbitMQ server installed & running on machine.

We will be creating a consumer service, which will subscribe to our topics and we will define a sender-service, which will publish random events to the exchange. Our lib folder, will hold some common configurations for both our consumer and sender. Before we begin, you will have to get the dependency for amqp:

go get github.com/streadway/amqp

But that's it, now we are ready to write some code.

I. Event Queue:
All files in this section will be placed in lib/event.

1. ./lib/event/event.go
In this file, we are defining three static methods. The getExchangeName function simply returns the name of our exchange. It isn't necessary, but nice for this tutorial, to make it simple to change your topic name. More interesting is the declareRandomQueue function. This function will create a nameless queue, which RabbitMQ will assign a random name, we don't want to worry about this and that is why we are letting RabbitMQ worry about it. The queue is also defined as exclusive, which means that when defined only one subscriber can be subscribed to this queue. The last function that we have declared is declareExchange which will declare an exchange, as the name suggests. This function is idempotent, so if the exchange already exists, no worries, it won't create duplicates. However, if we were to change the type of the Exchange (to direct or fanout), then we would have to either delete the old exchange or find a new name, as you cannot overwrite exchanges. The topic type is what enables us to publish an event with a topic such as log.WARN, which the subscribers can specify in their binding keys.

NOTE: You might have noticed that both functions need an amqp.Channel struct. This is simply a pointer to an AMQP connection channel. We will explain this a little better later

2. ./lib/event/emitter.go
In this file, at the very top of our code, we are defining our Emitter struct (a class), which contains an amqp.Connection.

setup - Makes sure that the exchange that we are sending messages to actually exists. We do this by retrieving a channel from our connection pool and calling the idempotent declareExchange function from our event.go file.

Push - Sends a message to our exchange. First we get a new channel from our connection pool and if we receive no errors when doing so, we publish our message. The function takes two input parameters event and severity; event is the message to be sent and severity is our logging serverity, which will define which messages are received by which subscribers, based on their binding keys.

NewEventEmitter - Will return a new Emitter, or an error, making sure that the connection is established to our AMQP server.

The last bit of code to write for our library, is our consumer struct and right away we can see that it is somewhat similar to our emitter struct.

3. ./lib/event/consumer.go
In this file, at the very top we define that our Consumer struct defines a connection to our AMQP server and a queueName. The queue name will store the randomly generated name of our declared nameless queue. We will use this for telling RabbitMQ that we want to bind/listen to this particular queue for messages.

setup() - We ensure that the exchange is declared, just like we do in our Emitter struct.

NewConsumer() - We return a new Consumer or an error, ensuring that everything went well connecting to our AMQP server.

Listen - We get a new channel from our connection pool. We declare our nameless queue and then we iterate over our input topics, which is just an array of strings, specifying our binding keys. For each string in topics, we will bind our queue to the exchange, specifying our binding key, for which messages we want to receive. As an example, this could be log.WARN and log.ERROR. Lastly, we will invoke the Consume function (to start listening on the queue) and define that we will iterate over all messages received from the queue and print out these message to the console.

The forever channel that we are making on line #69, and sending output from on line #77, is just a dummy. This is a simple way of ensuring a program will run forever. Essentially, we are defining a channel, which we will wait for until it receives input, but we will never actually send it any input. It's a bit dirty, but for this tutorial it will suffice.

II. Consumer Service:
All files in this section will be placed in the consumer folder.

1. ./consumer/consumer.go
As can be seen this is a really simple program which creates a connection to our docker instance of RabbitMQ, passes this connection to our NewConsumer function and then calls the Listen method, passing all the input arguments from the command line. Once we have written this code we can open up a few terminals to start up a few consumers:

#t1> go run consumer.go queueOne log.WARN log.ERROR

#t2> go run consumer.go queueTwo log.*

The first terminal in which we are running our consumer.go file, we are listening for all log.WARN and log.ERROR events. In the second terminal we are listening for all events. It is also possible to do a lot of other search filters with binding keys. There are only two different kind of binding keys * and #:

* substitutes exactly one word. So our binding key could be: apples.*.orangeand we would receive apples.dingdong.orange. Similarly, we would receive log.WARN if our binding was log.*, but we wouldn't receive log.WARN.joke #: substitutes zero or more words. So if we use the same example as above: If our binding is log.# we will receive log.WARN.joke as well as receiving log.WARN.

III. Emitter Service:

1. ./sender/sender.go
Again, a very simply little service. Connection to AMQP, create a new Event Emitter and then iterate to publish 10 messages to the exchange, using the console input as severity level. The Push function being input (message: "i - input", severity: input). Simples. So, run this a few times and see what happens:

#t3> go run sender.go log.WARN

#t3> go run sender.go log.ERROR

#t3> go run sender.go log.INFO

Wow! As expected our two other services are now receiving messages independently of each other, only receiving the messages that they are subscribed to.

