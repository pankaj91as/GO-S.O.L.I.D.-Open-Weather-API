package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pankaj91as/open-weather-api/app/service"
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
	book, err := weatherController.GetWeatherHistoryByLocation(vars["location"], q)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			weatherController.RespondWithError(w, http.StatusNotFound, "Location not found")
		default:
			weatherController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	weatherController.RespondWithJSON(w, http.StatusOK, book)
}

func (weatherController *WeatherController) GetWeatherData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	q := r.URL.Query()
	book, err := weatherController.GetCurrentWeatherByLocation(vars["location"], q)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			weatherController.RespondWithError(w, http.StatusNotFound, "Location not found")
		default:
			weatherController.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	weatherController.RespondWithJSON(w, http.StatusOK, book)
}
