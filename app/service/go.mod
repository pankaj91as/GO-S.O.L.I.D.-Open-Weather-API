module service

go 1.22.2

require (
	gorm.io/gorm v1.25.9
	models v0.0.0-00010101000000-000000000000
	paggination v0.0.0-00010101000000-000000000000
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

replace models => ../../common/models

replace paggination => ../../pkg/paggination
