package main

import (
	"jimdel/pkg/server"
	"log"
)

func main() {

	port := ":8080"
	err := server.Run(port)

	if err != nil {
		log.Fatal(err)
	}
}
