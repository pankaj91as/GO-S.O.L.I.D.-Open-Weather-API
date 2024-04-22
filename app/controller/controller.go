package controller

import "net/http"

type IController interface {
	LandingPage(w http.ResponseWriter, r *http.Request)
	FetchWeather(w http.ResponseWriter, r *http.Request)
	FetchLatestWeather(w http.ResponseWriter, r *http.Request)
}