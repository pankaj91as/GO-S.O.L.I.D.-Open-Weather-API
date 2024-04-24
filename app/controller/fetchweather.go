package controller

import (
	"encoding/json"
	"net/http"

	"github.com/pankaj91as/open-weather-api/common/models"
	"github.com/pankaj91as/open-weather-api/pkg/db"
)

func FetchWeather(w http.ResponseWriter, r *http.Request) {
	var weatherHistoricalData []models.WeatherData
	gormCon := db.InitDBConnection()
	result := gormCon.GormConn.Find(&weatherHistoricalData)
	json.NewEncoder(w).Encode(map[string]int64{"fetch-weather": result.RowsAffected})
}
