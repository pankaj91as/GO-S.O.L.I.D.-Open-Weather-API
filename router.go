package main

import (
	"net/http"
	"sync"

	"controller"
	"logger"

	"github.com/gorilla/mux"
)

type router struct{}

type IRouter interface {
	InitRouter() *mux.Router
}

var (
	r    *router
	once sync.Once
)

// Making Router instance as singleton
func Router() IRouter {
	if r == nil {
		once.Do(func() {
			r = &router{}
		})
	}
	return r
}

func (r *router) InitRouter() *mux.Router {
	mRouter := mux.NewRouter().StrictSlash(true)
	mRouter.Use(logger.LoggingMiddleware)
	kernel := Kernel()
	weatherController := kernel.InjectWeatherController()
	InitWeatherRoutes(mRouter, &weatherController)
	return mRouter
}

func InitWeatherRoutes(r *mux.Router, weatherController *controller.WeatherController) {
	r.HandleFunc("/weather", weatherController.WeatherEndpoint)
	r.HandleFunc("/weather/{location}", weatherController.GetWeatherHistoryData).Methods(http.MethodGet)
	r.HandleFunc("/weather/current/{location}", weatherController.GetWeatherData).Methods(http.MethodGet)
}
