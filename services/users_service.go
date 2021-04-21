package services

import (
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.APIErrors) {
	var errs errors.APIErrors
	if userID <=0{
		errs.AddError(errors.NewBadRequestError("Invalid user ID","Could not parse ID"))
	}
	user := users.User{
		ID: userID,
	}
	getErrs:=user.Get()
	if getErrs != nil{
		return nil, getErrs
	}

	return &user,nil

}

func CreateUser(user users.User) (*users.User, *errors.APIErrors) {
	if errs := user.Validate(); errs != nil {
		return nil, errs
	}
	saveErrs:=user.Save()
	if saveErrs!=nil{
		return nil,saveErrs
	}

	return &user, nil
}
