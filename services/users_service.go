package services

import (
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/utils"
)

func CreateUser(user users.User) (*users.User, *utils.APIErrors) {
	return &user, nil
}
