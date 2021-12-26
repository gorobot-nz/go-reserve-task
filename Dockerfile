#Build stage
FROM golang:latest AS builder
RUN mkdir -p /urs/src/app
WORKDIR /urs/src/app

RUN go version
ENV GOPATH=/

COPY . /urs/src/app

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go mod tidy
RUN go build -o go-tech-task ./cmd/api/main.go

CMD ["./go-tech-task"]