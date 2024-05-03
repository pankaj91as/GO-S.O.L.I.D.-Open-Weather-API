module main

go 1.22.2

require rabbitqueue v0.0.0-00010101000000-000000000000

require (
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	models v0.0.0-00010101000000-000000000000 // indirect
)

require (
	emailapi v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/rabbitmq/amqp091-go v1.9.0 // indirect
)

replace rabbitqueue => ../../pkg/queue

replace emailapi => ../../pkg/emailapi

replace models => ../../common/models
