package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type OpenWeatherAPI struct{}

func (ll *OpenWeatherAPI) FetchData(latitude float64, longitude float64) {
	// Waitgroup Done Signal
	defer wg.Done()

	// Create API Endpoint
	endpoint := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v", latitude, longitude, "3e7e4cafc4f58311c36ef0e546a7fecd")
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}

	// Read API Body Response
	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var APIResponse map[string]interface{}
	if err := json.Unmarshal(bodyData, &APIResponse); err != nil {
		fmt.Println(err)
	}

	weatherHistoryData := WeatherDataHistory{
		Lon:           APIResponse["coord"].(map[string]interface{})["lon"].(float64),
		Lat:           APIResponse["coord"].(map[string]interface{})["lat"].(float64),
		MainTemp:      APIResponse["main"].(map[string]interface{})["temp"].(float64),
		MainTempMin:   APIResponse["main"].(map[string]interface{})["temp_min"].(float64),
		MainTempMax:   APIResponse["main"].(map[string]interface{})["temp_max"].(float64),
		MainPressure:  int(APIResponse["main"].(map[string]interface{})["pressure"].(float64)),
		MainFeelsLike: APIResponse["main"].(map[string]interface{})["feels_like"].(float64),
		MainHumidity:  int(APIResponse["main"].(map[string]interface{})["humidity"].(float64)),
		WindSpeed:     APIResponse["wind"].(map[string]interface{})["speed"].(float64),
		WindDeg:       int(APIResponse["wind"].(map[string]interface{})["deg"].(float64)),
		CloudsAll:     int(APIResponse["clouds"].(map[string]interface{})["all"].(float64)),
		SysCountry:    APIResponse["sys"].(map[string]interface{})["country"].(string),
		SysSunrise:    time.Unix(int64(APIResponse["sys"].(map[string]interface{})["sunrise"].(float64))+int64(APIResponse["timezone"].(float64)), 0),
		SysSunset:     time.Unix(int64(APIResponse["sys"].(map[string]interface{})["sunset"].(float64))+int64(APIResponse["timezone"].(float64)), 0),
		Name:          APIResponse["name"].(string),
		Base:          APIResponse["base"].(string),
		Visibility:    int(APIResponse["visibility"].(float64)),
		Dt:            time.Unix(int64(APIResponse["dt"].(float64)), 0),
	}

	weatherCurrentData := WeatherData{
		Lon:           APIResponse["coord"].(map[string]interface{})["lon"].(float64),
		Lat:           APIResponse["coord"].(map[string]interface{})["lat"].(float64),
		MainTemp:      APIResponse["main"].(map[string]interface{})["temp"].(float64),
		MainTempMin:   APIResponse["main"].(map[string]interface{})["temp_min"].(float64),
		MainTempMax:   APIResponse["main"].(map[string]interface{})["temp_max"].(float64),
		MainPressure:  int(APIResponse["main"].(map[string]interface{})["pressure"].(float64)),
		MainFeelsLike: APIResponse["main"].(map[string]interface{})["feels_like"].(float64),
		MainHumidity:  int(APIResponse["main"].(map[string]interface{})["humidity"].(float64)),
		WindSpeed:     APIResponse["wind"].(map[string]interface{})["speed"].(float64),
		WindDeg:       int(APIResponse["wind"].(map[string]interface{})["deg"].(float64)),
		CloudsAll:     int(APIResponse["clouds"].(map[string]interface{})["all"].(float64)),
		SysCountry:    APIResponse["sys"].(map[string]interface{})["country"].(string),
		SysSunrise:    time.Unix(int64(APIResponse["sys"].(map[string]interface{})["sunrise"].(float64))+int64(APIResponse["timezone"].(float64)), 0),
		SysSunset:     time.Unix(int64(APIResponse["sys"].(map[string]interface{})["sunset"].(float64))+int64(APIResponse["timezone"].(float64)), 0),
		Name:          APIResponse["name"].(string),
		Base:          APIResponse["base"].(string),
		Visibility:    int(APIResponse["visibility"].(float64)),
		Dt:            time.Unix(int64(APIResponse["dt"].(float64)), 0),
	}

	// Forward Request to insert data
	// Initialize MySQL connector
	connector := MySQLConnect("localhost", 3306, "root", "Pankaj@569", "open_weather")

	// Connect to MySQL
	ormdb, dbConn, err := connector.Connect()
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}
	defer dbConn.Close()

	// Verify Table Exist
	var weatherHistoryTable []WeatherDataHistory
	var weatherTable []WeatherData
	historyTableCheck, weatherTableCheck := ormdb.First(&weatherHistoryTable), ormdb.First(&weatherTable)

	if historyTableCheck != nil {
		migrateTable(ormdb, &weatherHistoryTable)
	}

	if weatherTableCheck != nil {
		migrateTable(ormdb, &weatherTable)
	}

	// Insert Data into Weather Data History Table
	insertData := ormdb.Create(&weatherHistoryData)
	if insertData.Error != nil {
		fmt.Println(insertData.Error)
	}

	var weatherSearch = WeatherData{Lat: weatherHistoryData.Lat, Lon: weatherHistoryData.Lon}
	findResult := ormdb.First(&weatherSearch)
	if findResult.RowsAffected == 0 {
		findResult = ormdb.Create(&weatherCurrentData)
	} else {
		findResult = ormdb.Model(&WeatherData{}).Where("Lat = ?", weatherHistoryData.Lat).Where("Lon = ?", weatherHistoryData.Lon).Updates(&weatherCurrentData)
	}

	if findResult.Error != nil {
		fmt.Println(findResult.Error)
	}
}

func migrateTable(ormdb *gorm.DB, ns interface{}) {
	err := ormdb.AutoMigrate(ns)
	if err != nil {
		log.Fatalf("Error auto-migrating schema: %v", err)
	}
}
