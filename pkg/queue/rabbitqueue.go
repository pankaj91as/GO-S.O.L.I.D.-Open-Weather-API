package queue

import (
	"fmt"
	"log"

	"github.com/op/go-logging"
	amqp "github.com/rabbitmq/amqp091-go"
)

var Log = logging.MustGetLogger("API")

type IQueue interface {
	ConnectMQ() (c *amqp.Connection, e error)
	CreateMQChannel(cs *amqp.Connection) (c *amqp.Channel, e error)
	CloseMQ(mq *amqp.Connection) (err error)
	QueueConnect(host string, port int, username string, password string) *QueueConnection
}

type QueueConnection struct {
	host     string
	port     int
	username string
	password string
	MQCon    *amqp.Connection
	MQChan   *amqp.Channel
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

	// Create MQ Channel
	ch, err := CreateMQChannel(mb)
	if err != nil {
		Log.Error(err)
	}

	return &QueueConnection{
		MQCon:  mb,
		MQChan: ch,
	}
}

func CreateMQChannel(cs *amqp.Connection) (c *amqp.Channel, e error) {
	ch, err := cs.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
	return ch, err
}

func CloseChannel(ch *amqp.Connection) (err error) {
	return ch.Close()
}
