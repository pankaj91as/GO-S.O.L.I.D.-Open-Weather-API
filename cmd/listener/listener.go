package main

import (
	"flag"
	"log"
	"os"
	queue "rabbitqueue"
	"strconv"

	"emailapi"

	"github.com/joho/godotenv"
	"github.com/op/go-logging"
)

var (
	MQhost     string
	MQport     int
	MQusername string
	MQpassword string
)

var Log = logging.MustGetLogger("listener")

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Command line option for rabbit message que
	RABBIT_PORT, _ := strconv.Atoi(os.Getenv("RABBIT_PORT"))
	flag.StringVar(&MQhost, "mq-host", os.Getenv("RABBIT_HOST"), "RabbitMQ host domain/ip - e.g. localhost or 0.0.0.0")
	flag.IntVar(&MQport, "mq-port", RABBIT_PORT, "RabbitMQ port number - e.g. 3306")
	flag.StringVar(&MQusername, "mq-username", os.Getenv("RABBIT_USERNAME"), "RabbitMQ user name - e.g. admin or root")
	flag.StringVar(&MQpassword, "mq-password", os.Getenv("RABBIT_PASSWORD"), "RabbitMQ user secret/password")
	flag.Parse()

	//  init Rabbit Queue connection
	MqConnection := queue.InitQueueConnection(MQhost, MQport, MQusername, MQpassword)

	// Close Message Queue Connection
	// defer MqConnection.MQCon.Close()

	// Create MQ Channel
	channel := queue.CreateMQChannel(MqConnection.MQCon)

	// Create Exchange For Message Queue
	err = queue.DefineExchange(channel.MQChan, os.Getenv("MQ_TOPIC"))
	if err != nil {
		Log.Errorf("%s: %s", "Failed to open a channel", err)
	}

	// Define Queue
	q := queue.DefineQueue(channel.MQChan)

	// Bind Queue
	err = queue.BindQueue(channel.MQChan, q.MQueue, os.Getenv("MQ_TOPIC"))
	if err != nil {
		Log.Errorf("%s: %s", "Failed to bind a queue", err)
	}

	msgs, err := queue.MessageConsume(channel.MQChan, q.MQueue)
	if err != nil {
		Log.Errorf("%s: %s", "Failed to register a consumer", err)
	}

	var forever chan struct{}

	for d := range msgs {
		emailapi.EWG.Add(1)
		go emailapi.SendMail(d.Body)
	}

	emailapi.EWG.Done()

	<-forever
}
