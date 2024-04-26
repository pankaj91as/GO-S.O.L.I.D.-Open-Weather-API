package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"configs"
	"db"
	"models"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("cronjob")

type OpenWeatherAPI struct{}

func mapAPIResponseToWeatherData(APIResponse map[string]interface{}) models.WeatherData {
	return models.WeatherData{
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
}

func (ll *OpenWeatherAPI) SaveData(latitude float64, longitude float64, dbConnection *db.SQLConnection) {
	// Waitgroup Done Signal
	defer wg.Done()

	// Create API Endpoint
	endpoint := fmt.Sprintf(configs.OpenWeatherAPIEndpoint+"?lat=%v&lon=%v&appid=%v", latitude, longitude, configs.OpenWeatherAPIAPIKEY)
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

	weatherCurrentData := mapAPIResponseToWeatherData(APIResponse)

	// Forward Request to insert data
	insertData(dbConnection, weatherCurrentData)
}

func insertData(dbConnection *db.SQLConnection, weatherCurrentData models.WeatherData) {
	ormdb := dbConnection.GormConn

	// Insert Data into Weather Data History Table
	insertData := ormdb.Table("weather_data_history").Create(&weatherCurrentData)
	if insertData.Error != nil {
		Log.Error(insertData.Error)
	}

	var weatherSearch []models.WeatherData
	findResult := ormdb.Where(&models.WeatherData{Lat: weatherCurrentData.Lat, Lon: weatherCurrentData.Lon}).First(&weatherSearch)
	if findResult.RowsAffected == 0 {
		findResult = ormdb.Create(&weatherCurrentData)
	} else {
		findResult = ormdb.Model(&models.WeatherData{}).Where("Lat = ?", weatherCurrentData.Lat).Where("Lon = ?", weatherCurrentData.Lon).Updates(&weatherCurrentData)
	}

	if findResult.Error != nil {
		fmt.Println(findResult.Error)
	}

	Log.Info("Data processed successfully...")
}
