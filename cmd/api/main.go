package main

import (
	"go-tech-task/internal"
	"log"
)

func main() {
	server := new(internal.Server)
	if err := server.Run(); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
