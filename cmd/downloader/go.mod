module main

go 1.22.2

require (
	github.com/pankaj91as/open-weather-api/common/models v0.0.0-20240425104753-c0fd09243d3b
	github.com/pankaj91as/open-weather-api/pkg/db v0.0.0-20240425104753-c0fd09243d3b
	gorm.io/driver/mysql v1.5.6
	gorm.io/gorm v1.25.9
)

require filippo.io/edwards25519 v1.1.0 // indirect

require (
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
)

replace github.com/pankaj91as/open-weather-api/common/models => ../../common/models

replace github.com/pankaj91as/open-weather-api/pkg/db => ../../pkg/db
