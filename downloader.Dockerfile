FROM alpine:latest

WORKDIR /usr/src/app

COPY ./cmd/downloader/crontab/open-weather-api-downloader-cron .
COPY ./cmd/downloader/crontab/entry.sh .
RUN chmod +x /usr/src/app/entry.sh
COPY ./cmd/downloader/cities.json .
COPY ./cmd/downloader/cities.json /
COPY ./cmd/downloader/downloader-linux-amd64 .
RUN /usr/bin/crontab /usr/src/app/open-weather-api-downloader-cron

CMD ["/bin/sh","-c","/usr/src/app/entry.sh"]