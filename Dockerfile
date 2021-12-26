#Build stage
FROM golang:latest AS builder

WORKDIR /app

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go mod tidy
RUN go build -o go-tech-task ./cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder ./app/wait-for-postgres.sh .
COPY --from=builder ./app/go-tech-task .

RUN apk --update add postgresql-client

CMD ["./app/go-tech-task"]