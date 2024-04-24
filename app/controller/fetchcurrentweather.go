package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pankaj91as/open-weather-api/common/models"
	"github.com/pankaj91as/open-weather-api/pkg/db"
	"github.com/pankaj91as/open-weather-api/pkg/paggination"
)

func FetchCurrentWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var weatherHistoricalData []models.WeatherData
	gormCon := db.InitDBConnection()
	gormCon.GormConn.Scopes(paggination.Paginate(r)).Where("name = ?",vars["location"]).Find(&weatherHistoricalData)
	re, err := json.Marshal(weatherHistoricalData)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"ok": "error"})
	}
	n := len(re)
	s := string(re[:n])
	fmt.Fprint(w, s)
}