package main

import "time"

type WeatherData struct {
	ID            uint    `gorm:"primaryKey"`
	Lon           float64 `gorm:"index" json:"Coord.Lon"`
	Lat           float64 `gorm:"index" json:"Coord.Lat"`
	MainTemp      float64 `json:"Main.Temp"`
	MainTempMin   float64 `json:"Main.TempMin"`
	MainTempMax   float64 `json:"Main.TempMax"`
	MainPressure  int     `json:"Main.Pressure"`
	MainFeelsLike float64 `json:"Main.FeelsLike"`
	MainHumidity  int     `json:"Main.Humidity"`
	WindSpeed     float64
	WindDeg       int
	CloudsAll     int
	SysCountry    string
	SysSunrise    time.Time
	SysSunset     time.Time
	Name          string `gorm:"index"`
	Base          string
	Visibility    int
	Dt            time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type WeatherAPIResponse struct {
	Coord      Coord   `json:"coord"`
	Weather    Weather `json:"weather"`
	Base       string  `json:"base"`
	Main       Main    `json:"main"`
	Visibility int     `json:"visibility"`
	Wind       Wind    `json:"wind"`
	Clouds     Clouds  `json:"clouds"`
	Dt         int     `json:"dt"`
	Sys        Sys     `json:"sys"`
	Timezone   int     `json:"timezone"`
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Cod        int     `json:"cod"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Weather []struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}
