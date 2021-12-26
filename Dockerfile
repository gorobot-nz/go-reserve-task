#Build stage
FROM golang:latest AS builder
RUN mkdir -p /urs/src/app
WORKDIR /urs/src/app

RUN go version
ENV GOPATH=/

COPY . /urs/src/app

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go mod tidy
RUN go build -o go-tech-task ./cmd/api/main.go

FROM alpine:latest

RUN mkdir -p /urs/src/app
WORKDIR /urs/src/app

COPY --from=builder /urs/src/app/wait-for-postgres.sh /urs/src/app
COPY --from=builder /urs/src/app/go-tech-task /urs/src/app

RUN apk --update add postgresql-client

CMD ["./go-tech-task"]