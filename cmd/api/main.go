package main

import (
	log "github.com/sirupsen/logrus"

	"go-tech-task/internal/server"
)

// @title Books API
// @version 1.0
// @description Simple API for Books store

// @host localhost:8000
// @BasePath /

func main() {
	log.SetFormatter(new(log.JSONFormatter))

	app := server.NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
