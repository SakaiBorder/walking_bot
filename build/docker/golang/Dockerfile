FROM golang:1.16.3-alpine

RUN apk update && apk add git

RUN mkdir -p /go/app/src

WORKDIR /go/app/

RUN go get -u github.com/cosmtrek/air
