package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// prepare cities & lat,long data
	citiesArray := getCitiesWithLatLong()

	//  init database connection
	dbConnection := initDBConnection()
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

func createRequiredTables(dbConnection *SQLConnection) error {
	// Verify Table Exist
	var weatherDataTable []WeatherData
	tableNames := []string{"weather_data_history", "weather_data"}
	createErr := CreateTable(dbConnection.GormConn, tableNames, &weatherDataTable)
	if createErr != nil {
		return createErr
	}
	return nil
}

func initDBConnection() *SQLConnection {
	// Initialize MySQL connector
	connector := MySQLConnect("localhost", 3306, "root", "Pankaj@569", "open_weather")

	// Connect to MySQL
	ormdb, dbConn, err := connector.Connect()
	if err != nil {
		Log.Fatal("Error connecting to MySQL:", err)
	}

	return &SQLConnection{
		GormConn: ormdb,
		SqlCon:   dbConn,
	}
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
