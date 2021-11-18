FROM golang:1.14.0-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./cmd/api

CMD ["/app/main"]