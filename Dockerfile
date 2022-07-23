# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 8000

CMD [ "/app/main" ]