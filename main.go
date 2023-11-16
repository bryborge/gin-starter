package main

import (
	"log"
	"starter_api/server"

	"github.com/joho/godotenv"
)

func main() {
	setEnv()
	server.Init()
}

func setEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
