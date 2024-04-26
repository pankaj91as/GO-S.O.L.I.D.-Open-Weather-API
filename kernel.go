package main

import (
	"sync"

	"controller"
	"service"
)

type kernel struct{}

type IKernel interface {
	InjectWeatherController() controller.WeatherController
}

var (
	k          *kernel
	kernelOnce sync.Once
	err        error
)

// Singleton
func Kernel() IKernel {
	if k == nil {
		kernelOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

func (k *kernel) InjectWeatherController() controller.WeatherController {
	WeatherService := &service.BookService{}
	WeatherController := &controller.WeatherController{IWeatherService: WeatherService, ICommonController: &controller.CommonController{}}
	return *WeatherController
}
