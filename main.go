package main

import (
	"context"
	"gin-starter/server"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
	dbUri := "neo4j://" + os.Getenv("GRAPHDB_HOST")
	dbUser := os.Getenv("GRAPHDB_USER")
	dbPassword := os.Getenv("GRAPHDB_PASSWORD")
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	server.Init()
}
