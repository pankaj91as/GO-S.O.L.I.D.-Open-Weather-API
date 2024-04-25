package service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pankaj91as/open-weather-api/common/models"
	"github.com/pankaj91as/open-weather-api/pkg/db"
	"github.com/pankaj91as/open-weather-api/pkg/paggination"
)

func FindWhere(con *db.SQLConnection, r *http.Request) (res []byte, err error) {
	vars := mux.Vars(r)
	var weatherHistoricalData []models.WeatherData
	con.GormConn.Table("weather_data_history").Scopes(paggination.Paginate(r)).Where("name = ?", vars["location"]).Find(&weatherHistoricalData)
	re, err := json.Marshal(weatherHistoricalData)
	if err != nil {
		return nil, err
	}
	return re, nil
}
