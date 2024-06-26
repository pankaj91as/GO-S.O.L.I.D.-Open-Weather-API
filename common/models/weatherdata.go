package models

import "time"

type WeatherData struct {
	ID            uint    `gorm:"primaryKey;<-:create"`
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
