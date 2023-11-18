package main

import (
	"gin-starter/db"
	"gin-starter/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db.Init()
	server.Init()
}
