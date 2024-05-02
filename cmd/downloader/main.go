package main

import (
	"configs"
	"flag"
	"fmt"
	"log"
	"os"
	queue "rabbitqueue"
	"strconv"
	"sync"

	"db"
	"emailapi"
	"models"

	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

var (
	DBhost     string
	DBport     int
	DBusername string
	DBpassword string
	DBname     string
	MQhost     string
	MQport     int
	MQusername string
	MQpassword string
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Command Line Option To Set Server Gracefuls Shutdown Timeout
	MYSQL_PORT, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	flag.StringVar(&DBhost, "db-host", os.Getenv("MYSQL_HOST"), "database host domain/ip - e.g. localhost or 0.0.0.0")
	flag.IntVar(&DBport, "db-port", MYSQL_PORT, "database port number - e.g. 3306")
	flag.StringVar(&DBusername, "db-username", os.Getenv("MYSQL_USERNAME"), "database user name - e.g. admin or root")
	flag.StringVar(&DBpassword, "db-password", os.Getenv("MYSQL_PASSWORD"), "database user secret/password")
	flag.StringVar(&DBname, "database", os.Getenv("DB_USE"), "database user secret/password")

	// Command line option for rabbit message que
	RABBIT_PORT, _ := strconv.Atoi(os.Getenv("RABBIT_PORT"))
	flag.StringVar(&MQhost, "mq-host", os.Getenv("RABBIT_HOST"), "RabbitMQ host domain/ip - e.g. localhost or 0.0.0.0")
	flag.IntVar(&MQport, "mq-port", RABBIT_PORT, "RabbitMQ port number - e.g. 3306")
	flag.StringVar(&MQusername, "mq-username", os.Getenv("RABBIT_USERNAME"), "RabbitMQ user name - e.g. admin or root")
	flag.StringVar(&MQpassword, "mq-password", os.Getenv("RABBIT_PASSWORD"), "RabbitMQ user secret/password")
	flag.Parse()

	// prepare cities & lat,long data
	citiesArray := getCitiesWithLatLong()

	//  init database connection
	dbConnection := db.InitDBConnection(DBhost, DBport, DBusername, DBpassword)
	defer dbConnection.SqlCon.Close()

	// Create Required DB
	err = createRequiredTables(dbConnection)
	if err != nil {
		Log.Error(err)
	}

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

	// Print cities
	for _, city := range citiesArray {
		wg.Add(1)
		openWeatherAPI := &OpenWeatherAPI{}
		go openWeatherAPI.SaveData(city.Latitude, city.Longitude, dbConnection, channel)
	}

	wg.Wait()

	msgs, err := queue.MessageConsume(channel.MQChan, q.MQueue)
	if err != nil {
		Log.Errorf("%s: %s", "Failed to register a consumer", err)
	}

	go func() {
		for d := range msgs {
			emailapi.EWG.Add(1)
			go emailapi.SendMail(d.Body)
			log.Printf(" [x] %s", d.Body)
		}
	}()

	emailapi.EWG.Wait()
}

func createRequiredTables(cs *db.SQLConnection) error {
	// Verify Table Exist
	var weatherDataTable []models.WeatherData
	tableNames := []string{"weather_data_history", "weather_data"}
	createErr := db.CreateTable(cs.GormConn, tableNames, &weatherDataTable)
	if createErr != nil {
		return createErr
	}
	return nil
}

func getCitiesWithLatLong() (returnCity []City) {
	jsonParser := &CityJSONParser{}
	jsonData := []byte(configs.CitiesJson)

	// Parse JSON using the JSON parser
	cities, err := jsonParser.ParseJSON(jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	return cities
}
