package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/op/go-logging"
	amqp "github.com/rabbitmq/amqp091-go"
)

var Log = logging.MustGetLogger("API")

type IQueue interface {
	ConnectMQ() (c *amqp.Connection, e error)
	CreateMQChannel(cs *QueueConnection) *QueueConnection
	CloseMQ() (err error)
	QueueConnect(host string, port int, username string, password string) *QueueConnection
	DefineExchange(ch *QueueConnection, topic string)
	PublishMessage(ch *amqp.Channel, ctx context.Context) (e error)
	DefineQueue(ch *amqp.Channel) *QueueConnection
}

type QueueConnection struct {
	host     string
	port     int
	username string
	password string
	MQCon    *amqp.Connection
	MQChan   *amqp.Channel
	MQueue   *amqp.Queue
}

func QueueConnect(host string, port int, username string, password string) *QueueConnection {
	return &QueueConnection{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (cs *QueueConnection) ConnectMQ() (c *amqp.Connection, e error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", cs.username, cs.password, cs.host, cs.port)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	return conn, err
}

func CloseMQ(mq *amqp.Connection) (err error) {
	return mq.Close()
}

func InitQueueConnection(host string, port int, username string, password string) *QueueConnection {
	// Initialize MySQL connector
	connection := QueueConnect(host, port, username, password)

	// Connect to MySQL
	mb, err := connection.ConnectMQ()
	if err != nil {
		Log.Error(err)
	}

	return &QueueConnection{
		MQCon: mb,
	}
}

func CreateMQChannel(cs *amqp.Connection) *QueueConnection {
	ch, err := cs.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
	return &QueueConnection{
		MQChan: ch,
	}
}

func (ch *QueueConnection) CloseChannel() (err error) {
	return ch.MQChan.Close()
}

func DefineExchange(ch *amqp.Channel, topic string) {
	err := ch.ExchangeDeclare(
		topic,   // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
}

func PublishMessage(ch *amqp.Channel, ctx context.Context, topic string, body string) (e error) {
	err := ch.PublishWithContext(ctx,
		topic,            // exchange
		"anonymous.info", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	return err
}

func DefineQueue(ch *amqp.Channel) *QueueConnection {
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Panicf("%s: %s", "Failed to declare a queue", err)
	}

	return &QueueConnection{
		MQueue: q,
	}
}
