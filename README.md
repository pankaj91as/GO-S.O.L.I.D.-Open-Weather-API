# Open Weather API Implimentation In GO

Using SOLID Principle We have integrated Open Weather API in Go.

## Step 1 - Clone repository

git clone https://github.com/pankaj91as/Open-Weather-API.git

## CRON (optional)

Weather API data downloader CRON is available in cmd/downloader/crontab/open-weather-api-downloader-cron
Modify CRON time if required.

## Step 2 - Generate Linux Build

(root) Open-Weather-API > env GOOS=linux GOARCH=amd64 go build -o weatherapi-linux-amd64 .

(root) Open-Weather-API > cd cmd/downloader

Open-Weather-API/cmd/downloader > env GOOS=linux GOARCH=amd64 go build -o downloader-linux-amd64 .

## Step 3 - Create Docker image from Linux Build

(root) Open-Weather-API > docker compose up

## Step 4 - Start/Restart docker container

Please follow the squence

1. Wait for MySQL instance up & running
2. Start downloader container
3. Start service container

#### Weather Data Downloader Log

Downloader log is available in downloader container on following path
var/log/open-weather-api-downloader-cron.log
