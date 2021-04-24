package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/services/userservice"

	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

var (
	_IUserService = &userservice.UserService{}
)

func Get(ctx *gin.Context) {
	userID := getUserID(ctx)

	user, errsGetUser := _IUserService.GetUser(userID)
	if errsGetUser != nil {
		ctx.JSON(http.StatusNotFound, errsGetUser)
		return
	}
	sendValidResponse(ctx, user)
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
	res, createErr := _IUserService.CreateUser(user)

	if createErr != nil {
		errs.Errors = append(errs.Errors, createErr.Errors...)
		sendWithLastErrorCode(ctx, &errs)
		return
	}
	sendValidResponse(ctx, res)
}

func Search(ctx *gin.Context) {
	var errs errors.APIErrors
	status := ctx.Query("status")
	users, err := _IUserService.FindByStatus(status)
	if err != nil {
		errs.Errors = append(errs.Errors, err.Errors...)
		sendWithLastErrorCode(ctx, &errs)
		return
	}
	isPublic := ctx.GetHeader("X-Public") == "true"
	ctx.JSON(http.StatusFound, users.Marshall(isPublic))
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
	//isPartial check the  http method of the request
	isPartial := ctx.Request.Method == http.MethodPatch
	res, updateErr := _IUserService.UpdateUser(isPartial, user)
	if updateErr != nil {
		errs.Errors = append(errs.Errors, updateErr.Errors...)
		sendWithLastErrorCode(ctx, &errs)
		return
	}
	sendValidResponse(ctx, res)
}

func Delete(ctx *gin.Context) {
	userID := getUserID(ctx)
	errs := _IUserService.DeleteUser(userID)
	if errs != nil {
		sendWithLastErrorCode(ctx, errs)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"Status": "Deleted"})
}

//----------------------utils for the controllers-------------------------
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
func sendWithLastErrorCode(ctx *gin.Context, errs *errors.APIErrors) {
	lastIndex := len(errs.Errors) - 1
	lastErrorCode := errs.Errors[lastIndex].(*errors.UserError).Code
	ctx.JSON(lastErrorCode, errs)
}

func sendValidResponse(ctx *gin.Context, user *users.User) {
	isPublic := ctx.GetHeader("X-Public") == "true"
	ctx.JSON(http.StatusFound, user.Marshall(isPublic))
}
