package controllers

import (
	"gin-starter/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreatePokemon(context *gin.Context) {
	// TODO: wrap this as an adapter
	dbUri := "neo4j://" + os.Getenv("GRAPHDB_HOST")
	dbUser := os.Getenv("GRAPHDB_USER")
	dbPassword := os.Getenv("GRAPHDB_PASSWORD")
	neo4jDriver, err := neo4j.NewDriver(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))

	if err != nil {
		log.Fatal("Error creating Neo4j driver:", err)
	}
	defer neo4jDriver.Close()

	var newPokemon models.Pokemon
	if err := context.BindJSON(&newPokemon); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.CreatePokemonInDb(neo4jDriver, newPokemon.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Pokemon in DB"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Pokemon created successfully - " + newPokemon.Name})
}

func GetPokemonByName(context *gin.Context) {
	// TODO
	dbUri := "neo4j://" + os.Getenv("GRAPHDB_HOST")
	dbUser := os.Getenv("GRAPHDB_USER")
	dbPassword := os.Getenv("GRAPHDB_PASSWORD")
	neo4jDriver, err := neo4j.NewDriver(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))

	if err != nil {
		log.Fatal("Error creating Neo4j driver:", err)
	}
	defer neo4jDriver.Close()

	name := context.Param("name")

	pokemon, err := models.GetPokemonInDbByName(neo4jDriver, name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Pokemon from DB"})
		return
	}

	if (models.Pokemon{}) == pokemon {
		context.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}

	context.JSON(http.StatusOK, pokemon)

	context.IndentedJSON(http.StatusOK, pokemon)
}
