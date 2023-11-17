package server

import (
	"gin-starter/controllers"
	ApiV1 "gin-starter/controllers/api/v1"
	"gin-starter/middleware"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

func NewRouter() *gin.Engine {
	// Create Server
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Public routes
	publicRoutes := router.Group("/")
	publicRoutes.GET("/health", new(controllers.HealthController).Status)

	// Protected routes
	protectedRoutes := router.Group("/api/v1")
	protectedRoutes.Use(adapter.Wrap(middleware.EnsureValidToken()))
	protectedRoutes.GET("/pokemon", ApiV1.GetPokemon)

	return router
}
