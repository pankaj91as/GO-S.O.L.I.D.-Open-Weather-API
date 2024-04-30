package main

import (
	"configs"
	"flag"
	"fmt"
	queue "rabbitqueue"
	"sync"

	"db"
	"models"
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
	// Command Line Option To Set Server Gracefuls Shutdown Timeout
	flag.StringVar(&DBhost, "db-host", "0.0.0.0", "database host domain/ip - e.g. localhost or 0.0.0.0")
	flag.IntVar(&DBport, "db-port", 3306, "database port number - e.g. 3306")
	flag.StringVar(&DBusername, "db-username", "root", "database user name - e.g. admin or root")
	flag.StringVar(&DBpassword, "db-password", "password", "database user secret/password")
	flag.StringVar(&DBname, "database", "open_weather", "database user secret/password")

	// Command line option for rabbit message que
	flag.StringVar(&MQhost, "mq-host", "localhost", "RabbitMQ host domain/ip - e.g. localhost or 0.0.0.0")
	flag.IntVar(&MQport, "mq-port", 5672, "RabbitMQ port number - e.g. 3306")
	flag.StringVar(&MQusername, "mq-username", "guest", "RabbitMQ user name - e.g. admin or root")
	flag.StringVar(&MQpassword, "mq-password", "guest", "RabbitMQ user secret/password")
	flag.Parse()

	// prepare cities & lat,long data
	citiesArray := getCitiesWithLatLong()

	//  init database connection
	dbConnection := db.InitDBConnection(DBhost, DBport, DBusername, DBpassword)
	defer dbConnection.SqlCon.Close()

	// Create Required DB
	err := createRequiredTables(dbConnection)
	if err != nil {
		Log.Error(err)
	}

	//  init Rabbit Queue connection
	mqConnection := queue.InitQueueConnection(MQhost, MQport, MQusername, MQpassword)
	defer mqConnection.MQCon.Close()

	// Print cities
	for _, city := range citiesArray {
		wg.Add(1)
		openWeatherAPI := &OpenWeatherAPI{}
		go openWeatherAPI.SaveData(city.Latitude, city.Longitude, dbConnection)
	}

	wg.Wait()
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
