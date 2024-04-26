module github.com/pankaj91as/open-weather-api/app/controller

go 1.22.2

require (
	github.com/gorilla/mux v1.8.1
	github.com/pankaj91as/open-weather-api/app/service v0.0.0-20240426003759-0f10fb8eb655
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pankaj91as/open-weather-api/common/models v0.0.0-20240426003759-0f10fb8eb655 // indirect
	github.com/pankaj91as/open-weather-api/pkg/db v0.0.0-20240426003418-12d4da1e124e // indirect
	github.com/pankaj91as/open-weather-api/pkg/paggination v0.0.0-20240426003759-0f10fb8eb655 // indirect
	gorm.io/driver/mysql v1.5.6 // indirect
	gorm.io/gorm v1.25.9 // indirect
)

replace github.com/pankaj91as/open-weather-api/common/models => ../../common/models

replace github.com/pankaj91as/open-weather-api/pkg/db => ../../pkg/db

replace github.com/pankaj91as/open-weather-api/pkg/paggination => ../../pkg/paggination

replace github.com/pankaj91as/open-weather-api/app/service => ../service
