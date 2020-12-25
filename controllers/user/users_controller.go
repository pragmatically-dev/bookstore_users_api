package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {

	ctx.JSON(http.StatusNotImplemented, "not implemented")
}

func CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "not implemented")
}

func SearchUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "not implemented")
}
