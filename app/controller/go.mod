module github.com/pankaj91as/open-weather-api/app/controller

go 1.22.2

require (
	github.com/pankaj91as/open-weather-api/common/models v0.0.0-00010101000000-000000000000
	github.com/pankaj91as/open-weather-api/pkg/db v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	gorm.io/driver/mysql v1.5.6 // indirect
	gorm.io/gorm v1.25.9 // indirect
)

replace github.com/pankaj91as/open-weather-api/common/models => ../../common/models

replace github.com/pankaj91as/open-weather-api/pkg/db => ../../pkg/db
