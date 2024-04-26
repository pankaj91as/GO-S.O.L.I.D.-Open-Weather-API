package service

import (
	"encoding/json"
	"net/url"

	"models"
	"paggination"

	"db"
)

type IWeatherService interface {
	GetWeatherHistoryByLocation(location string, q url.Values) (res []byte, err error)
	GetCurrentWeatherByLocation(location string, q url.Values) (res []byte, err error)
}

type BookService struct {
	DB *db.SQLConnection
}

func (bookRepository *BookService) GetWeatherHistoryByLocation(location string, q url.Values) (res []byte, err error) {
	var weatherHistoricalData []models.WeatherData
	bookRepository.DB.GormConn.Table("weather_data_history").Scopes(paggination.Paginate(q)).Where("name = ?", location).Find(&weatherHistoricalData)
	re, err := json.Marshal(weatherHistoricalData)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (bookRepository *BookService) GetCurrentWeatherByLocation(location string, q url.Values) (res []byte, err error) {
	var weatherCurrentData []models.WeatherData
	bookRepository.DB.GormConn.Scopes(paggination.Paginate(q)).Where("name = ?", location).Find(&weatherCurrentData)
	re, err := json.Marshal(weatherCurrentData)
	if err != nil {
		return nil, err
	}
	return re, nil
}
