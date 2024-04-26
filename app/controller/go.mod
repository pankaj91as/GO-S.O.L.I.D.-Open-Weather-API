module controller

go 1.22.2

require (
	github.com/gorilla/mux v1.8.1
	service v0.0.0-00010101000000-000000000000
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/gorm v1.25.9 // indirect
	models v0.0.0-00010101000000-000000000000 // indirect
	paggination v0.0.0-00010101000000-000000000000 // indirect
)

replace models => ../../common/models

replace db => ../../pkg/db

replace paggination => ../../pkg/paggination

replace service => ../service
