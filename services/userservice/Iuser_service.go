package userservice

import (
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

type IUserService interface {
	GetUser(userID int64) (*users.User, *errors.APIErrors)
	CreateUser(user users.User) (*users.User, *errors.APIErrors)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.APIErrors)
	DeleteUser(userID int64) *errors.APIErrors
	FindByStatus(status string) (users.Users, *errors.APIErrors)
	LoginUser(users.LoginRequest) (*users.User, *errors.APIErrors)
}
