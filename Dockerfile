FROM golang:latest

ENV PROJECT_REPO=github.com/gorobot-nz/go-reserve-task
ENV APP_PATH=/go/src/${PROJECT_REPO}/
WORKDIR ${APP_PATH}
COPY . ${APP_PATH}
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api/main.go

FROM alpine:latest
ENV PROJECT_REPO=github.com/gorobot-nz/go-reserve-task
ENV APP_PATH=/go/src/${PROJECT_REPO}/
RUN adduser -S nonrootuser
WORKDIR ${APP_PATH}
COPY --from=0 ${APP_PATH}/app ${APP_PATH}/app
COPY --from=0 ${APP_PATH}/wait-for-postgres.sh ${APP_PATH}/wait-for-postgres.sh
COPY --from=0 ${APP_PATH}/configs/config.yml ${APP_PATH}/configs/config.yml
COPY --from=0 ${APP_PATH}/.env ${APP_PATH}/.env
COPY --from=0 ${APP_PATH}/internal/schema/book.sql ${APP_PATH}/internal/schema/book.sql
RUN chmod +x ${APP_PATH}/wait-for-postgres.sh
USER nonrootuser
CMD ["./app"]
