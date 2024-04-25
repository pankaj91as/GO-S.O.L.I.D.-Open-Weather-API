package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/pankaj91as/open-weather-api/common/models"
	"github.com/pankaj91as/open-weather-api/pkg/db"
)

var wg sync.WaitGroup

var (
	DBhost	string
	DBport	int
	DBusername	string
	DBpassword	string
	DBname	string
)

func main() {
	// Command Line Option To Set Server Gracefuls Shutdown Timeout
	flag.StringVar(&DBhost, "db-host", "localhost", "database host domain/ip - e.g. localhost or 0.0.0.0")
	flag.IntVar(&DBport, "db-port", 3306, "database port number - e.g. 3306")
	flag.StringVar(&DBusername, "db-username", "root", "database user name - e.g. admin or root")
	flag.StringVar(&DBpassword, "db-password", "Pankaj@569", "database user secret/password")
	flag.StringVar(&DBname, "database", "open_weather", "database user secret/password")
	flag.Parse()

	// prepare cities & lat,long data
	citiesArray := getCitiesWithLatLong()

	//  init database connection
	dbConnection := db.InitDBConnection(DBhost, DBport, DBusername, DBpassword)
	defer dbConnection.SqlCon.Close()

	// Create Required DB
	err := createRequiredTables(dbConnection)
	if err != nil {
		Log.Fatal(err)
	}

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

	// Open JSON file
	jsonFile, err := os.Open("cities.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Read JSON File
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	// Parse JSON using the JSON parser
	cities, err := jsonParser.ParseJSON(jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	return cities
}
