package main

import (
	log "github.com/sirupsen/logrus"
	"go-tech-task/internal/server"
)

func main() {
	log.SetFormatter(new(log.JSONFormatter))

	app := server.NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
