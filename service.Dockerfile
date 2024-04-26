FROM alpine:latest

WORKDIR /usr/src/app

COPY .env .
COPY ./weatherapi-linux-amd64 /usr/local/bin/weatherapi-linux-amd64

EXPOSE 8080:8080

CMD ["/bin/sh", "-c", "/usr/local/bin/weatherapi-linux-amd64 -db-host=db -db-port=3306 -db-username=root -db-password=password -database=open_weather"]