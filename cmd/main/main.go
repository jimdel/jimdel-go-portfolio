package main

import (
	"jimdel/pkg/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := ":8080"
	if envPort := os.Getenv("APP_PORT"); envPort != "" {
		port = ":" + envPort
	}
	err = server.Run(port)

	if err != nil {
		log.Fatal(err)
	}
}
