package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/services"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

func GetUser(ctx *gin.Context) {
	var errs errors.APIErrors
	raw, _ := ctx.Params.Get("id")
	UserID, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		errs.AddError(errors.NewBadRequestError("Invalid ID", "Could not parse ID"))
		ctx.JSON(http.StatusBadRequest, errs)
		return
	}
	user, errsGetUser := services.GetUser(UserID)
	if errsGetUser != nil {
		ctx.JSON(http.StatusNotFound, errsGetUser)
		return
	}
	ctx.JSON(http.StatusFound, &user)
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
		errs.Errors = append(errs.Errors, createErr.Errors...)
		ctx.JSON(errs.Errors[len(errs.Errors)-1].(*errors.UserError).Code, errs)
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func SearchUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "not implemented")
}
