#Build stage
FROM golang:latest
RUN mkdir -p /urs/src/app
WORKDIR /urs/src/app

RUN go version
ENV GOPATH=/

COPY . .

RUN go mod download
RUN go mod tidy
RUN go build -o go-tech-task ./cmd/api/main.go

FROM alpine:latest

COPY --from=0 /urs/src/app/wait-for-postgres.sh .
COPY --from=0 /urs/src/app/go-tech-task .

RUN chmod +x wait-for-postgres.sh

RUN apk --update add postgresql-client

RUN useradd -u 8877 john
USER john

CMD ["./go-tech-task"]