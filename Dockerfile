FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o go-tech-task ./cmd/api/main.go

CMD ["./go-tech-task"]