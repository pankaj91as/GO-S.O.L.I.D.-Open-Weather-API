module emailapi

go 1.22.2

require (
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	models v0.0.0-00010101000000-000000000000
)

replace models => ../../common/models
