module open-weather-api

go 1.22.2

require (
	controller v0.0.0-00010101000000-000000000000
	db v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	logger v0.0.0-00010101000000-000000000000
	service v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/mysql v1.5.6 // indirect
	gorm.io/gorm v1.25.9 // indirect
	models v0.0.0-00010101000000-000000000000 // indirect
	paggination v0.0.0-00010101000000-000000000000 // indirect
)

replace controller => ./app/controller

replace service => ./app/service

replace paggination => ./pkg/paggination

replace models => ./common/models

replace db => ./pkg/db

replace logger => ./pkg/logger
