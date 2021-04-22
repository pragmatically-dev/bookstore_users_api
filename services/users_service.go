package services

import (
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.APIErrors) {
	var errs errors.APIErrors
	if userID <= 0 {
		errs.AddError(errors.NewBadRequestError("Invalid user ID", "Could not parse ID"))
	}
	user := users.User{
		ID: userID,
	}
	getErrs := user.Get()
	if getErrs != nil {
		return nil, getErrs
	}

	return &user, nil

}

func CreateUser(user users.User) (*users.User, *errors.APIErrors) {
	if errs := user.Validate(); errs != nil {
		return nil, errs
	}
	saveErrs := user.Save()
	if saveErrs != nil {
		return nil, saveErrs
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.APIErrors) {
	current, err := GetUser(user.ID) //respecting SOLID
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userID int64) *errors.APIErrors {
	user, err := GetUser(userID) //respecting SOLID
	if err != nil {
		return err
	}
	errs := user.Delete()
	if errs != nil {
		return errs
	}
	return nil
}

func FindByStatus(status string) ([]users.User, *errors.APIErrors) {
	dao := &users.User{}
	return dao.FindByStatus(status)

}
