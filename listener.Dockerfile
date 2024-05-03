FROM alpine:latest

WORKDIR /usr/src/app

COPY .env .
COPY ./cmd/listener/listener-linux-amd64 .

CMD ["/bin/sh", "-c", "/usr/src/app/listener-linux-amd64"]