package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/services"
	"github.com/pragmatically-dev/bookstore_users_api/utils"
)

func GetUser(ctx *gin.Context) {

	ctx.JSON(http.StatusNotImplemented, "not implemented")
}

func CreateUser(ctx *gin.Context) {
	var user users.User
	var errors utils.APIErrors
	if err := ctx.BindJSON(&user); err != nil {
		errors.AddError(&utils.UserError{Reason: "Invalid json body", Msg: err.Error()}, http.StatusBadRequest)
		ctx.JSON(errors.Code, errors)
		return
	}
	res, createErr := services.CreateUser(user)
	if createErr != nil {
		errors.Errors = append(errors.Errors, createErr.Errors...)
		ctx.JSON(errors.Code, errors)
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func SearchUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "not implemented")
}
