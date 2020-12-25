package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/services"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

func GetUser(ctx *gin.Context) {

	ctx.JSON(http.StatusNotImplemented, "not implemented")
}

//CreateUser controller for create an user
func CreateUser(ctx *gin.Context) {
	var user users.User
	var errs errors.APIErrors
	if err := ctx.BindJSON(&user); err != nil {
		errs.AddError(errors.NewBadRequestError("Invalid JSON body", err.Error()))
		ctx.JSON(http.StatusBadRequest, errs)
		return
	}
	res, createErr := services.CreateUser(user)
	if createErr != nil {
		errs.Errors = append(errs.Errors, createErr. ...)
		ctx.JSON(http.StatusInternalServerError, errs)
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func SearchUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "not implemented")
}
