FROM alpine:latest

WORKDIR /usr/src/app

COPY .env .
COPY ./weatherapi-linux-amd64 /usr/local/bin/weatherapi-linux-amd64

EXPOSE 8080:8080

CMD ["weatherapi-linux-amd64"]