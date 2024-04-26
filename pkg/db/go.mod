module github.com/pankaj91as/open-weather-api/pkg/db

go 1.22.2

require (
	gorm.io/driver/mysql v1.5.6
	gorm.io/gorm v1.25.9
	github.com/pankaj91as/open-weather-api/pkg/logger v0.0.0-20240426003759-0f10fb8eb655
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

replace github.com/op/go-logging => ../logger
