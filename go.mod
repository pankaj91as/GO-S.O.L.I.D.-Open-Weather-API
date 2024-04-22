module weatherapi

go 1.22.2

require (
	controller v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	logger v0.0.0-00010101000000-000000000000
)

replace logger => ./pkg/logger

replace controller => ./app/controller
