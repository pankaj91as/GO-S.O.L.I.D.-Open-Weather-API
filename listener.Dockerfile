FROM alpine:latest

WORKDIR /usr/src/app

COPY .env /usr/.env
COPY ./cmd/listener/listener-linux-amd64 .

CMD ["/bin/sh", "-c", "/usr/src/app/listener-linux-amd64 -RABBIT_HOST=queue -RABBIT_PORT=5672 -RABBIT_USERNAME=guest -RABBIT_PASSWORD=guest"]