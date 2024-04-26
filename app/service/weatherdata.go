package service

import (
	"db"
	"encoding/json"
	"net/url"

	"models"
	"paggination"
)

type IWeatherService interface {
	GetWeatherHistoryByLocation(location string, q url.Values) (res []byte, err error)
	GetCurrentWeatherByLocation(location string, q url.Values) (res []byte, err error)
}

type WeatherService struct {
	DB *db.SQLConnection
}

func (WeatherService *WeatherService) GetWeatherHistoryByLocation(location string, q url.Values) (res []byte, err error) {
	var weatherHistoricalData []models.WeatherData
	WeatherService.DB.GormConn.Table("weather_data_history").Scopes(paggination.Paginate(q)).Where("name = ?", location).Find(&weatherHistoricalData)
	re, err := json.Marshal(weatherHistoricalData)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (WeatherService *WeatherService) GetCurrentWeatherByLocation(location string, q url.Values) (res []byte, err error) {
	var weatherCurrentData []models.WeatherData
	WeatherService.DB.GormConn.Scopes(paggination.Paginate(q)).Where("name = ?", location).Find(&weatherCurrentData)
	re, err := json.Marshal(weatherCurrentData)
	if err != nil {
		return nil, err
	}
	return re, nil
}
