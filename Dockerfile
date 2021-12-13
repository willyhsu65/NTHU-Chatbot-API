FROM golang:1.13-alpine

WORKDIR /usr/src/app

RUN apk add build-base

COPY . /usr/src/app
RUN cd /usr/src/app && go build -o app
EXPOSE 8080
ENTRYPOINT ./app