package main

import (
	"fmt"
	"starter_api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	serveApp()
}

func serveApp() {
	router := gin.Default()

	publicRoutes := router.Group("/api/v1")
	publicRoutes.GET("/pokemon", controllers.GetPokemon)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
