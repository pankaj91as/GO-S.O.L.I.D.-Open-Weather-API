module github.com/pankaj91as/open-weather-api

go 1.22.2

require (
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/pankaj91as/open-weather-api/app/controller v0.0.0-20240425230305-552e6e4f8144
	github.com/pankaj91as/open-weather-api/pkg/logger v0.0.0-20240425230305-552e6e4f8144
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pankaj91as/open-weather-api/app/service v0.0.0-20240426000203-65257dfca010 // indirect
	github.com/pankaj91as/open-weather-api/common/models v0.0.0-20240425230305-552e6e4f8144 // indirect
	github.com/pankaj91as/open-weather-api/pkg/db v0.0.0-20240425230305-552e6e4f8144 // indirect
	github.com/pankaj91as/open-weather-api/pkg/paggination v0.0.0-20240425235456-74d9634ea02d // indirect
	gorm.io/driver/mysql v1.5.6 // indirect
	gorm.io/gorm v1.25.9 // indirect
)

replace github.com/pankaj91as/open-weather-api/app/controller => ./app/controller

replace github.com/pankaj91as/open-weather-api/pkg/logger => ./pkg/logger
