# Open Weather API Implimentation In GO

Using SOLID Principle We have integrated Open Weather API in Go.

## Clone repository

git clone https://github.com/pankaj91as/Open-Weather-API.git

## Create Docker image from Linux Build

docker build -f Dockerfile.downloader -t weatherdata-downloader .

docker build -f Dockerfile.api -t weatherapi .

## Generate Linux Build

env GOOS=linux GOARCH=amd64 go build -o downloader-linux-amd64 .

env GOOS=linux GOARCH=amd64 go build -o weatherapi-linux-amd64 .
