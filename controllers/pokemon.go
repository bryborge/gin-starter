package controllers

import (
	"net/http"
	"starter_api/models"

	"github.com/gin-gonic/gin"
)

func GetPokemon(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.MockData)
}
