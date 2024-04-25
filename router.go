package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/pankaj91as/open-weather-api/app/controller"
	"github.com/pankaj91as/open-weather-api/pkg/logger"
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
	InitWeatherRoutes(mRouter)
	return mRouter
}

func InitWeatherRoutes(r *mux.Router) {
	r.HandleFunc("/weather", controller.LandingPage)
	r.HandleFunc("/weather/{location}", controller.FetchWeather).Methods(http.MethodGet)
	r.HandleFunc("/weather/current/{location}", controller.FetchCurrentWeather).Methods(http.MethodGet)
}
