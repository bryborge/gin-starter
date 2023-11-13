package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"starter_api/controllers"
	ApiV1 "starter_api/controllers/api/v1"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	setEnv()
	serveApp()
}

func setEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApp() {
	// Setup Auth0 JWT Validation
	// TODO: Create Auth0 middleware wrapper
	issuerURL, _ := url.Parse(os.Getenv("AUTH0_ISSUER_URL"))
	audience := os.Getenv("AUTH0_AUDIENCE")
	provider := jwks.NewCachingProvider(issuerURL, time.Duration(5*time.Minute))
	jwtValidator, _ := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
	)

	// Create Middlewares
	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	// Create Server
	router := gin.Default()

	// Public Routes
	publicRoutes := router.Group("/")
	publicRoutes.GET("/", controllers.Home)

	// Private Routes
	privateRoutes := router.Group("/api/v1")
	privateRoutes.Use(adapter.Wrap(jwtMiddleware.CheckJWT))
	privateRoutes.GET("/pokemon", ApiV1.GetPokemon)

	// Serve
	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
