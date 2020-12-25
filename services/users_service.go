package services

import (
	"errors"

	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.APIErrors) {

	return &user, nil
}
