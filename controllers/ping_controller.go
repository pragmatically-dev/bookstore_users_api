package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping ping endpoint
func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
