package main

import (
	"go-tech-task/internal/server"

	"log"
)

func main() {
	app := server.NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
