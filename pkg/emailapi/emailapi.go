package emailapi

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"models"

	"gopkg.in/gomail.v2"
)

var EWG sync.WaitGroup

type IEmailAPI interface {
	PrepareBody(body []models.WeatherData) string
	SendMail(b []byte)
	IncrementWG()
	WaitWG()
}

func PrepareBody(body []models.WeatherData) string {
	b := fmt.Sprintf(`Weather Change Detected!

City: %s,
Temperature: %f,
temperature Min: %f,
temperature Max: %f,
Pressure: %v,
Feels Like: %f,
Humidity: %v,
Wind Speed: %f,
WindDeg: %v,
Clouds All: %v,
Visibility: %v

`, body[0].Name,
		body[0].MainTemp,
		body[0].MainTempMin,
		body[0].MainTempMax,
		body[0].MainPressure,
		body[0].MainFeelsLike,
		body[0].MainHumidity,
		body[0].WindSpeed,
		body[0].WindDeg,
		body[0].CloudsAll,
		body[0].Visibility)

	return b
}

func SendMail(b []byte) {
	defer EWG.Done()

	m := gomail.NewMessage()
	m.SetHeader("From", "nextgensoft@zohomail.in")
	m.SetHeader("To", "pankaj91as@gmail.com")
	m.SetHeader("Subject", "Weather Update!")
	m.SetBody("text/html", string(b))

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
