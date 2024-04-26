package service

import (
	"encoding/json"
	"net/url"

	"github.com/pankaj91as/open-weather-api/common/models"
	"github.com/pankaj91as/open-weather-api/pkg/db"
	"github.com/pankaj91as/open-weather-api/pkg/paggination"
)

type IWeatherService interface {
	GetWeatherHistoryByLocation(location string, q url.Values) (res []byte, err error)
	GetCurrentWeatherByLocation(location string, q url.Values) (res []byte, err error)
}

type BookRepository struct {
	DB *db.SQLConnection
}

func (bookRepository *BookRepository) GetWeatherHistoryByLocation(location string, q url.Values) (res []byte, err error) {
	var weatherHistoricalData []models.WeatherData
	bookRepository.DB.GormConn.Table("weather_data_history").Scopes(paggination.Paginate(q)).Where("name = ?", location).Find(&weatherHistoricalData)
	re, err := json.Marshal(weatherHistoricalData)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (bookRepository *BookRepository) GetCurrentWeatherByLocation(location string, q url.Values) (res []byte, err error) {
	var weatherCurrentData []models.WeatherData
	bookRepository.DB.GormConn.Scopes(paggination.Paginate(q)).Where("name = ?", location).Find(&weatherCurrentData)
	re, err := json.Marshal(weatherCurrentData)
	if err != nil {
		return nil, err
	}
	return re, nil
}
