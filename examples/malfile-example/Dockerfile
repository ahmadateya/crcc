FROM golang:1.10.0-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN apk add --update --no-cache bash vim