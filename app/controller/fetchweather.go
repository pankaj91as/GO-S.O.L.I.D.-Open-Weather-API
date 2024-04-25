package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pankaj91as/open-weather-api/common/models"
	"github.com/pankaj91as/open-weather-api/pkg/db"
	"github.com/pankaj91as/open-weather-api/pkg/paggination"
)

func FetchWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var weatherHistoricalData []models.WeatherData
	port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	gormCon := db.InitDBConnection(os.Getenv("MYSQL_HOST"), port, os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"))
	gormCon.GormConn.Table("weather_data_history").Scopes(paggination.Paginate(r)).Where("name = ?",vars["location"]).Find(&weatherHistoricalData)
	re, err := json.Marshal(weatherHistoricalData)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"ok": "error"})
	}
	n := len(re)
	s := string(re[:n])
	fmt.Fprint(w, s)
}
