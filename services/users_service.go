package services

import (
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.APIErrors) {
	if errs := user.Validate(); errs != nil {
		return nil, errs
	}
	return &user, nil
}
