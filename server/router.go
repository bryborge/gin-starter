package server

import (
	"gin-starter/controllers"
	ApiV1 "gin-starter/controllers/api/v1"
	"net/url"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
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

	// TODO: Encapsulate and move to auth middleware
	issuerURL, _ := url.Parse(os.Getenv("AUTH0_ISSUER_URL"))
	audience := os.Getenv("AUTH0_AUDIENCE")
	provider := jwks.NewCachingProvider(issuerURL, time.Duration(5*time.Minute))
	jwtValidator, _ := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
	)
	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	// Protected routes
	protectedRoutes := router.Group("/api/v1")
	protectedRoutes.Use(adapter.Wrap(jwtMiddleware.CheckJWT))
	protectedRoutes.GET("/pokemon", ApiV1.GetPokemon)

	return router
}
