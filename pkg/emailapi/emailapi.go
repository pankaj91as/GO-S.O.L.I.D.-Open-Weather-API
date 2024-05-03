package emailapi

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"models"

	"github.com/op/go-logging"
	"gopkg.in/gomail.v2"
)

var Log = logging.MustGetLogger("emailapi")

var EWG sync.WaitGroup

type IEmailAPI interface {
	PrepareBody(body []models.WeatherData) string
	SendMail(b []byte)
	IncrementWG()
	WaitWG()
}

func PrepareBody(body models.WeatherData) string {
	b := fmt.Sprintf(`Weather Change Detected!
<br><br>
City: %s,<br>
Temperature: %f,<br>
temperature Min: %f,<br>
temperature Max: %f,<br>
Pressure: %v,<br>
Feels Like: %f,<br>
Humidity: %v,<br>
Wind Speed: %f,<br>
WindDeg: %v,<br>
Clouds All: %v,<br>
Visibility: %v

`, body.Name,
		body.MainTemp,
		body.MainTempMin,
		body.MainTempMax,
		body.MainPressure,
		body.MainFeelsLike,
		body.MainHumidity,
		body.WindSpeed,
		body.WindDeg,
		body.CloudsAll,
		body.Visibility)

	return b
}

func SendMail(b []byte) {
	defer EWG.Done()

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", os.Getenv("NOTIFICATION_USERS"))
	m.SetHeader("Subject", "Weather Update!")
	m.SetBody("text/html", string(b))

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		Log.Error(err)
	}

	Log.Info("Email Notification Sent...")
}
