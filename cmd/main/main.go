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

	port := ":" + os.Getenv("APP_PORT")
	err = server.Run(port)

	if err != nil {
		log.Fatal(err)
	}
}
