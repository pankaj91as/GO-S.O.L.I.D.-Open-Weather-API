package service

import (
	"encoding/json"
	"net/url"

	"models"
	"paggination"

	"gorm.io/gorm"
)

type IWeatherService interface {
	GetWeatherHistoryByLocation(location string, q url.Values) (res []byte, err error)
	GetCurrentWeatherByLocation(location string, q url.Values) (res []byte, err error)
}

type BookService struct {
	DB *gorm.DB
}

func (bookRepository *BookService) GetWeatherHistoryByLocation(location string, q url.Values) (res []byte, err error) {
	var weatherHistoricalData []models.WeatherData
	bookRepository.DB.Table("weather_data_history").Scopes(paggination.Paginate(q)).Where("name = ?", location).Find(&weatherHistoricalData)
	re, err := json.Marshal(weatherHistoricalData)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (bookRepository *BookService) GetCurrentWeatherByLocation(location string, q url.Values) (res []byte, err error) {
	var weatherCurrentData []models.WeatherData
	bookRepository.DB.Scopes(paggination.Paginate(q)).Where("name = ?", location).Find(&weatherCurrentData)
	re, err := json.Marshal(weatherCurrentData)
	if err != nil {
		return nil, err
	}
	return re, nil
}
