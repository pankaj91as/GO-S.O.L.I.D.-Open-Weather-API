# Open Weather API Implimentation In GO

Using SOLID Principle We have integrated Open Weather API in Go.

# Available locations

- Mumbai
- Delhi
- Bangalore
- Kolkata
- Chennai
- Hyderabad
- Ahmedabad
- Pune
- Surat
- Jaipur
- Lucknow
- Kanpur
- Nagpur
- Indore
- Thane
- Bhopal
- Visakhapatnam
- Pimpri-Chinchwad
- Patna
- Vadodara
- Ghaziabad
- Ludhiana
- Agra
- Nashik
- Faridabad
- Meerut
- Rajkot
- Kalyan-Dombivli
- Vasai-Virar
- Varanasi

## Step 1 - Clone repository

```
www % git clone https://github.com/pankaj91as/Open-Weather-API.git
```

## CRON (optional)

Weather API data downloader CRON is available in `cmd/downloader/crontab/open-weather-api-downloader-cron`

Modify CRON time if required, Current it execute **every min**.

## Step 2 - Generate Linux Build

```
Open-Weather-API % env GOOS=linux GOARCH=amd64 go build -o weatherapi-linux-amd64 .
Open-Weather-API % cd cmd/downloader
Open-Weather-API/cmd/downloader % env GOOS=linux GOARCH=amd64 go build -o downloader-linux-amd64 .
```

## Step 3 - Create Docker image from Linux Build

```
Open-Weather-API % docker compose up
```

## Step 4 - Available Endpoints

- http://localhost:8080/weather
- http://localhost:8080/weather/nashik
- http://localhost:8080/weather/current/nashik

## Weather API CRON Log

Downloader log is available in downloader container on following path

```
tail -5 /var/log/open-weather-api-downloader-cron.log
```

## Database Access

- http://localhost:8081
- **host:** db
- **username:** root
- **password:** password

<br>
<br>

---

Note: if any loopback services **e.q: apache, httpd, mysql** running on your machine please stop before deploying container.
