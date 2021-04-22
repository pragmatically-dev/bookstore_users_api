package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/services"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

func getUserID(ctx *gin.Context) int64 {
	var errs errors.APIErrors
	raw, _ := ctx.Params.Get("id")
	userID, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		errs.AddError(errors.NewBadRequestError("Invalid ID", "Could not parse ID"))
		ctx.JSON(http.StatusBadRequest, errs)
		return 0
	}
	return userID
}

func Get(ctx *gin.Context) {
	userID := getUserID(ctx)
	user, errsGetUser := services.GetUser(userID)
	if errsGetUser != nil {
		ctx.JSON(http.StatusNotFound, errsGetUser)
		return
	}
	ctx.JSON(http.StatusFound, &user)
}

//CreateUser controller for create an user
func Create(ctx *gin.Context) {
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
		//TODO: Transform this statement to a function. in order to respect DRY principle
		lastErrorCode := errs.Errors[len(errs.Errors)-1].(*errors.UserError).Code
		ctx.JSON(lastErrorCode, errs)
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func Search(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "not implemented")
}

func Update(ctx *gin.Context) {
	var user users.User
	var errs errors.APIErrors
	userID := getUserID(ctx)
	if err := ctx.BindJSON(&user); err != nil {
		errs.AddError(errors.NewBadRequestError("Invalid JSON body", err.Error()))
		ctx.JSON(http.StatusBadRequest, errs)
		return
	}
	user.ID = userID
	isPartial := ctx.Request.Method == http.MethodPatch

	res, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		errs.Errors = append(errs.Errors, updateErr.Errors...)
		//TODO: Transform this statement to a function. in order to respect DRY principle
		lastErrorCode := errs.Errors[len(errs.Errors)-1].(*errors.UserError).Code
		ctx.JSON(lastErrorCode, errs)
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func Delete(ctx *gin.Context) {
	userID := getUserID(ctx)
	errs := services.DeleteUser(userID)
	if errs != nil {
		//TODO: Transform this statement to a function. in order to respect DRY principle
		lastErrorCode := errs.Errors[len(errs.Errors)-1].(*errors.UserError).Code
		ctx.JSON(lastErrorCode, errs)
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"Status": "Deleted"})
}
