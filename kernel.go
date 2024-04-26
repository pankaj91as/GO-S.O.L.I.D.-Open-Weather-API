package main

import (
	"db"
	"sync"

	"controller"
	"service"
)

type kernel struct{}

type IKernel interface {
	InjectDB() *db.SQLConnection
	InjectWeatherController(*db.SQLConnection) controller.WeatherController
}

var (
	k          *kernel
	kernelOnce sync.Once
)

// Singleton Pattern
func Kernel() IKernel {
	if k == nil {
		kernelOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

func (k *kernel) InjectDB() (d *db.SQLConnection) {
	// Implement DB Connection
	database := db.InitDBConnection(DBhost, DBport, DBusername, DBpassword)
	return database
}

func (k *kernel) InjectWeatherController(d *db.SQLConnection) controller.WeatherController {
	WeatherService := &service.WeatherService{DB: d}
	WeatherController := &controller.WeatherController{IWeatherService: WeatherService, ICommonController: &controller.CommonController{}}
	return *WeatherController
}
