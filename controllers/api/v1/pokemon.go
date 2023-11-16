package controllers

import (
	"gin-starter/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPokemon(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.MockData)
}
