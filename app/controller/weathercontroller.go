package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"service"

	"github.com/gorilla/mux"
)

type WeatherController struct {
	service.IWeatherService
	ICommonController
}

func (weatherController *WeatherController) WeatherEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"weather-api": true})
}

func (weatherController *WeatherController) GetWeatherHistoryData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	q := r.URL.Query()
	weatherData, err := weatherController.GetWeatherHistoryByLocation(vars["location"], q)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			weatherController.RespondWithError(w, http.StatusNotFound, "Location not found")
		default:
			weatherController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	n := len(weatherData)
	s := string(weatherData[:n])
	fmt.Fprint(w, s)
}

func (weatherController *WeatherController) GetWeatherData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	q := r.URL.Query()
	weatherData, err := weatherController.GetCurrentWeatherByLocation(vars["location"], q)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			weatherController.RespondWithError(w, http.StatusNotFound, "Location not found")
		default:
			weatherController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	n := len(weatherData)
	s := string(weatherData[:n])
	fmt.Fprint(w, s)
}
