package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Hello World")
}
