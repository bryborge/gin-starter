package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Status(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, [1]string{"Working!"})
}
