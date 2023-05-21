package main

import (
	"awesomeProject/internal/api"
	"awesomeProject/internal/config"
	"log"
)

// @BasePath /

func main() {
	log.Println("Application start")

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(cfg)
	server.StartServer()

	log.Println("Application terminated")
}
